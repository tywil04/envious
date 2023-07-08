<script>
    // javascript
    import { GetSelectedInvidiousInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"
    import { pushRouteAction } from "../lib/router.js"


    // components
    import WindowShell from "../components/window/WindowShell.svelte";
</script>


<WindowShell>
    <button use:pushRouteAction={"instanceSetup"}>Set or Change Invidious instance used.</button>
    
    {#await GetSelectedInvidiousInstance()}
        <p>Loading...</p>
    {:then instance} 
        {#if instance !== ""}
            <p>Your selected instance is: {instance}.</p>
        {:else}
            <p>You dont have a selected instance.</p>
        {/if}
    {:catch error}
        <span class="errorText">{error}</span>
    {/await}
</WindowShell>


<style lang="postcss">
    .errorText {
        @apply text-sm text-red-600;
    }
</style>