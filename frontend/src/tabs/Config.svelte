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
    <div class="seperator">
        <SelectInput bind:selected={instance} label="Select Instance" options={[{ display: "Custom", value: "" }, ...instances.map((instance) => ({ display: instance.apiUrl + " [cors: " + instance.cors + "] [region: " + instance.region + "]", value: instance }))]}/>
    </div>

    <!-- <TextInput bind:value={sessionId} label="Session Id" type="text"/>
    <div class="seperator"></div> -->

    <Button disabled={instance === ""} on:click={configure}>Confirm (this will reload the app)</Button>
{/await}


<style lang="postcss">
    .seperator {
        @apply mb-4;
    }
</style>