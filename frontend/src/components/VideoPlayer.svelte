<script>
    import { Play as PlayIcon, Pause as PauseIcon, SpeakerWave as SpeakerWaveIcon, SpeakerXMark as SpeakerXMarkIcon } from "@steeze-ui/heroicons";
    import { Maximize as MaximiseIcon, Minimize as MinimizeIcon } from "@steeze-ui/lucide-icons";
    import { Icon } from "@steeze-ui/svelte-icon";
    import dashjs from "dashjs";
    import { WindowUnfullscreen, WindowFullscreen } from "../../wailsjs/runtime/runtime.js";


    export let video
    

    let playerContainer = null
    let playerHtml = null
    let playerDash = null
    let controlsContainer = null

    let temporarilyShowControlsTimeout = null

    let playing = false 
    let muted = false 
    let fullscreen = false
    let appFullscreen = false

    let volume = 100
    let currentTime = 0
    let duration = 1
    let watchedPercentage = 0

    let quality = "dash"

    
    $: watchedPercentage = (currentTime / duration) * 100


    function loadPlayer(element, [quality, video]) {
        const actions = {
            update([quality]) {
                if (quality === "dash") {
                    if (playerDash === null) {
                        playerDash = dashjs.MediaPlayer().create()
                        playerDash.initialize(element, video.dashUrl, false, currentTime)
                        playerDash.updateSettings({
                            streaming: {
                                abr: {
                                    minBitrate: {
                                        video: 2000,
                                    },
                                    limitBitrateByPortal: true,
                                }
                            }
                        })
                        playerDash.setQualityFor("video", 5, true)
                        playerDash.setVolume(volume / 100)   
                    }
                } else {
                    if (playerDash !== null) {
                        playerDash.destroy()
                        playerDash = null
                    }

                    let selectedStream = null
                    for (let stream of video.formatStreams) {
                        if (stream.qualityLabel === quality) {
                            selectedStream = stream
                            break
                        }
                    }

                    playerHtml.src = selectedStream.url 
                    playerHtml.type = selectedStream.type
                }
            },

            destroy() {
                if (playerDash !== null) {
                    playerDash.destroy()
                    playerDash = null
                }
            }
        }

        actions.update([quality])
        return actions
    }

    function play() {
        if (quality === "dash") {
            playerDash.play()
        } else {
            playerHtml.play()
        }
        playing = true
    }

    function pause() {
        if (quality === "dash") {
            playerDash.pause()
        } else {
            playerHtml.pause()
        }
        playing = false
    }

    function togglePlay() {
        if (playing) {
            pause()
        } else {
            play()
        }
    }

    function setVolume(newVolume) {
        if (quality === "dash") {
            playerDash.setVolume(newVolume / 100)
        } else {
            playerHtml.volume = newVolume / 100
        }
        volume = newVolume
    }

    function mute() {
        playerHtml.muted = true 
        muted = true
    }

    function unmute() {
        playerHtml.muted = false 
        muted = false
    }

    function toggleMute() {
        if (muted) {
            unmute()
        } else {
            mute()
        }
    }

    function enterFullscreen() {
        if (document.fullscreenElement === null) {
            WindowFullscreen()
            appFullscreen = true

            playerContainer.requestFullscreen().then(() => {
                if (quality === "dash") {
                    playerDash.updatePortalSize()
                }
            })

            fullscreen = true
        }
    }

    function exitFullscreen() {
        if (document.fullscreenElement === playerContainer) {
            document.exitFullscreen()
        }
    }

    function playerToggleFullscreen() {
        if (fullscreen) {
            exitFullscreen()
        } else {
            enterFullscreen()
        }
    }

    function showControls() {
        controlsContainer.style.opacity = "100%"
        controlsContainer.style.cursor = "auto"
    }

    function hideControls() {
        if (playing) {
            controlsContainer.style.opacity = "0%"
            controlsContainer.style.cursor = "none"
        }
    }

    function temporarilyShowControls() {
        if (temporarilyShowControlsTimeout !== null) {
            clearTimeout(temporarilyShowControlsTimeout)
        }
        showControls()
        temporarilyShowControlsTimeout = setTimeout(hideControls, 3000)
    }    

    
    document.addEventListener("fullscreenchange", (event) => {
        if (event.target === playerContainer && document.fullscreenElement === null) {
            if (quality === "dash") {
                playerDash.updatePortalSize()
            }
            fullscreen = false

            if (appFullscreen === true) {
                WindowUnfullscreen()
                appFullscreen = false
            }
        }
    })
</script>


<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div bind:this={playerContainer} on:click={togglePlay} class="flex relative flex-col rounded-md cursor-default aspect-video">
    <!-- svelte-ignore a11y-media-has-caption -->
    <video on:ended={pause} bind:this={playerHtml} bind:duration={duration} bind:currentTime={currentTime} use:loadPlayer={[quality, video]} class="my-auto w-full rounded-md">
        {#each video.captions as caption}
            <track kind="captions" label={caption.label} src={caption.url} srclang={caption.languageCode}>
        {/each}
    </video>

    <div bind:this={controlsContainer} class="flex absolute top-0 flex-col w-full h-full duration-[400ms]" style={(playing || fullscreen) && "opacity: 0% important;"} on:mouseleave={hideControls} on:mousemove={temporarilyShowControls}>
        {#if playerHtml !== null}
            <div class="p-1.5 mt-auto w-full h-22 child:w-full">
                <div class="flex flex-row gap-2 p-1 mt-9 h-11 child:bg-black/70 child:backdrop-blur-[100px] child:p-2 child:rounded-md hover:child:text-white/60 child:text-white/95">
                    <button on:click|stopPropagation={togglePlay}>
                        <Icon src={playing ? PauseIcon : PlayIcon} theme="mini" size="20"></Icon>
                    </button>

                    <div class="!px-2.5 w-full">
                        <div style={`background: linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) ${watchedPercentage}%, rgb(255 255 255 / 0.6) ${watchedPercentage}%, rgb(255 255 255 / 0.6) 100%);`} class="absolute w-[calc(100%-20px)] top-[calc(50%-3px)] h-[6px] rounded-md"></div>
                        <input class="w-full opacity-0 cursor-pointer" type="range" min="0" max={duration} bind:value={currentTime} on:click|stopPropagation/>
                    </div>

                    <div class="flex flex-row gap-1 !pr-2.5 w-fit hover:!text-white/95">
                        <button class="mr-1 duration-100 hover:text-white/60" on:click|stopPropagation={toggleMute}>
                            <Icon src={muted || volume <= 0 ? SpeakerXMarkIcon : SpeakerWaveIcon} theme="mini" size="20"></Icon> 
                        </button>
                        <div style={`background: linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) ${volume}%, rgb(255 255 255 / 0.6) ${volume}%, rgb(255 255 255 / 0.6) 100%);`} class="absolute left-[36px] top-[calc(50%-3px)] w-16 h-[6px] rounded-md"></div>
                        <input class="w-16 opacity-0 cursor-pointer" type="range" height="10" min="0" max="100" value={volume} on:input|stopPropagation={(e)=>setVolume(e.currentTarget.value)} on:click|stopPropagation/>
                    </div>

                    <select on:click|stopPropagation bind:value={quality} class="w-fit cursor-pointer text-sm font-bold leading-5 appearance-none !border-none !outline-none !ring-0">
                        <option value="dash" selected>Quality: Dash</option>
                        {#each video.formatStreams as stream}
                            {#if playerHtml.canPlayType(stream.type) === "probably"}
                                <option value={stream.qualityLabel}>Quality: {stream.qualityLabel}</option>
                            {/if}
                        {/each}
                    </select>

                    <button data-b on:click|stopPropagation={playerToggleFullscreen}>
                        <Icon src={fullscreen ? MinimizeIcon : MaximiseIcon} stroke-width="3" size="20"></Icon>
                    </button>
                </div>
            </div>
        {/if}
    </div> 
</div>