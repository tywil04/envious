<script>
    import { GetTrendingVideos } from "../../wailsjs/go/main/Tubed.js"

    import VideoGrid from "../components/VideoGrid.svelte";


    export let selectedSubTab
</script>


<h1 class="mb-2 text-lg font-semibold">Trending {selectedSubTab || ""}</h1>

{#await GetTrendingVideos({ type: selectedSubTab }) then videos}
    <VideoGrid {videos}/>
{/await}