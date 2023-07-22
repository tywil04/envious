<script>
    import { GetInstancesApi, Configure } from "../../wailsjs/go/main/Tubed.js"
    import validations from "../lib/validations.js";

    import SelectInput from '../components/inputs/SelectInput.svelte'
    import TextInput from "../components/inputs/TextInput.svelte"
    import Button from "../components/buttons/Button.svelte"


    let provider = ""
    let instance = ""
    // let sessionId = ""
    let selectedIndex = 0

    
    const configure = async () => {
        if (instance !== "") {
            await Configure(provider, instance)
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
                <SelectInput bind:selected={instance} bind:selectedIndex label="Select Instance" options={[{ display: "Custom", value: "" }, ...instances]}/>
                <div class="seperator"></div>
                {#if selectedIndex === 0}
                    <TextInput bind:value={instance} label="Custom Instance" type="url" errorMessage="Please enter a valid url" placeholder="https://example.org" validation={validations.url}/>
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