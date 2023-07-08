<script>
    import { GetSelectedInvidiousInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"

    import { link } from "svelte-spa-router"

    import ErrorSpan from "../components/text/ErrorSpan.svelte";
</script>

<p>Hi</p>

<a href="/instanceSetup" use:link>To Instance Setup</a>

{#await GetSelectedInvidiousInstance()}
    <p>Loading...</p>
{:then instance} 
    {#if instance !== ""}
        <p>{instance}</p>
    {:else}
        <p>No instance :(</p>
    {/if}
{:catch error}
    <ErrorSpan>{error}</ErrorSpan>
{/await}