<script context="module">
    import RxPlayer from "rx-player"
    import * as utils from "../lib/utils.js"
    import * as wails from "../../wailsjs/runtime/runtime.js"
    import * as Window from "../Window.svelte";

    class Player extends RxPlayer {
        video = null
        quality = "dash"

        rootElement = null
        // videoElement (already part of RxPlayer)
        controlsElement = null
        timelineSliderElement = null
        timelineTextElement = null
        volumeSliderElement = null

        volumeIcon = null
        playIcon = null
        fullscreenIcon = null

        wasMinimisedBeforeFullscreen = false

        controlsVisibleDuration = 5000 // ms
        controlsVisibleTimeout = null

        constructor(video) {
            super({
                minVideoBitrate: 200000,
                maxVideoBitrate: 30000000,
                initialVideoBitrate: 3000000,
            })

            this.video = video

            this.addEventListener("positionUpdate", (data) => {
                let watchedPercentage = (data.position / data.duration) * 100
                this.timelineSliderElement.style.setProperty("--percentage", watchedPercentage + "%")
                this.timelineSliderElement.setAttribute("aria-valuenow", watchedPercentage)
                this.timelineTextElement.innerText = utils.calculateWatchedDuration(data.position, data.duration)
            })

            this.addEventListener("playerStateChange", (state) => {
                switch (state) {
                    case "PLAYING": {
                        if (this.isFullscreen()) {
                            this.hideControlsAfterTimeout()
                        } else {
                            this.showControls()
                        }
                        break
                    }
                    case "PAUSED": {
                        this.showControls()
                        break
                    }
                    case "ENDED": {
                        this.reload({ 
                            reloadAt: { 
                                position: 0 
                            },
                            autoPlay: false,
                        })
                        this.playIcon.$$set({ src: PlayIcon })
                        break
                    }
                } 
            })

            document.addEventListener("fullscreenchange", this.#eventDocumentFullscreenChange.bind(this))
            document.addEventListener("keydown", this.#eventDocumentKeyDown.bind(this))
            wails.EventsOn("videoPlayer:pauseUnlessActive", this.#eventWailsVideoPlayerPauseUnlessActive.bind(this))

            // I dont know why but waiting 1ms makes it work
            // without this the dash video doesnt get loaded automatically
            setTimeout(() => {
                this.loadVideo(this.video)
            }, 1)
        }

        dispose() {
            document.removeEventListener("fullscreenchange", this.#eventDocumentFullscreenChange.bind(this))
            document.removeEventListener("keydown", this.#eventDocumentKeyDown.bind(this))
            wails.EventsOff("videoPlayer:pauseUnlessActive")

            super.dispose()
        }

        #eventDocumentFullscreenChange() {
            if (this.isFullscreen()) {
                wails.WindowFullscreen()
                this.fullscreenIcon.$$set({ src: MinimizeIcon })
            } else {
                wails.WindowUnfullscreen()
                this.fullscreenIcon.$$set({ src: MaximiseIcon })
                this.showControls()
            }
        }

        #eventDocumentKeyDown(event) {
            if (event.code === "Space" && Window.tabs.isElementFromActiveTab(this.rootElement) === true) {
                event.preventDefault()
                this.togglePlay()
            }
        }

        #eventWailsVideoPlayerPauseUnlessActive() {
            if (!Window.tabs.isElementFromActiveTab(this.rootElement)) {
                this.pause()
            }
        }

        loadVideo(video) {
            let currentPosition = 0
            let state = this.getPlayerState()
            if (state === "LOADED" || state === "PLAYING" || state === "PAUSED" || state === "ENDED" || state === "BUFFERING") {
                currentPosition = this.getPosition()
            }
            
            if (this.quality === "dash") {
                super.loadVideo({
                    url: video.dashUrl,
                    transport: "dash",
                    autoPlay: false,
                })
            } else {
                for (let stream of video.formatStreams) {
                    if (this.quality === stream.qualityLabel) {
                        // we know that this stream can be played because only valid streams are listed to the user
                        // thanks to Player.filterCompatibleFormatStreams
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

        async play() {
            if (this.isContentLoaded()) {
                wails.EventsEmit("videoPlayer:pauseUnlessActive")
                this.playIcon.$$set({ src: PauseIcon })
                await super.play()
            }
        }

        pause() {
            if (this.isContentLoaded()) {
                this.playIcon.$$set({ src: PlayIcon })
                super.pause()
            }
        }

        togglePlay() {
            if (this.isContentLoaded()) {
                if (this.isPlaying()) {
                    this.pause()
                } else {
                    this.play()
                }
            }
        }

        setVolume(volume) {
            if (volume > 0) {
                this.volumeIcon.$$set({ src: SpeakerWaveIcon })
            } else {
                this.volumeIcon.$$set({ src: SpeakerXMarkIcon })
            }

            let volumePercentage = Math.round(volume * 100)
            this.volumeSliderElement.style.setProperty("--percentage", volumePercentage + "%")
            this.volumeSliderElement.setAttribute("aria-valuenow", volumePercentage)

            super.setVolume(volume)
        }

        setQuality(quality) {
            this.quality = quality
            this.loadVideo(this.video)
        }

        getQuality() {
            return this.quality
        }

        toggleMute() {
            if (this.isMute()) {
                this.unMute()
            } else {
                this.mute()
            }
        }

        isFullscreen() {
            return document.fullscreenElement === this.rootElement
        }

        enterFullscreen() {
            if (document.fullscreenElement === null) {
                this.rootElement.requestFullscreen()
            }
        }

        exitFullscreen() {
            if (this.isFullscreen()) {
                document.exitFullscreen()
            }
        }

        toggleFullscreen() {
            if (this.isFullscreen()) {
                this.exitFullscreen()
            } else {
                this.enterFullscreen()
            }
        }

        setPosition(position) {
            this.seekTo({ position })
        }

        filterCompatibleFormatStreams() {
            return this.video.formatStreams.filter((stream) => {
                return this.videoElement.canPlayType(stream.type) === "probably"
            })
        }

        isPlaying() {
            return this.getPlayerState() === "PLAYING"
        }

        showControls() {
            if (this.controlsVisibleDuration !== null) {
                clearTimeout(this.controlsVisibleTimeout)
            }

            this.controlsElement.style.opacity = "1"
        }

        showControlsUntilTimeout() {
            if (this.controlsVisibleDuration !== null) {
                clearTimeout(this.controlsVisibleTimeout)
            }

            this.showControls()
            
            if (this.isFullscreen()) {
                this.controlsVisibleTimeout = setTimeout(() => {
                    this.hideControls()
                    this.controlsVisibleTimeout = null
                }, this.controlsVisibleDuration)
            }
        }

        hideControls() {
            this.controlsElement.style.opacity = "0"
        }

        hideControlsAfterTimeout() {
            if (this.isFullscreen()) {
                if (this.controlsVisibleDuration !== null) {
                    clearTimeout(this.controlsVisibleTimeout)
                }

                this.controlsVisibleTimeout = setTimeout(() => {
                    this.hideControls()
                    this.controlsVisibleTimeout = null
                }, this.controlsVisibleDuration)
            }
        }
    }
</script>


<script>
    import { Play as PlayIcon, Pause as PauseIcon, SpeakerWave as SpeakerWaveIcon, SpeakerXMark as SpeakerXMarkIcon } from "@steeze-ui/heroicons";
    import { Maximize as MaximiseIcon, Minimize as MinimizeIcon } from "@steeze-ui/lucide-icons";
    import { Icon } from "@steeze-ui/svelte-icon";
    import { onDestroy } from "svelte"

    export let video

    const player = new Player(video)

    onDestroy(() => player.dispose())

    console.log(video)
</script>


<div class="player" bind:this={player.rootElement} role="none" on:mousemove={()=>player.showControlsUntilTimeout()}>
    <video bind:this={player.videoElement} class="video">
        {#each player.video.captions as caption}
            <track kind="captions" src={caption.url} lang={caption.languageCode} type={caption.type} label={caption.label}/>
        {/each}
    </video>

    <div class="controlsOverlay">  
        <button data-focus-invisible class="clickToPlay" on:click={()=>player.togglePlay()} ></button>
        
        <div bind:this={player.controlsElement} class="controls">
            <button class="segment interactive" on:click={()=>player.togglePlay()}>
                <Icon bind:this={player.playIcon} src={PlayIcon} theme="mini" size="20"></Icon>
            </button>

            <div class="segment timeline">
                <button role="slider" aria-valuenow={0} bind:this={player.timelineSliderElement} class="slider interactive" on:click={(e)=>player.setPosition((e.offsetX / e.currentTarget.clientWidth) * player.getVideoDuration())}></button> 
            
                <p bind:this={player.timelineTextElement} class="watchedDuration">00:00 / 00:00</p>
            </div>

            <div class="segment volume">
                <button class="interactive mute" on:click={()=>player.toggleMute()}>
                    <Icon bind:this={player.volumeIcon} src={SpeakerWaveIcon} theme="mini" size="20"></Icon> 
                </button>

                <button role="slider" aria-valuenow={100} style="--percentage: 100%;" on:click={(e)=>player.setVolume(e.offsetX / e.currentTarget.clientWidth)} bind:this={player.volumeSliderElement} class="slider interactive"></button>
            </div>

            <select on:change={(e)=>player.setQuality(e.currentTarget.value)} value={player.getQuality()} class="segment quality interactive">
                <option value="dash" selected>Quality: Dash</option>

                {#each player.filterCompatibleFormatStreams() as stream}
                    <option value={stream.qualityLabel}>Quality: {stream.qualityLabel}</option>
                {/each}
            </select>

            <button class="segment interactive" data-b on:click={()=>player.toggleFullscreen()}>
                <Icon bind:this={player.fullscreenIcon} src={MaximiseIcon} stroke-width="3" size="20"></Icon>
            </button>
        </div>
    </div> 
</div>


<style>
    .player {
        display: flex;
        position: relative;
        flex-direction: column;
        cursor: default;
        aspect-ratio: 16 / 9;
        background: transparent;
        border-radius: inherit;

        & .video {
            margin-top: auto;
            margin-bottom: auto;
            width: 100%;
            border-radius: inherit;
        }

        & .controlsOverlay {
            display: flex;
            position: absolute;
            top: 0;
            flex-direction: column;
            width: 100%;
            height: 100%;
            
            & .clickToPlay {
                height: 100%;
                cursor: default;
            }

            & .controls {
                width: 100%;
                display: flex;
                flex-direction: row;
                gap: 4px;
                padding: 4px;
                height: 40px;
                transition-duration: 200ms;

                & .segment {
                    background: rgba(0, 0, 0, 0.6);
                    padding: 6px;
                    border-radius: 6px;
                    box-shadow: inset 0 0 0.313rem 0.125rem rgba(0, 0, 0, 0.15);
                    backdrop-filter: blur(10px) contrast(65%) brightness(90%);
                    transition-duration: 100ms;

                    &.interactive, & .interactive {
                        color: rgba(255, 255, 255, 0.95);
                    }

                    &.interactive:hover, & .interactive:hover {
                        color: rgba(255, 255, 255, 0.7);
                    }

                    &.quality {
                        width: fit-content;
                        cursor: pointer;
                        font-size: 0.875rem;
                        line-height: 1.25rem;
                        font-weight: 400;
                        appearance: none;
                        border: none !important;
                        outline: none !important;
                        padding-left: 8px;
                        padding-right: 8px;
                    }

                    &.timeline {
                        width: 100%;
                        display: flex;
                        flex-direction: row;

                        & .watchedDuration {
                            font-size: 0.875rem;
                            line-height: 1.25rem;
                            font-weight: 400;
                            width: fit-content;
                            white-space: nowrap;
                            margin-left: 8px;
                        }
                    }

                    &.volume {
                        display: flex;
                        flex-direction: row;
                        width: fit-content;

                        & .mute {
                            margin-right: 8px;
                        }

                        & .slider {
                            width: 64px !important;
                        }
                    }

                    & > .slider {
                        --percentage: 0%;
                        background: linear-gradient(to right, rgb(255 255 255 / 0.95) 0%, rgb(255 255 255 / 0.95) var(--percentage), rgb(255 255 255 / 0.6) var(--percentage), rgb(255 255 255 / 0.6) 100%);
                        margin-top: 7px;
                        cursor: pointer;
                        height: 6px;
                        border-radius: 6px;
                        width: 100%;
                        margin-left: 2px;
                        margin-right: 2px;
                        transition-duration: 100ms;

                        &.interactive:hover {
                            background: linear-gradient(to right, rgb(255 255 255 / 0.9) 0%, rgb(255 255 255 / 0.9) var(--percentage), rgb(255 255 255 / 0.55) var(--percentage), rgb(255 255 255 / 0.55) 100%);
                        }
                    }
                }
            }
        }
    }
</style>