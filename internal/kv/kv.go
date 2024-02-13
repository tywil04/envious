package kv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const pathSeperator = string(os.PathSeparator)

var userConfigDir string

func init() {
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	userConfigDir = dir
}

type KV struct {
	Key   string
	Value string
}

type DB struct {
	path     string
	internal map[string]string
}

func NewDB(name string) (*DB, error) {
	db := &DB{
		path:     userConfigDir + pathSeperator + name,
		internal: map[string]string{},
	}

	pathParts := strings.Split(db.path, pathSeperator)
	parentDir := strings.Join(pathParts[0:len(pathParts)-1], pathSeperator)
	err := os.MkdirAll(parentDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	err = db.readFromFile()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DB) writeToFile() error {
	file, err := os.OpenFile(d.path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = json.NewEncoder(file).Encode(&d.internal)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) readFromFile() error {
	file, err := os.OpenFile(d.path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		data = []byte("{}")
	}

	err = json.Unmarshal(data, &d.internal)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Get(key string) (string, error) {
	fmt.Println(d.internal)

	if d.internal == nil {
		return "", errors.New("db not initialised")
	}

	value, ok := d.internal[key]
	if !ok {
		return "", nil
	}

	return value, nil
}

func (d *DB) GetMultiple(keys []string) ([]string, error) {
	values := make([]string, len(keys))
	for i, key := range keys {
		value, err := d.Get(key)
		if err != nil {
			return nil, err
		}
		values[i] = value
	}
	return values, nil
}

func (d *DB) Set(kv KV) error {
	if d.internal == nil {
		return errors.New("db not initialised")
	}

	d.internal[kv.Key] = kv.Value

	return nil
}

func (d *DB) SetMultiple(kvs []KV) error {
	for _, kv := range kvs {
		err := d.Set(kv)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DB) Close() error {
	return d.writeToFile()
}
