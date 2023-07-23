<script>    
    import { spawnTab } from "../window/Window.svelte"
    import * as StackBlur from "../../../node_modules/stackblur-canvas/dist/stackblur-es.min.js";

    import { Icon } from "@steeze-ui/svelte-icon"
    import { ChevronLeft, ChevronRight } from "@steeze-ui/carbon-icons"

    import Video from "../../tabs/Video.svelte";


    export let rawData = []                        
    export let dataFunction = async () => rawData // allow for data to be directly entered into the component or use custom dataFunction


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


    const blurBehindText = (element) => {
        let thumbnail
        let canvas

        for (let child of element.children) {
            if (child.tagName === "IMG") {
                thumbnail = child
                continue
            }

            if (child.tagName === "CANVAS") {
                canvas = child
            }
        }

        if (thumbnail != undefined && canvas != undefined) { 
            canvas.getContext("2d").drawImage(thumbnail, 0, 0, 320, 150)
            StackBlur.canvasRGB(canvas, 0, 0, 320, 150, 100)
        }
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
                <button class="video" title={video.title} use:blurBehindText use:spawnTab={[
                    { 
                        name: video.title,
                        component: Video, 
                        props: {
                            videoId: video.id
                        } 
                    }, 
                    false
                ]}>
                    <img crossorigin="anonymous" class="image" src={video.thumbnailUrl} alt={video.title}/>
                    <canvas id="canvas" class="textImage"/>
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

    .icon {
        @apply z-10 rounded-full text-zinc-200 p-2 shadow-md opacity-0 duration-100 w-[48px] h-[48px] pointer-events-none bg-zinc-800;
    } 

    .icon:hover {
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

    .image {
        @apply w-[320px] h-[180px] min-w-fit bg-black rounded-t-md;
    }

    .textImage {
        @apply w-[320px] h-[124px] rounded-b-md bg-black;
    }

    .text {
        @apply -translate-y-[124px] flex flex-col p-3 text-white space-y-1 rounded-b-md saturate-200 bg-glassBlack;
    }

    .title {
        /* 12px + 296px + 12px = 320px (12px padding each side)*/
        @apply max-w-[296px] break-words font-semibold line-clamp-3 min-h-[72px] text-left;
    }

    .author {
        @apply text-left;
    }
</style>