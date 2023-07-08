<script>
    import { GetInvidiousApiInstances, SetSelectedInvidiousInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"

    import {push} from 'svelte-spa-router'

    import SelectInput from '../components/inputs/SelectInput.svelte'
    import TextInput from "../components/inputs/TextInput.svelte"
    import Button from "../components/buttons/Button.svelte"
    import ErrorSpan from "../components/text/ErrorSpan.svelte";

    const urlValidation = /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/

    let instance = ""
    let selectedIndex = 0

    const onInstanceConfirm = async () => {
        console.log(instance)
        if (instance !== "") {
            const success = await SetSelectedInvidiousInstance(instance)
            if (success) {
                push("/")
            } else {
                window.location.reload()
            }
        }
    }
</script>

<main class="prose p-5 space-y-6 max-w-full max-h-full">
    {#await GetInvidiousApiInstances()}
        <p>Loading...</p>
    {:then instances} 
        <h2>Instance Configuration</h2>

        <p>Please select an instance from the list below or select 'Custom' and enter a custom url pointing towards an API enabled Invidious instance.</p>

        <div class="space-y-4">
            <SelectInput bind:selected={instance} bind:selectedIndex label="Select Instance" options={[{ display: "Custom", value: "" }, ...instances]}/>
            {#if selectedIndex === 0}
                <TextInput bind:value={instance} label="Custom Instance" type="url" errorMessage="Please enter a valid url" placeholder="https://example.org" validation={urlValidation}/>
            {/if}
        </div>

        <Button disabled={instance === ""} on:click={onInstanceConfirm}>Confirm</Button>
    {:catch error}
        <ErrorSpan>{error}</ErrorSpan>
    {/await}
</main>