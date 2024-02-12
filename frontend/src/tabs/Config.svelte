<script>
    import SelectInput from '../components/SelectInput.svelte'
    import Button from "../components/Button.svelte"
    import * as go from "../../wailsjs/go/main/Envious.js"
    import * as wails from "../../wailsjs/runtime/runtime.js"


    let instance = null

    
    const configure = () => {
        if (instance !== null) {
            go.SetInvidiousInstance(instance).then(() => {
                go.RestartApp()
            })
        }
    }

    console.log("hi")
</script>


{#await go.GetInvidiousInstances()}
    <p>Loading...</p>
{:then instances}
    <div class="mb-4">
        <SelectInput bind:selected={instance} label="Select Instance" options={instances.map((i) => ({display: i.apiUrl + " [region: " + i.region + "]", value: i }))}/>
    </div>

    <!-- <TextInput bind:value={sessionId} label="Session Id" type="text"/>
    <div class="seperator"></div> -->

    <Button disabled={instance === ""} on:click={configure}>Confirm (this will reload the app)</Button>
{/await}