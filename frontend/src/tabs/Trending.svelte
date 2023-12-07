<script>
    import { GetTrendingVideos } from "../../wailsjs/go/main/Tubed.js"

    import VideoGrid from "../components/VideoGrid.svelte";
</script>

    
<p class="label">Trending:</p>

{#await GetTrendingVideos() then trendingVideos}
    <VideoGrid videos={trendingVideos}/>
{/await}


<style lang="postcss">
    .label {
        @apply mt-4 mb-1 text-zinc-400;
    }
</style>