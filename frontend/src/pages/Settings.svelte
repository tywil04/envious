<script>
    import { blur } from "svelte/transition"

    import { SetToken, GetApiInstances, SetSelectedInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"
    import validations from "../lib/validations.js";

    import SelectInput from '../components/inputs/SelectInput.svelte.js'
    import TextInput from "../components/inputs/TextInput.svelte.js"
    import Button from "../components/buttons/Button.svelte.js"


    let instance = ""
    let sessionId = ""
    let selectedIndex = 0
    let instanceConfirmError = false


    const selectInstance = async () => {
        if (instance !== "") {
            instanceConfirmError = false

            const success = await SetSelectedInstance(instance)
            await SetToken(sessionId)

            if (success) {
                // await push("/")
            } else {
                instanceConfirmError = true
            }
        }
    }


    $: instanceConfirmError = selectedIndex !== 0 ? false: instanceConfirmError
</script>


{#await GetApiInstances()}
    <p>Loading...</p>
{:then instances} 
    <p class="info">Please select an instance from the list below or select 'Custom' and enter a custom url pointing towards an API enabled Invidious instance.</p>

    {#if selectedIndex === 0 && instanceConfirmError}
        <p class="info errorText" out:blur={{ duration: 250 }} in:blur={{ duration: 250 }}>The custom Invidious instance provided doesn't seem to have its API enabled or its not an Invidious instance.</p>
    {/if}

    <div class="seperator">
        <SelectInput bind:selected={instance} bind:selectedIndex label="Select Instance" options={[{ display: "Custom", value: "" }, ...instances]}/>
        <div class="seperator"></div>
        {#if selectedIndex === 0}
            <TextInput bind:value={instance} label="Custom Instance" type="url" errorMessage="Please enter a valid url" placeholder="https://example.org" validation={validations.url}/>
        {/if}
    </div>

    <TextInput bind:value={sessionId} label="Session Id" type="text"/>
    <div class="seperator"></div>

    <Button disabled={instance === ""} on:click={selectInstance}>Confirm</Button>
{:catch error}
    <span class="error">{error}</span>
{/await}


<style lang="postcss">
    .info {
        @apply mb-3;
    }

    .seperator {
        @apply mb-4;
    }

    .error {
        @apply text-red-600;
    }
</style>