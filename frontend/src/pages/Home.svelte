<script>
    import { GetSelectedInvidiousInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"

    import { link } from "svelte-spa-router"

    import ErrorSpan from "../components/text/ErrorSpan.svelte";
</script>

<main>
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

<style lang="postcss">
    main {
        @apply space-y-2;
    }
</style>