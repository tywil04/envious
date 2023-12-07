<script>    
    import Video from "../tabs/Video.svelte";

    import tabSystem from "../window/tabSystem.js"


    export let videos = []                        


    const openVideoTab = async (video) => {
        tabSystem.createTab({
            group: "Videos",
            name: video.title,
            component: Video,
            props: { videoId: video.videoId },
            active: true,
            // backgroundUrl: getLowestQualityThumbnail(video.videoThumbnails).url,
            backgroundUrl: video.thumbnailImg.currentSrc,
        })
    }

    function getLowestQualityThumbnail(videoThumbnails) {
        let lowestQuality = null
        for (let thumbnail of videoThumbnails) {
            if (lowestQuality === null || thumbnail.width < lowestQuality.width) {
                lowestQuality = thumbnail
            }
        }
        return lowestQuality
    }
</script>


<div class="grid grid-cols-1 gap-4 gap-y-6 video-grid-1:grid-cols-1 video-grid-2:grid-cols-2 video-grid-3:grid-cols-3 video-grid-4:grid-cols-4 video-grid-5:grid-cols-5 video-grid-6:grid-cols-6">
    {#each videos as video}
        <button class="flex flex-col hover:scale-[1.025] duration-100" title={video.title} on:click={() => openVideoTab(video)}>
            <picture class="w-full h-full">
                {#each video.videoThumbnails as thumbnail}
                    {#if (thumbnail.width/16)*9 == thumbnail.height}
                        <source media={`(min-width: ${thumbnail.width}px)`} srcset={"/" + thumbnail.url}/>
                    {/if}
                {/each}
                <img bind:this={video.thumbnailImg} class="object-cover w-full h-full rounded-t-md" alt={video.title}/>
            </picture>
            
            <div class="flex flex-col p-2 w-full h-full rounded-b-md bg-white/5">
                <p class="font-semibold text-left line-clamp-2">{video.title}</p>
                <a class="pt-2 mt-auto text-left line-clamp-1 text-white/50" href="/author/{video.id}">{video.author}</a>
            </div>
        </button>
    {/each}
</div>
