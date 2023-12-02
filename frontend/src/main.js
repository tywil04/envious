import { WindowFullscreen, WindowUnfullscreen, WindowMaximise, WindowUnmaximise } from "../wailsjs/runtime/runtime.js"
import { DBGet } from "../wailsjs/go/main/Tubed.js"

import { Search as SearchIcon, Cog as CogIcon, Play as PlayIcon, Flame as FlameIcon } from "@steeze-ui/lucide-icons"

import Window from "./components/Window.svelte";

import Settings from "./tabs/Settings.svelte";
import Search from "./tabs/Search.svelte";
import Trending from "./tabs/Trending.svelte";
import Config from "./tabs/Config.svelte";
import Subscriptions from "./tabs/Subscriptions.svelte";

import "./style.css";


const onFullscreen = (event) => {
    if (document.fullscreenElement !== null) {
        WindowMaximise()
        setTimeout(WindowFullscreen, 1)
    } else {
        WindowUnfullscreen()
        setTimeout(WindowUnmaximise, 1)
    }
}

document.addEventListener("fullscreenchange", onFullscreen)


let window 
Promise.all([
    DBGet("backend.configured", "bool"),
]).then(([
    configured,
]) => {
    if (configured) {
        window = new Window({
            target: document.getElementById("app"),
            props: {
                defaultTabs: [
                    {
                        name: "Trending",
                        group: "App",
                        active: true,
                        locked: true,
                        fallback: true,
                        component: Trending,
                        icon: FlameIcon
                    },
                    {
                        name: "Subscriptions",
                        group: "App",
                        locked: true,
                        component: Subscriptions,
                        icon: PlayIcon
                    },
                    {
                        name: "Search",
                        group: "App",
                        locked: true,
                        component: Search,
                        icon: SearchIcon,
                    },
                    {
                        name: "Settings",
                        group: "App",
                        locked: true,
                        component: Settings,
                        icon: CogIcon,
                    },
                ]
            }
        });
    } else {
        window = new Window({
            target: document.getElementById("app"),
            props: {
                defaultTabs: [
                    {
                        name: "Config",
                        group: "App",
                        active: true,
                        locked: true,
                        fallback: true,
                        component: Config,
                    },
                ]
            }
        });  
    }
})

export default window;
