<script>
    // javascript
    import { GetPopular, GetSelectedInvidiousInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"
    import { link } from "svelte-spa-router"


    // components
    import WindowShell from "../components/window/WindowShell.svelte";
    import VideoRow from "../components/video/VideoRow.svelte";
</script>


<WindowShell>
    <a href="/instanceSetup" use:link>Set or Change Invidious instance used.</a>
    
    {#await GetSelectedInvidiousInstance()}
        <p>Loading...</p>
    {:then instance} 
        {#if instance !== ""}
            <p>Your selected instance is: {instance}.</p>
        {:else}
            <p>You dont have a selected instance.</p>
        {/if}

        <VideoRow dataFunction={GetPopular}/>
    {:catch error}
        <span class="errorText">{error}</span>
    {/await}
</WindowShell>


<style lang="postcss">
    .errorText {
        @apply text-red-600;
    }
</style>