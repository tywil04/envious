<script>
    // javascript
    import { GetInvidiousApiInstances, SetSelectedInvidiousInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"
    import { pushRoute } from "../lib/router.js"
    import { urlRegex } from "../lib/validations.js";

    
    // components
    import SelectInput from '../components/inputs/SelectInput.svelte'
    import TextInput from "../components/inputs/TextInput.svelte"
    import Button from "../components/buttons/Button.svelte"
    import WindowShell from "../components/window/WindowShell.svelte";


    let instance = ""
    let selectedIndex = 0

    const onInstanceConfirm = async () => {
        console.log(instance)
        if (instance !== "") {
            const success = await SetSelectedInvidiousInstance(instance)
            if (success) {
                pushRoute("home")
            } else {
                window.location.reload()
            }
        }
    }
</script>


<WindowShell>
    {#await GetInvidiousApiInstances()}
        <p>Loading...</p>
    {:then instances} 
        <p class="info">Please select an instance from the list below or select 'Custom' and enter a custom url pointing towards an API enabled Invidious instance.</p>

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
</WindowShell>


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
        @apply text-sm text-red-600;
    }
</style>