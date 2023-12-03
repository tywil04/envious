<script>    
    import Video from "../tabs/Video.svelte";

    import tabSystem from "../window/tabSystem.js"
    import adaptiveBackground from "../window/adaptiveBackground";


    export let rawData = []                        
    export let dataFunction = async () => rawData // allow for data to be directly entered into the component or use custom dataFunction

    
    const openVideoTab = (video) => {
        tabSystem.createTab({
            group: "Videos",
            name: video.title,
            component: Video,
            props: { videoId: video.id },
            active: true,
            backgroundUrl: video.thumbnailUrl,
        })
    }
</script>


{#await dataFunction()}
    <p>Loading...</p>
{:then data}
    <div class="grid grid-cols-1 gap-4 gap-y-6 video-grid-1:grid-cols-1 video-grid-2:grid-cols-2 video-grid-3:grid-cols-3 video-grid-4:grid-cols-4 video-grid-5:grid-cols-5 video-grid-6:grid-cols-6">
        {#each data as video}
            <button class="flex flex-col hover:scale-[1.025] duration-100" title={video.title} on:click={() => openVideoTab(video)}>
                <img class="object-cover w-full h-full rounded-t-md" crossorigin="anonymous" src={video.thumbnailUrl} alt={video.title}/>
                <div class="flex flex-col p-2 w-full h-full rounded-b-md bg-white/5">
                    <p class="font-semibold text-left line-clamp-2">{video.title}</p>
                    <a class="pt-2 mt-auto text-left line-clamp-1 text-white/50" href="/author/{video.id}">{video.author}</a>
                </div>
            </button>
        {/each}
    </div>
{:catch error}
    <p class="error">{error}</p>
{/await}
 