<script context="module">
    export class player {
        static container = null 
        static controls = null 
        static playbackDuration = 1
        static playbackPosition = 0
        static video = null
        
        static #htmlPlayer = null 
        static #dashPlayer = null 
        static #controlsTimeout = null 
        static #controlsVisible = false
        static #playing = false 
        static #muted = false
        static #fullscreen = false 
        static #volume = 100        
        static #quality = null
        static #loaded = false

        static init(containerElement, videoElement, controlsElement, video, quality="dash") {
            this.#htmlPlayer = videoElement
            this.container = containerElement
            this.controls = controlsElement
            this.video = video
            this.quality = quality
            this.#loaded = true
        }

        static get loaded() {
            return this.#loaded
        }

        static get quality() {
            return this.#quality
        }

        static set quality(quality) {
            if (quality === "dash") {
                if (this.#dashPlayer === null) {
                    this.#dashPlayer = dashjs.MediaPlayer().create()
                    this.#dashPlayer.initialize(this.#htmlPlayer, this.video.dashUrl, false, this.playbackPosition)
                    this.#dashPlayer.updateSettings({
                        streaming: {
                            abr: {
                                minBitrate: {
                                    video: 2000,
                                },
                                limitBitrateByPortal: true,
                            }
                        }
                    })
                    this.#dashPlayer.setQualityFor("video", 5, true)
                    this.#dashPlayer.setVolume(this.volume / 100)   
                }

                this.#quality = quality
            } else {
                if (this.#dashPlayer !== null) {
                    this.#dashPlayer.destroy()
                    this.#dashPlayer = null
                }

                let selectedStream = null
                for (let stream of this.video.formatStreams) {
                    if (this.#htmlPlayer.canPlayType(stream.type) === "probably") {
                        if (stream.qualityLabel === quality) {
                            selectedStream = stream
                            break
                        }
                    }
                }

                this.#htmlPlayer.src = selectedStream.url 
                this.#htmlPlayer.type = selectedStream.type
                this.#quality = selectedStream.qualityLabel
            }
        }

        static get playing() { 
            return this.#playing
        }

        static set playing(playing) {
            let player = this.#htmlPlayer
            if (this.#quality === "dash") {
                player = this.#dashPlayer
            }
            
            if (playing) {
                player.play()
            } else {
                player.pause()
            }

            this.#playing = playing
        }
       
        static get watchedPercentage() { 
            return (this.playbackPosition / this.playbackDuration) * 100 
        }

        static get volume() {
            return this.#volume
        }

        static set volume(volume) {
            if (this.#quality === "dash") {
                this.#dashPlayer.setVolume(volume / 100)
            } else {
                this.#htmlPlayer.volume = volume / 100
            }
            this.#volume = volume
        }

        static get muted() {
            return this.#muted
        }

        static set muted(muted) {
            if (this.#quality === "dash") {
                this.#dashPlayer.setMute(muted)
            } else {
                this.#htmlPlayer.muted = muted
            }
            this.#muted = muted
        }

        static get fullscreen() {
            return this.#fullscreen
        }

        static set fullscreen(fullscreen) {
            if (fullscreen) {
                document.addEventListener("fullscreenchange", (event) => {
                    if (event.target === this.container && document.fullscreenElement === null) {
                        if (this.#quality === "dash") {
                            this.#dashPlayer.updatePortalSize()
                        }
                        
                        if (this.#fullscreen) {
                            WindowUnfullscreen()
                            this.#fullscreen = false
                        }
                    }
                })

                if (document.fullscreenElement === null) {
                    WindowFullscreen()

                    this.container.requestFullscreen().then(() => {
                        if (this.#quality === "dash") {
                            this.#dashPlayer.updatePortalSize()
                        }
                    })

                    this.#fullscreen = true
                }
            } else {
                if (this.#fullscreen) {
                    document.exitFullscreen()
                }
            }
        }

        static get controlsVisible() {
            return this.#controlsVisible
        }

        static set controlsVisible(controlsVisible) {
            if (controlsVisible) {
                if (this.#controlsTimeout !== null) {
                    clearTimeout(this.#controlsTimeout)
                }

                this.controls.style.opacity = "100%"
                this.controls.style.cursor = "auto"

                this.#controlsTimeout = setTimeout(() => {
                    this.controls.style.opacity = "0%"
                    this.controls.style.cursor = "none"
                }, 3000)
            } else {
                if (this.#playing) {
                    this.controls.style.opacity = "0%"
                    this.controls.style.cursor = "none"
                }
            }
            
            this.#controlsVisible = controlsVisible
        }
    }
</script>


<script>
    import { onMount } from "svelte"
    import { Play as PlayIcon, Pause as PauseIcon, SpeakerWave as SpeakerWaveIcon, SpeakerXMark as SpeakerXMarkIcon } from "@steeze-ui/heroicons";
    import { Maximize as MaximiseIcon, Minimize as MinimizeIcon, PlaySquare, Volume } from "@steeze-ui/lucide-icons";
    import { Icon } from "@steeze-ui/svelte-icon";
    import dashjs from "dashjs";
    import { WindowUnfullscreen, WindowFullscreen, LogDebug } from "../../wailsjs/runtime/runtime.js";


    export let video
    

    let playerElement
    let playerContainer
    let playerControls


    onMount(() => player.init(playerContainer, playerElement, playerControls, video))
</script>


<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div bind:this={playerContainer} class="flex relative flex-col rounded-md cursor-default aspect-video">
    <!-- svelte-ignore a11y-media-has-caption -->
    <video bind:this={playerElement} on:click={()=>player.playing=!player.playing} on:ended={()=>player.playing=false} bind:duration={player.playbackDuration} bind:currentTime={player.playbackPosition} class="my-auto w-full rounded-md">
        {#each video.captions as caption}
            <track kind="captions" label={caption.label} src={caption.url} srclang={caption.languageCode}>
        {/each}
    </video>

    <div bind:this={playerControls} class="flex absolute top-0 flex-col w-full h-full duration-[400ms]" style={(player.playing || player.fullscreen) && "opacity: 0% important;"} on:mouseleave={()=>player.controlsVisible=false} on:mousemove={()=>player.controlsVisible=true}>        
        <div class="p-1.5 mt-auto w-full h-22 child:w-full">
            <div class="flex flex-row gap-2 p-1 mt-9 h-11 child:bg-black/70 child:backdrop-blur-[100px] child:p-2 child:rounded-md hover:child:text-white/60 child:text-white/95">
                <button on:click|stopPropagation={()=>player.playing=!player.playing}>
                    <Icon src={player.playing ? PauseIcon : PlayIcon} theme="mini" size="20"></Icon>
                </button>

                <div class="!px-2.5 w-full">
                    <div style={`background: linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) ${player.watchedPercentage}%, rgb(255 255 255 / 0.6) ${player.watchedPercentage}%, rgb(255 255 255 / 0.6) 100%);`} class="absolute w-[calc(100%-20px)] top-[calc(50%-3px)] h-[6px] rounded-md"></div>
                    <input class="w-full opacity-0 cursor-pointer" type="range" min="0" max={player.playbackDuration} bind:value={player.playbackPosition} on:click|stopPropagation/>
                </div>

                <div class="flex flex-row gap-1 !pr-2.5 w-fit hover:!text-white/95">
                    <button class="mr-1 duration-100 hover:text-white/60" on:click|stopPropagation={()=>player.muted=!player.muted}>
                        <Icon src={player.muted || player.volume <= 0 ? SpeakerXMarkIcon : SpeakerWaveIcon} theme="mini" size="20"></Icon> 
                    </button>
                    <div style={`background: linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) ${player.volume}%, rgb(255 255 255 / 0.6) ${player.volume}%, rgb(255 255 255 / 0.6) 100%);`} class="absolute left-[36px] top-[calc(50%-3px)] w-16 h-[6px] rounded-md"></div>
                    <input class="w-16 opacity-0 cursor-pointer" type="range" height="10" min="0" max="100" on:change={(e)=>player.volume=parseInt(e.currentTarget.value)} value={player.volume} on:click|stopPropagation/>
                </div>

                <select on:click|stopPropagation bind:value={player.quality} class="w-fit cursor-pointer text-sm font-bold leading-5 appearance-none !border-none !outline-none !ring-0">
                    <option value="dash" selected>Quality: Dash</option>
                    {#each video.formatStreams as stream}
                        <option value={stream.qualityLabel}>Quality: {stream.qualityLabel}</option>
                    {/each}
                </select>

                <button data-b on:click|stopPropagation={()=>player.fullscreen=!player.fullscreen}>
                    <Icon src={player.fullscreen ? MinimizeIcon : MaximiseIcon} stroke-width="3" size="20"></Icon>
                </button>
            </div>
        </div>
    </div> 
</div>