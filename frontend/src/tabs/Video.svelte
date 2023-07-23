<script>
    import { BrowserOpenURL } from "../../wailsjs/runtime/runtime"
    import { GetVideo } from "../../wailsjs/go/main/Tubed";
    import { spawnTab } from "../components/window/Window.svelte";

    import VideoRow from "../components/video/VideoRow.svelte";

    import Video from "./Video.svelte"


    export let videoId


    let frame


    const onLoad = (_) => {
        setTimeout(() => {
            frame.classList.remove("unloaded")
        }, 1000)
    }


    const captureAnchors = (element) => {
        for (let child of element.children) {
            if (child.tagName === "A") {
                const href = child.href 
                child.href = ""

                const withoutProtocol = href.replace("https://", "").replace("http://", "")
                const isYoutube = withoutProtocol.startsWith("youtube.com") || withoutProtocol.startsWith("www.youtube.com")

                if (isYoutube) {
                    const videoId = withoutProtocol.split("watch?v=")[1]
                    if (videoId !== undefined) {
                        // use action that adds a click event listener
                        spawnTab(child, [
                            { 
                                name: href,
                                component: Video, 
                                props: {
                                    videoId: videoId
                                } 
                            }, 
                            true
                        ])
                    }
                } else {
                    child.addEventListener("click", (event) => {
                        event.preventDefault()
                        BrowserOpenURL(href)
                    })
                }
            }
        }
    }
</script>


{#await GetVideo(videoId)}
    <p>Loading...</p>
{:then video} 
    <div class="video">
        <div class="background">
            <iframe bind:this={frame} on:load={onLoad} allow="fullscreen" class="embed unloaded" title={video.title} src={video.embedUrl}/>
        </div>

        <p class="title">{video.title}</p>
        
        <div class="avatar">
            <img class="icon" alt={video.author} src={video.authorAvatarUrl}/>
            <p class="author">{video.author}</p>
        </div>
    </div>

    <details class="description">
        <summary>
            Click to view description
        </summary>
        <div use:captureAnchors>
            {@html video.description}
        </div>
    </details>

    <div class="divider"></div>

    <p>Recommended:</p>
    <VideoRow rawData={video.recommendedVideos}/>
{:catch error}
    <span class="error">{error}</span>
{/await}


<style lang="postcss">
    .embed {
        @apply w-full aspect-video rounded-lg duration-100;
    }

    .background {
        @apply bg-black rounded-lg;
    }

    .title {
        @apply font-semibold text-2xl my-2;
    }

    .avatar {
        @apply flex flex-row;
    }

    /* .description {
        @apply ;
    } */

    .icon {
        @apply w-9 rounded-full;
    }

    .author {
        @apply leading-9 ml-2;
    }

    .unloaded {
        @apply opacity-0;
    }

    .divider {
        @apply mb-4;
    }

    .error {
        @apply text-red-600;
    }
</style>