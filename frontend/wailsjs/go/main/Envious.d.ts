// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {invidious} from '../models';

export function DBGet(arg1:string):Promise<string>;

export function DBSet(arg1:string,arg2:string):Promise<void>;

export function GetInvidiousInstance():Promise<invidious.Instance>;

export function GetInvidiousInstances():Promise<Array<invidious.Instance>>;

export function GetTrendingVideos(arg1:invidious.TrendingOption):Promise<Array<invidious.Video>>;

export function GetVideo(arg1:string):Promise<invidious.Video>;

export function RestartApp():Promise<void>;

export function Search(arg1:string,arg2:invidious.SearchOption):Promise<Array<invidious.SearchItem>>;

export function SetInvidiousInstance(arg1:invidious.Instance):Promise<void>;
