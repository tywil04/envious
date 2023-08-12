<script>
    import { BrowserOpenURL } from "../../wailsjs/runtime/runtime.js"
    import { GetVideo } from "../../wailsjs/go/main/Tubed.js";
    import { spawnTab } from "../components/Window.svelte";

    import VideoRow from "../components/VideoRow.svelte";

    import Video from "./Video.svelte"


    export let videoId


    let frame
    let descriptionElement


    const toggleDescription = (event) => {
        if (event.target.innerText === "Read more...") {
            event.target.innerText = "Read less..."
            descriptionElement.classList.remove("clamped")
        } else {
            event.target.innerText = "Read more..."
            descriptionElement.classList.add("clamped")
        }
    }   


    const loadEmbed = (_) => {
        setTimeout(() => {
            frame.classList.remove("unloaded")
        }, 1000)
    }


    const processDescription = (element) => {
        for (let child of element.children) {
            if (child.tagName === "A") {
                const href = child.href 
                child.href = "javascript:void(0)"

                const withoutProtocol = href.replace("https://", "").replace("http://", "")
                const isYouTubeVideo = withoutProtocol.startsWith("youtube.com/watch?v=") || withoutProtocol.startsWith("www.youtube.com/watch?v=")

                if (isYouTubeVideo) {
                    const urlArguments = withoutProtocol.split("watch?")[1].split("&")

                    let hasTimestamp
                    let id
                    for (let argument of urlArguments) {
                        let split = argument.split("=")
                        if (split[0] === "v") {
                            id = split[1]
                        }

                        if (split[0] === "t") {
                            hasTimestamp = true
                        }
                    }

                    if (id !== undefined && videoId != id) {
                        GetVideo(id).then((video) => {
                            child.addEventListener("click", (event) => {
                                event.preventDefault()
                                spawnTab({ 
                                    name: video.title,
                                    component: Video, 
                                    props: {
                                        videoId: video.id
                                    } 
                                }, false)
                            })
                            child.classList.add("link")
                            child.innerText = "Watch " + video.title
                        })
                    } else if (hasTimestamp && videoId == id) {
                        child.addEventListener("click", (event) => event.preventDefault())
                        child.classList.add("disabledLink")
                    }
                } else {
                    child.addEventListener("click", (event) => {
                        event.preventDefault()
                        BrowserOpenURL(href)
                    })
                    child.classList.add("link")
                    child.innerText = href
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
            <iframe bind:this={frame} on:load={loadEmbed} allow="fullscreen" class="embed unloaded" title={video.title} src={video.embedUrl}/>
        </div>

        <div class="info">
            <p class="title">{video.title}</p>
        
            <div class="author">
                <img class="avatar" alt={video.author} src={video.authorAvatarUrl}/>
                <p class="name">{video.author}</p>
            </div>
    
            <p class="descriptionLabel">Description:</p>
            <div bind:this={descriptionElement} class="description clamped" use:processDescription>
                {@html video.description}
            </div>
            <button on:click={toggleDescription} class="descriptionToggle">Read more...</button>
        </div>
    </div>

    <p class="label">Other video's you might like:</p>
    <VideoRow rawData={video.recommendedVideos}/>
{:catch error}
    <span class="error">{error}</span>
{/await}


<style lang="postcss">
    /* for processDescription */
    :global(.link) {
        @apply duration-100 text-blue-400;
    }

    :global(.disabledLink) {
        @apply pointer-events-none;
    }

    :global(.link:hover) {
        @apply underline;
    }


    .label {
        @apply mt-4 mb-1 text-zinc-400;
    }

    .error {
        @apply text-red-600;
    }


    .video > .background {
        @apply bg-black rounded-lg flex flex-row justify-center;
    }

    .video > .background > .embed {
        @apply w-full aspect-video rounded-t-lg duration-100;
    }

    .video > .background > .embed.unloaded {
        @apply opacity-0;
    }

    .video > .info {
        @apply bg-zinc-800 p-4 rounded-b-lg;
    }

    .video > .info > .title {
        @apply font-semibold text-2xl mb-4;
    }

    .video > .info > .author {
        @apply flex flex-row;
    }

    .video > .info > .author > .avatar {
        @apply w-9 rounded-full;
    }

    .video > .info > .author > .name {
        @apply leading-9 ml-3 font-semibold;
    }

    .video > .info > .descriptionLabel {
        @apply text-zinc-400 mt-4;
    }

    .video > .info > .description {
        @apply text-zinc-300/90;
    }

    .video > .info > .description.clamped {
        @apply line-clamp-3;
    }

    .video > .info > .descriptionToggle {
        @apply duration-100 mt-1;
    }

    .video > .info > .descriptionToggle:hover {
        @apply text-zinc-200;
    }
</style>