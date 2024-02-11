<script>
    import * as go from "../../wailsjs/go/main/Envious.js"
    import VideoGrid from "../components/VideoGrid.svelte";
    import { RotateCw as RotateCwIcon } from "@steeze-ui/lucide-icons"
    import { Icon } from "@steeze-ui/svelte-icon"

    let trendingType = "all"

    let videos = {
        all: [],
        music: [],
        gaming: [],
        movies: [],
    }

    function selectTrendingType(event) {
        trendingType = event.currentTarget.dataset.t
        if (videos[trendingType].length === 0) {
            loadCurrentTrendingVideos()
        }
    }

    function loadCurrentTrendingVideos() {
        // @ts-ignore
        videos[trendingType] = go.GetTrendingVideos({ type: trendingType })
    }

    loadCurrentTrendingVideos()
</script>


<menu class="menu"> 
    <div class="left">
        <div class="trendingTypeSelector"> 
            <button data-t="all" data-a={trendingType==="all"} on:click={selectTrendingType} class="item">All</button>
            <button data-t="music" data-a={trendingType==="music"} on:click={selectTrendingType} class="item">Music</button>
            <button data-t="gaming" data-a={trendingType==="gaming"} on:click={selectTrendingType} class="item">Gaming</button>
            <button data-t="movies" data-a={trendingType==="movies"} on:click={selectTrendingType} class="item">Movies</button>
        </div>
    </div>

    <div class="right">
        <button class="reload" title="Reload" on:click={loadCurrentTrendingVideos}>
            <Icon src={RotateCwIcon} stroke-width="3" size="18"></Icon>
        </button>
    </div>
</menu>

{#await videos[trendingType] then videos}
    <VideoGrid {videos}/>
{/await}



<style>
    .menu {
        display: flex;
        flex-direction: row;
        z-index: 100;
        margin-bottom: 20px;
        width: 100%;

        & > .left {
            margin-right: auto;

            & > .trendingTypeSelector {
                display: flex;
                flex-direction: row;
                background: rgba(255 255 255 / 0.05);
                border-radius: 999px;
                padding: 4px;
                gap: 4px;
                justify-content: space-between;
                box-shadow: 0 0 0.5rem 0.250rem rgba(0, 0, 0, 0.05);
                width: fit-content;

                & > .item {
                    width: 100%;
                    height: 30px;
                    text-align: center;
                    padding: 3px;
                    padding-left: 24px;
                    padding-right: 24px;
                    border-radius: 999px;
                    background: transparent;
                    transition-duration: 150ms;
                    cursor: pointer;
                    color: rgba(255 255 255 / 0.65);

                    &:hover {
                        background: rgba(255 255 255 / 0.05);
                        box-shadow: 0 0 0.313rem 0.05rem rgba(0, 0, 0, 0.15);
                    }

                    &[data-a=true] {
                        background: rgba(255 255 255 / 0.1);
                        box-shadow: 0 0 0.313rem 0.125rem rgba(0, 0, 0, 0.15);
                        cursor: default;
                        color: rgba(255 255 255 / 0.8)
                    }
                }
            }
        }

        & > .right {
            margin-left: auto;

            & > .reload {
                border-radius: 999px;
                padding: 10px;
                height: 38px;
                width: 38px;
                box-shadow: 0 0 0.5rem 0.250rem rgba(0, 0, 0, 0.05);
                background: rgba(255 255 255 / 0.05);
            }
        }
    }
</style>