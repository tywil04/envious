<script>
    // svelte
    import { blur } from "svelte/transition"


    // javascript
    import * as runtime from "../../../wailsjs/runtime/runtime"


    // components
    import { Icon } from "@steeze-ui/svelte-icon"
    import { Close, Add, Subtract } from "@steeze-ui/carbon-icons"


    // props
    export let title = ""


    const animationDuration = 150
    const outBlur = {
        duration: animationDuration
    }
    const inBlur = {
        delay: animationDuration,
        duration: animationDuration,
    }
</script> 


<nav class="nav">
    {#key title}
        <div class="navSection left" out:blur={outBlur} in:blur={inBlur}>
            <b class="titlebarText">{title}</b>
        </div>
    {/key}

    <div class="navSection right"> 
        <button title="Minimise Window" class="windowButton" on:click={runtime.WindowMinimise}>
            <Icon src={Subtract} size="24"/>
        </button>
        <button title="Maximise Window" class="windowButton" on:click={runtime.WindowToggleMaximise}>
            <Icon src={Add} size="24"/>
        </button>
        <button title="Close Window" class="windowButton close" on:click={runtime.Quit}>
            <Icon src={Close} size="24"/>
        </button>
    </div>
</nav>


<style lang="postcss">
    .nav {
        @apply flex flex-row p-2 border-b-2 bg-black border-black text-white;
        --wails-draggable:drag;
    }

    .navSection {
        @apply select-none h-6;
        --wails-draggable:no-drag
    }

    .navSection.left {
        @apply mr-auto;
    }

    .navSection.right {
        @apply ml-auto flex flex-row z-50;
    }


    .windowButton {
        @apply flex flex-col justify-center duration-100 rounded-full cursor-auto text-white;
    }

    .windowButton:hover {
        @apply bg-blue-50/15;
    }

    .windowButton.close:hover {
        @apply bg-red-600;
    }


    .titlebarText {
        @apply ml-1.5;
    }
</style>