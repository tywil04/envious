<script>
    import { GetTrending, GetInstanceApi } from "../../wailsjs/go/main/Tubed.js"
    import { spawnTab } from "../components/window/Window.svelte"

    import VideoRow from "../components/video/VideoRow.svelte";
    import Settings from "./Settings.svelte";
</script>

    
{#await GetInstanceApi()}
    <p>Loading...</p>
{:then instance}
    <div class="root">
        {#if instance !== ""}
            <p>Your selected instance is: {instance}.</p>

            <button use:spawnTab={[
                { 
                    name: "Settings",
                    component: Settings, 
                }, 
                false
            ]}>
                Goto settings to select an instance now
            </button>

            <div class="divider"/>
        
            <p>Trending:</p>
            <VideoRow dataFunction={GetTrending}/>
        {:else}
            <p>You dont have a selected instance.</p>

            <button use:spawnTab={[
                { 
                    name: "Settings",
                    component: Settings, 
                }, 
                false
            ]}>
                Goto settings to select an instance now
            </button>
        {/if}
    </div>
{:catch error}
    <span class="error">{error}</span>
{/await}


<style lang="postcss">
    .root {
        @apply flex flex-col max-w-full;
    }

    .divider {
        @apply mb-4;
    }

    .error {
        @apply text-red-600;
    }
</style>