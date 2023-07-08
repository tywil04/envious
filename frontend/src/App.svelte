<script>
    import * as runtime from "../wailsjs/runtime"

    import Router from "svelte-spa-router"

    import Home from "./pages/Home.svelte"
    import InstanceSetup from "./pages/InstanceSetup.svelte";

    import { Icon } from "@steeze-ui/svelte-icon"
    import { Close, Add, Subtract } from "@steeze-ui/carbon-icons"
</script> 

<nav>
    <div class="left">
        <b>Invidious Desktop</b>
    </div>

    <div class="right"> 
        <button title="Minimise Window" class="windowControl" on:click={runtime.WindowMinimise}>
            <Icon src={Subtract} size="32"/>
        </button>
        <button title="Maximise Window" class="windowControl" on:click={runtime.WindowMaximise}>
            <Icon src={Add} size="32"/>
        </button>
        <button title="Close Window" class="windowControl close" on:click={runtime.Quit}>
            <Icon src={Close} size="32"/>
        </button>
    </div>
</nav>

<main>
    <Router routes={{
        "/": Home,
        '/instanceSetup': InstanceSetup,
    }}/>
</main>

<style lang="postcss">
    main {
        @apply p-2;
    }

    nav {
        @apply flex flex-row bg-black p-2;
    }

    nav > div.left {
        @apply prose h-8;
    }

    nav > div.left > b {
        @apply text-lg leading-8 text-white;
    }

    nav > div.right {
        @apply ml-auto flex flex-row -space-x-0.5;
    }

    button.windowControl {
        @apply flex flex-col justify-center duration-100 rounded-full cursor-auto text-white;
    }

    button.windowControl:hover {
        @apply bg-blue-50/20 text-white;
    }

    button.windowControl.close:hover {
        @apply bg-red-600 text-white;
    }
</style>