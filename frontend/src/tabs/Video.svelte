<script>
    import VideoGrid from "../components/VideoGrid.svelte";
    import VideoPlayer from "../components/VideoPlayer.svelte"
    import * as wails from "../../wailsjs/runtime/runtime.js"
    import * as go from "../../wailsjs/go/main/Envious.js";
    import { tabs } from "../Window.svelte"
    import Video from "../tabs/Video.svelte"

    export let videoId

    let descriptionElement

    const toggleDescription = (event) => {
        if (descriptionElement.dataset.c === "true") {
            event.target.innerText = "Read less..."
            descriptionElement.dataset.c = "false"
        } else {
            event.target.innerText = "Read more..."
            descriptionElement.dataset.c = "true"
        }
    } 

    const processDescription = (element) => {
        const anchors = Array.from(element.children).filter((element) => element.tagName === "A")
        for (let anchor of anchors) {
            let url = new URL(anchor.href)

            if (url.pathname === "/watch" && url.search.startsWith("?v=")) {
                let videoId = url.searchParams.get("v")

                go.GetVideo(videoId).then((video) => {
                    if (video.adaptiveFormats !== null) {
                        let img = new Image()
                        img.src = video.videoThumbnails.filter((v)=>v.quality==="medium")[0].url

                        anchor.addEventListener("click", (event) => {
                            console.log(url.searchParams.get("v") + " : : : ")
                            event.preventDefault()
                            tabs.create({
                                group: "Videos",
                                name: video.title,
                                component: Video,
                                props: { videoId: videoId },
                                active: true,
                                backgroundUrl: img.src,
                            })
                        })

                        anchor.innerText = "Watch " + video.title
                    } else {
                        anchor.innerText = "Failed to load video " + videoId
                    }
                })
            } else {
                if (url.hostname === "wails.localhost") {
                    url.protocol = "https"
                    url.host = "youtube.com"
                    url.port = "443"
                }

                anchor.addEventListener("click", (event) => {
                    event.preventDefault()
                    wails.BrowserOpenURL(url.toString())
                })

                anchor.innerText = url.toString()
            }
            
            anchor.href = "javascript:void(0)"
        }
    }
</script>


{#await go.GetVideo(videoId)}
    <p>Loading...</p>
{:then video}
    <div>
        <div class="w-full !rounded-t-md !rounded-b-none child:child:rounded-b-none bg-black/50 h-fit aspect-video">
            <VideoPlayer {video}/>
        </div>

        <div class="p-4 rounded-b-md bg-white/5">
            <p class="mb-4 text-2xl font-semibold">{video.title}</p>
        
            <button class="flex flex-row duration-100 hover:text-white/60">
                <img class="w-10 rounded-full" alt={video.author + " avatar"} src={video.authorThumbnails.filter((t)=>t.width===48)[0].url}/>
                
                <p class="ml-3 font-semibold leading-10">{video.author}</p>
            </button>

            <div class="overflow-hidden mt-4 word-break break-words w-fit decendant-type-[a]:duration-100 decendant-type-[a]:text-blue-400 hover:decendant-type-[a]:underline">
                <div bind:this={descriptionElement} data-c="true" class="data-[c=true]:line-clamp-3 truncate whitespace-normal" use:processDescription>
                    {@html video.descriptionHtml}
                </div>

                <button on:click={toggleDescription} class="mt-1 duration-100 text-white/60 hover:text-white/75">Read more...</button>
            </div>
        </div>
    </div>

    {#if video.recommendedVideos !== null}
        <p class="mt-4 mb-1 text-white/60">Other video's you might like:</p>
        <VideoGrid videos={video.recommendedVideos}/>
    {/if}
{/await}