<script>
    import * as go from "../../wailsjs/go/main/Tubed.js"

    import VideoGrid from "../components/VideoGrid.svelte";


    let selected
</script>


<h1 class="mb-2 text-lg font-semibold">Trending {selected || ""}</h1>


<input type="radio" name="selectedType" value="All" on:input={(e)=>selected=undefined}>
<input type="radio" name="selectedType" value="Music" on:input={(e)=>selected=e.currentTarget.value}>
<input type="radio" name="selectedType" value="Gaming" on:input={(e)=>selected=e.currentTarget.value}>
<input type="radio" name="selectedType" value="Movies" on:input={(e)=>selected=e.currentTarget.value}>


{#await go.GetTrendingVideos({ type: selected }) then videos}
    <VideoGrid {videos}/>
{/await}