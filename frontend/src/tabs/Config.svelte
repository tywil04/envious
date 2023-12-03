<script>
    import SelectInput from '../components/SelectInput.svelte'
    import TextInput from "../components/TextInput.svelte"
    import Button from "../components/Button.svelte"

    import { DBSet, GetInstancesApi } from "../../wailsjs/go/main/Tubed.js"


    let provider = ""
    let instance = ""
    let instanceIndex = 0

    
    const configure = async () => {
        if (instance !== "") {
            await DBSet("backend.provider", provider)
            await DBSet("backend.instanceApi", instance)
            await DBSet("backend.configured", true)
        }
    }
</script>


<SelectInput bind:selected={provider} label="Select Provider" options={[
    { display: "Not selected", value: "" },
    { display: "Piped", value: "piped" },
    { display: "Invidious", value: "invidious" },
]}/>

<div class="seperator"></div>

{#key provider}
    {#if provider !== ""}
        {#await GetInstancesApi(provider)}
            <p>Loading...</p>
        {:then instances} 
            <div class="seperator">
                <SelectInput bind:selected={instance} bind:selectedIndex={instanceIndex} label="Select Instance" options={[{ display: "Custom", value: "" }, ...instances]}/>
                <div class="seperator"></div>
                {#if instanceIndex === 0}
                    <TextInput bind:value={instance} label="Custom Instance" type="url" errorMessage="Please enter a valid url" placeholder="https://example.org"/>
                {/if}
            </div>

            <!-- <TextInput bind:value={sessionId} label="Session Id" type="text"/>
            <div class="seperator"></div> -->

            <Button disabled={instance === ""} on:click={configure}>Confirm (this will reload the app)</Button>
        {:catch error}
            <span class="error">{error}</span>
        {/await}
    {/if}
{/key}


<style lang="postcss">
    .seperator {
        @apply mb-4;
    }

    .error {
        @apply text-red-600;
    }
</style>