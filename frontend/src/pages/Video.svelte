<script>
    import { GetVideo } from "../../wailsjs/go/main/InvidiousDesktop";

    
    export let videoId
</script>


{#await GetVideo(videoId)}
    <p>Loading...</p>
{:then video} 
    <iframe allow="fullscreen" class="videoEmbed" title={video.title} src={video.embedUrl}/>
{:catch error}
    <span class="errorText">{error}</span>
{/await}


<style lang="postcss">
    .videoEmbed {
        @apply w-full aspect-video rounded-lg;
    }

    :global(div.vjs-overlay.vjs-overlay-top.vjs-overlay-background > h1) {
        display: none;
    }
</style>