<script>
    // exports
    export let dataFunction
</script>


{#await dataFunction()}
    <p>Loading...</p>
{:then data}
    <div class="videosContainer">
        {#each data as video}
            <article class="videoContainer">
                <img class="video" loading="lazy" src={video.thumbnailUrl} alt={video.title} width={320} height={180}/>
                <a href="https://vid.puffyan.us/watch?v={video.videoId}" class="title">{video.title}</a>
            </article>
        {/each}
    </div>
{:catch error}
    <p class="errorText">{error}</p>
{/await}


<style lang="postcss">
    .videosContainer {
        @apply flex flex-row overflow-x-scroll space-x-4;
    }

    .videoContainer {
        @apply min-w-fit;
    }


    .video {
        @apply w-[320px] h-[180px] min-w-fit rounded-4px bg-black;
    }

    .title {
        @apply max-w-[320px];
    }


    .errorText {
        @apply mb-3 text-red-600;
    }
</style>