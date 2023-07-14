<script>    
    // components
    import { Icon } from "@steeze-ui/svelte-icon"
    import { ChevronLeft, ChevronRight } from "@steeze-ui/carbon-icons"


    // exports
    export let dataFunction


    let videoContainer
    let showScrollLeft = false 
    let showScrollRight = true

    const onScroll = (event) => {
        showScrollLeft = videoContainer.scrollLeft > 0
        showScrollRight = videoContainer.scrollLeft < videoContainer.scrollWidth - videoContainer.clientWidth
    }

    const scrollLeft = (event) => {
        let newX
        let calc1 = (videoContainer.scrollLeft) % 336
        
        if (calc1 === 0) {
            newX = videoContainer.scrollLeft - calc1 - 336
        } else {
            newX = videoContainer.scrollLeft - calc1
        }

        videoContainer.scrollTo({
            top: 0,
            left: newX,
            behavior: "smooth",
        })
    }

    const scrollRight = (event) => {
        let newX
        let calc1 = (320 - (videoContainer.scrollLeft + videoContainer.clientWidth) % 336)

        if (calc1 === 0) {
            newX = videoContainer.scrollLeft + calc1 + 336
        } else {
            newX = videoContainer.scrollLeft + calc1
        }

        videoContainer.scrollTo({
            top: 0,
            left: newX,
            behavior: "smooth",
        })
    }
</script>


{#await dataFunction()}
    <p>Loading...</p>
{:then data}
    <div class="container">
        <div class="iconContainerLeft">
            <button class="icon" class:iconAllowed={showScrollLeft} on:click={scrollLeft}>
                <Icon src={ChevronLeft} size="32"/>
            </button>
        </div>

        <div bind:this={videoContainer} class="videosContainer" on:scroll={onScroll}>
            {#each data as video}
                <a class="videoContainer" href="/video/{video.videoId}" title={video.title}>
                    <img class="videoImage" src={video.thumbnailUrl} alt={video.title}/>
                    <img class="videoTextContainerImage" src={video.thumbnailUrl} alt={video.title}/>
                    <div class="videoTextContainer">
                        <p class="videoTitle">{video.title}</p>
                        <a href="/author/{video.videoId}">{video.author}</a>
                    </div>
                </a>
            {/each}
        </div>

        <div class="iconContainerRight">
            <button class="icon" class:iconAllowed={showScrollRight} on:click={scrollRight}>
                <Icon src={ChevronRight} size="32"/>
            </button>
        </div>
    </div>
{:catch error}
    <p class="errorText">{error}</p>
{/await}


<style lang="postcss">
    .container {
        @apply flex flex-row max-w-full w-full drop-shadow-lg;
    }

    .iconContainerLeft {
        @apply absolute w-[32px] h-[304px] left-[40px] flex flex-col justify-center bg-transparent pointer-events-none;
    }

    .iconContainerRight {
        @apply absolute w-[32px] h-[304px] right-[56px] flex flex-col justify-center bg-transparent pointer-events-none;
    }

    .icon {
        @apply z-10 rounded-full bg-gray-200 text-neutral-950 p-2 drop-shadow-2xl opacity-0 duration-100 w-[48px] h-[48px] pointer-events-none;
    } 

    .container:hover .icon.iconAllowed {
        @apply opacity-100 pointer-events-auto;
    }

    ::-webkit-scrollbar {
        display: none;
    }

    .videosContainer {
        @apply flex flex-row overflow-x-scroll overflow-y-hidden space-x-4 rounded-lg min-h-fit mx-0;

        -ms-overflow-style: none;
        scrollbar-width: none;
    }

    .videoContainer {
        @apply min-w-fit h-fit rounded-lg -mb-[124px] duration-100;
    }

    .videoContainer:hover {
        @apply brightness-125;
    }

    .videoImage {
        @apply w-[320px] h-[180px] min-w-fit bg-black rounded-t-md;
    }

    .videoTextContainerImage {
        @apply w-[320px] h-[124px] rounded-b-md bg-black;
    }

    .videoTextContainer {
        @apply -translate-y-[124px] flex flex-col p-3 will-change-contents text-white space-y-1 rounded-b-md backdrop-blur-2xl saturate-200 bg-glassBlack;
    }

    .videoTitle {
        /* 12px + 296px + 12px = 320px (12px padding each side)*/
        @apply max-w-[296px] break-words font-semibold line-clamp-3 min-h-[72px];
    }


    .errorText {
        @apply mb-3 text-red-600;
    }
</style>