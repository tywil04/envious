<script>
    import { GetTrending, GetInstanceApi } from "../../wailsjs/go/main/Tubed"

    import VideoRow from "../components/video/VideoRow.svelte";
</script>

    
{#await GetInstanceApi()}
    <p>Loading...</p>
{:then instance}
    <div class="root">
        {#if instance !== ""}
            <p>Your selected instance is: {instance}.</p>

            <div class="divider"/>
        
            <p>Trending:</p>
            <VideoRow dataFunction={GetTrending}/>
        {:else}
            <p>You dont have a selected instance.</p>
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