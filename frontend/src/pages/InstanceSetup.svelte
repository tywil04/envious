<script>
    // svelte
    import { blur } from "svelte/transition"


    // javascript
    import { GetApiInstances, SetSelectedInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"
    import { push } from "svelte-spa-router"
    import { urlRegex } from "../lib/validations.js";

    
    // components
    import SelectInput from '../components/inputs/SelectInput.svelte'
    import TextInput from "../components/inputs/TextInput.svelte"
    import Button from "../components/buttons/Button.svelte"
    import Main from "../components/window/Main.svelte";


    let instance = ""
    let selectedIndex = 0
    let instanceConfirmError = false

    const onInstanceConfirm = async () => {
        if (instance !== "") {
            instanceConfirmError = false

            const success = await SetSelectedInstance(instance)

            if (success) {
                await push("/")
            } else {
                instanceConfirmError = true
            }
        }
    }

    $: instanceConfirmError = selectedIndex !== 0 ? false: instanceConfirmError
</script>


<Main>
    {#await GetApiInstances()}
        <p>Loading...</p>
    {:then instances} 
        <p class="info">Please select an instance from the list below or select 'Custom' and enter a custom url pointing towards an API enabled Invidious instance.</p>

        {#if selectedIndex === 0 && instanceConfirmError}
            <p class="info errorText" out:blur={{ duration: 250 }} in:blur={{ duration: 250 }}>The custom Invidious instance provided doesn't seem to have its API enabled or its not an Invidious instance.</p>
        {/if}

        <div class="inputContainer">
            <SelectInput bind:selected={instance} bind:selectedIndex label="Select Instance" options={[{ display: "Custom", value: "" }, ...instances]}/>
            <div class="inputSeperator"></div>
            {#if selectedIndex === 0}
                <TextInput bind:value={instance} label="Custom Instance" type="url" errorMessage="Please enter a valid url" placeholder="https://example.org" validation={urlRegex}/>
            {/if}
        </div>

        <Button disabled={instance === ""} on:click={onInstanceConfirm}>Confirm</Button>
    {:catch error}
        <span class="errorText">{error}</span>
    {/await}
</Main>


<style lang="postcss">
    .info {
        @apply mb-3;
    }

    
    .inputSeperator {
        @apply mb-4;
    }

    .inputContainer {
        @apply mb-6;
    }


    .errorText {
        @apply text-red-600;
    }
</style>