<script>
    import SelectInput from '../components/SelectInput.svelte'
    import Button from "../components/Button.svelte"

    import { GetInstances, SetSelectedInstance, SetBackendConfigured } from "../../wailsjs/go/main/Tubed.js"


    let instance = null

    
    const configure = async () => {
        if (instance !== null) {
            await SetSelectedInstance(instance)
            await SetBackendConfigured()
        }
    }
</script>


{#await GetInstances()}
    <p>Loading...</p>
{:then instances} 
    <div class="mb-4">
        <SelectInput bind:selected={instance} label="Select Instance" options={instances.map((i) => ({display: i.apiUrl + " [cors: " + i.cors + "] [region: " + i.region + "]", value: i }))}/>
    </div>

    <!-- <TextInput bind:value={sessionId} label="Session Id" type="text"/>
    <div class="seperator"></div> -->

    <Button disabled={instance === ""} on:click={configure}>Confirm (this will reload the app)</Button>
{/await}