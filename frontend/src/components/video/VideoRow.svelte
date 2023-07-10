<script>
    // components
    import { Icon } from "@steeze-ui/svelte-icon"
    import { ChevronLeft, ChevronRight } from "@steeze-ui/carbon-icons"


    // exports
    export let dataFunction
</script>


{#await dataFunction()}
    <p>Loading...</p>
{:then data}
    <Icon src={ChevronLeft} size="32"/>

    <div class="videosContainer">
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
{:catch error}
    <p class="errorText">{error}</p>
{/await}


<style lang="postcss">
    ::-webkit-scrollbar {
        display: none;
    }

    .videosContainer {
        @apply flex flex-row overflow-x-scroll space-x-4 rounded-4px min-h-fit;

        -ms-overflow-style: none;
        scrollbar-width: none;
    }

    .videoContainer {
        @apply min-w-fit h-fit rounded-lg -mb-[124px];
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