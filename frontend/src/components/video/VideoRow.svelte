<script>    
    import { spawnTab } from "../window/Window.svelte"
    
    import { Icon } from "@steeze-ui/svelte-icon"
    import { ChevronLeft, ChevronRight } from "@steeze-ui/carbon-icons"

    import Video from "../../pages/Video.svelte";


    export let data = []                        
    export let dataFunction = async () => data // allow for data to be directly entered into the component or use custom dataFunction


    let videosElement
    let showScrollLeft = false 
    let showScrollRight = true


    const calculateScrollIndicators = (_) => {
        showScrollLeft = videosElement.scrollLeft > 0
        showScrollRight = videosElement.scrollLeft < videosElement.scrollWidth - videosElement.clientWidth
    }


    const scrollLeft = (_) => {
        let offset = (videosElement.scrollLeft) % 336

        videosElement.scrollTo({
            top: 0,
            left: videosElement.scrollLeft - offset - (offset === 0 ? 336: 0),
            behavior: "smooth",
        })
    }


    const scrollRight = (_) => {
        let offset = (320 - (videosElement.scrollLeft + videosElement.clientWidth) % 336)

        videosElement.scrollTo({
            top: 0,
            left: videosElement.scrollLeft + offset + (offset === 0 ? 336: 0),
            behavior: "smooth",
        })
    }


    const openVideoInTab = (_, video) => {
        spawnTab({ 
            name: video.title,
            component: Video, 
            props: {
                videoId: video.id
            } 
        }, false)
    }
</script>


{#await dataFunction()}
    <p>Loading...</p>
{:then data}
    <div class="root">
        <div class="scrollIndicator left">
            <button class="icon" class:iconAllowed={showScrollLeft} on:click={scrollLeft}>
                <Icon src={ChevronLeft} size="32"/>
            </button>
        </div>

        <div bind:this={videosElement} class="videos" on:scroll={calculateScrollIndicators}>
            {#each data as video}
                <button class="video" title={video.title} on:click={(event) => openVideoInTab(event, video)}>
                    <img class="image" src={video.thumbnailUrl} alt={video.title}/>
                    <img class="textImage" src={video.thumbnailUrl} alt={video.title}/>
                    <div class="text">
                        <p class="title">{video.title}</p>
                        <a class="author" href="/author/{video.id}">{video.author}</a>
                    </div>
                </button>
            {/each}
        </div>

        <div class="scrollIndicator right">
            <button class="icon" class:iconAllowed={showScrollRight} on:click={scrollRight}>
                <Icon src={ChevronRight} size="32"/>
            </button>
        </div>
    </div>
{:catch error}
    <p class="error">{error}</p>
{/await}


<style lang="postcss">
    ::-webkit-scrollbar {
        display: none;
    }

    .root {
        @apply flex flex-row max-w-full w-full;
    }

    .error {
        @apply mb-3 text-red-600;
    }

    
    /* scroll indicators */
    .scrollIndicator {
        @apply absolute w-[32px] h-[304px] flex flex-col justify-center bg-transparent pointer-events-none;
    }

    .scrollIndicator.left {
        @apply left-[40px];
    }

    .scrollIndicator.right {
        @apply right-[56px];
    }

    .scrollIndicator > .icon {
        @apply z-10 rounded-full text-zinc-200 p-2 shadow-md opacity-0 duration-100 w-[48px] h-[48px] pointer-events-none bg-zinc-800;
    } 

    .scrollIndicator > .icon:hover {
        @apply brightness-125;
    }

    .root:hover .icon.iconAllowed { /* show icon when root is hovered over */
        @apply opacity-100 pointer-events-auto;
    }


    /* videos container and video */
    .videos {
        @apply flex flex-row overflow-x-scroll overflow-y-hidden space-x-4 rounded-lg min-h-fit mx-0;

        -ms-overflow-style: none;
        scrollbar-width: none;
    }

    .video {
        @apply min-w-fit h-fit rounded-lg -mb-[124px] duration-100;
    }

    .video:hover {
        @apply brightness-125;
    }

    .video > .image {
        @apply w-[320px] h-[180px] min-w-fit bg-black rounded-t-md;
    }

    .video > .textImage {
        @apply w-[320px] h-[124px] rounded-b-md bg-black;
    }

    .video > .text {
        @apply -translate-y-[124px] flex flex-col p-3 will-change-contents text-white space-y-1 rounded-b-md backdrop-blur-2xl saturate-200 bg-glassBlack;
    }

    .video > .text > .title {
        /* 12px + 296px + 12px = 320px (12px padding each side)*/
        @apply max-w-[296px] break-words font-semibold line-clamp-3 min-h-[72px] text-left;
    }

    .video > .text > .author {
        @apply text-left;
    }
</style>