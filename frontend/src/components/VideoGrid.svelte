<script>    
    import Video from "../tabs/Video.svelte";

    import { tabs } from "../Window.svelte"


    export let videos = []                        


    const openVideoTab = (video) => {
        tabs.create({
            group: "Videos",
            name: video.title,
            component: Video,
            props: { videoId: video.videoId },
            active: true,
            backgroundUrl: video.thumbnailImg.src,
        })
    }
</script>


<div class="grid grid-cols-1 gap-4 gap-y-6 video-grid-1:grid-cols-1 video-grid-2:grid-cols-2 video-grid-3:grid-cols-3 video-grid-4:grid-cols-4 video-grid-5:grid-cols-5 video-grid-6:grid-cols-6">
    {#each videos as video}
        <button class="flex flex-col hover:scale-[1.025] duration-100" title={video.title} on:click={() => openVideoTab(video)}>
            <img bind:this={video.thumbnailImg} src={video.videoThumbnails.filter((v)=>v.quality==="medium")[0].url} class="w-full h-full rounded-t-md aspect-video" alt={video.title}/>
            <div class="flex flex-col p-3 w-full h-full rounded-b-md bg-white/5">
                <p class="font-semibold text-left line-clamp-2">{video.title}</p>
                <button on:click|stopPropagation class="pt-2 mt-auto text-left duration-100 line-clamp-1 text-white/60">{video.author}</button>
            </div>
        </button>
    {/each}
</div>
