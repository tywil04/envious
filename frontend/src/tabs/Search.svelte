<script>
    import VideoGrid from "../components/VideoGrid.svelte";
    import * as go from "../../wailsjs/go/main/Tubed";


    let results = null


    async function search(event) {
        console.log("s  ")
        results = await go.Search(event.currentTarget.value, { type: "video" })
        console.log(results)
    }
</script>


<input on:input={search} class="text-white bg-black"/>

{#if results !== null}
    <VideoGrid videos={results}/>
{/if}