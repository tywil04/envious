<script>
    import { WindowFullscreen, WindowUnfullscreen } from "../../wailsjs/runtime/runtime"

    import { GetVideo } from "../../wailsjs/go/main/Tubed.js";


    export let videoId
</script>


{#await GetVideo(videoId)}
    <p>Loading...</p>
{:then video} 
    <iframe allow="fullscreen" class="video" title={video.title} src={video.embedUrl}/>
{:catch error}
    <span class="error">{error}</span>
{/await}


<style lang="postcss">
    .video {
        @apply w-full aspect-video rounded-lg;
    }

    :global(div.vjs-overlay.vjs-overlay-top.vjs-overlay-background > h1) {
        display: none;
    }

    .error {
        @apply text-red-600;
    }
</style>