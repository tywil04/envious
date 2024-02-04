<script context="module">
    import RxPlayer from "rx-player"

    class Player extends RxPlayer {
        #quality = "dash"
        #video = null

        controlsElement = null
        positionElement = null
        volumeElement = null 

        muteIcon = null
        playIcon = null

        constructor() {
            super({
                minVideoBitrate: 200000,
                initialVideoBitrate: 2000000,
            })

            this.addEventListener("positionUpdate", (data) => {
                let watchedPercentage = (data.position / data.duration) * 100
                this.positionElement.style.background = `linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) ${watchedPercentage}%, rgb(255 255 255 / 0.6) ${watchedPercentage}%, rgb(255 255 255 / 0.6) 100%)`
            })

            this.addEventListener("volumeChange", (volume) => {
                if (this.isMute()) {
                    this.muteIcon.$$set({ src: SpeakerXMarkIcon })
                    this.volumeElement.style.background = `linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.6) 0%, rgb(255 255 255 / 0.6) 100%)`
                } else {
                    this.muteIcon.$$set({ src: SpeakerWaveIcon })
                    let volumePercentage = volume * 100
                    this.volumeElement.style.background = `linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) ${volumePercentage}%, rgb(255 255 255 / 0.6) ${volumePercentage}%, rgb(255 255 255 / 0.6) 100%)`
                }
            })

            this.addEventListener("play", () => {
                this.playIcon.$$set({ src: PauseIcon })
            })

            this.addEventListener("pause", () => {
                this.playIcon.$$set({ src: PlayIcon })
            })

            this.addEventListener("playerStateChange", (state) => {
                if (state === "ENDED") {
                    this.reload({ 
                        reloadAt: { 
                            position: 0 
                        },
                        autoPlay: false,
                    })
                    this.playIcon.$$set({ src: PlayIcon })
                }
            })
        }

        loadVideo(video) {
            let currentPosition = 0
            let state = this.getPlayerState()
            if (state === "LOADED" || state === "PLAYING" || state === "PAUSED" || state === "ENDED" || state === "BUFFERING") {
                currentPosition = this.getPosition()
            }
            
            if (this.#quality === "dash") {
                super.loadVideo({
                    url: video.dashUrl,
                    transport: "dash",
                    autoPlay: false,
                })
            } else {
                for (let stream of video.formatStreams) {
                    if (this.#quality === stream.qualityLabel) {
                        super.loadVideo({
                            url: stream.url,
                            transport: "directfile",
                            autoPlay: false,
                        })
                        break
                    }
                }
            }

            const seekWhenLoaded = (state) => {
                if (state === "LOADED" && currentPosition !== 0) {
                    this.seekTo({ position: currentPosition })
                    this.removeEventListener("playerStateChange", seekWhenLoaded)
                }
            }
            this.addEventListener("playerStateChange", seekWhenLoaded)
        }

        setVideo(video) {
            this.#video = video
            this.loadVideo(this.#video)
        }

        getVideo() {
            return this.#video
        }

        setQuality(quality) {
            this.#quality = quality
            this.loadVideo(this.#video)
        }

        getQuality() {
            return this.#quality
        }

        togglePlay() {
            if (this.isPaused()) {
                this.play()
            } else {
                this.pause()
            }
        }

        toggleMute() {
            if (this.isMute()) {
                this.unMute()
            } else {
                this.mute()
            }
        }

        isFullscreen() {
            return document.fullscreenElement === this.videoElement
        }

        setPosition(position) {
            this.seekTo({ position })
        }
    }
</script>


<script>
    import { Play as PlayIcon, Pause as PauseIcon, SpeakerWave as SpeakerWaveIcon, SpeakerXMark as SpeakerXMarkIcon } from "@steeze-ui/heroicons";
    import { Maximize as MaximiseIcon, Minimize as MinimizeIcon } from "@steeze-ui/lucide-icons";
    import { Icon } from "@steeze-ui/svelte-icon";
    import { onMount, onDestroy } from "svelte";


    export let video
    

    let player = new Player()


    onMount(() => {
        player.setVideo(video)
    })

    onDestroy(() => {
        player.dispose()
    })
</script>


<div class="flex relative flex-col rounded-md cursor-default aspect-video">
    <video bind:this={player.videoElement} on:click={()=>player.togglePlay()} class="my-auto w-full rounded-md">
        <track kind="captions"/>
        {#each video.captions as caption}
            <track kind="captions" label={caption.label} src={caption.url} srclang={caption.languageCode}/>
        {/each}
    </video>

    <div bind:this={player.controlsElement} class="flex absolute top-0 flex-col w-full h-full duration-[400ms]" style={(!player.isPaused() || player.isFullscreen()) && "opacity: 0% important;"}>        
        <div class="p-1.5 mt-auto w-full h-22 child:w-full">
            <div class="flex flex-row gap-2 p-1 mt-9 h-11 child:bg-black/70 child:backdrop-blur-[100px] child:p-2 child:rounded-md hover:child:text-white/60 child:text-white/95">
                <button on:click|stopPropagation={()=>player.togglePlay()}>
                    <Icon bind:this={player.playIcon} src={PlayIcon} theme="mini" size="20"></Icon>
                </button>

                <div class="!px-2.5 w-full">
                    <div 
                        role="slider"
                        aria-valuemax={player.getVideoDuration()}
                        aria-valuemin={0}
                        aria-valuenow={0}
                        tabindex="-1"
                        on:click={(e)=>player.setPosition((e.offsetX / e.currentTarget.clientWidth) * player.getVideoDuration())}
                        on:keypress={(e)=>console.log(e)} 
                        bind:this={player.positionElement}  
                        style={`background: linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.6) 0%, rgb(255 255 255 / 0.6) 100%);`} 
                        class="z-50 mt-[7px] w-full cursor-pointer h-[6px] rounded-md"
                    ></div>                    
                </div>

                <div class="flex flex-row gap-1 !pr-2.5 w-fit hover:!text-white/95">
                    <button class="mr-1 duration-100 hover:text-white/60" on:click|stopPropagation={()=>player.toggleMute()}>
                        <Icon bind:this={player.muteIcon} src={SpeakerWaveIcon} theme="mini" size="20"></Icon> 
                    </button>
                    <div on:click={(e)=>player.setVolume(e.offsetX / e.currentTarget.clientWidth)} bind:this={player.volumeElement} style={`background: linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) ${player.getVolume() * 100}%, rgb(255 255 255 / 0.6) ${player.getVolume() * 100}%, rgb(255 255 255 / 0.6) 100%);`} class="w-16 mt-[7px] cursor-pointer h-[6px] rounded-md"></div>
                </div>

                <select on:click|stopPropagation on:change={(e)=>player.setQuality(e.currentTarget.value)} value={player.getQuality()} class="w-fit cursor-pointer text-sm font-bold leading-5 appearance-none !border-none !outline-none !ring-0">
                    <option value="dash" selected>Quality: Dash</option>
                    {#each video.formatStreams as stream}
                        <option value={stream.qualityLabel}>Quality: {stream.qualityLabel}</option>
                    {/each}
                </select>

                <button data-b on:click|stopPropagation>
                    <Icon src={player.isFullscreen() ? MinimizeIcon : MaximiseIcon} stroke-width="3" size="20"></Icon>
                </button>
            </div>
        </div>
    </div> 
</div>