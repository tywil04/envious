<script>
    import { GetSelectedInvidiousInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"

    import { link } from "svelte-spa-router"

    import ErrorSpan from "../components/text/ErrorSpan.svelte";
</script>

<main>
    <h2>Instance Configuration</h2>

    <a href="/instanceSetup" use:link>Set or Change Invidious instance used.</a>

    {#await GetSelectedInvidiousInstance()}
        <p>Loading...</p>
    {:then instance} 
        {#if instance !== ""}
            <p>Your selected instance is: {instance}.</p>
        {:else}
            <p>You dont have a selected instance.</p>
        {/if}
    {:catch error}
        <ErrorSpan>{error}</ErrorSpan>
    {/await}
</main>