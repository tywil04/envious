import Window from "./window/Window.svelte";

import Settings from "./tabs/Settings.svelte";
import Search from "./tabs/Search.svelte";
import Trending from "./tabs/Trending.svelte";
import Config from "./tabs/Config.svelte";
import Subscriptions from "./tabs/Subscriptions.svelte";

import { Search as SearchIcon, Cog as CogIcon, Play as PlayIcon, Flame as FlameIcon } from "@steeze-ui/lucide-icons"

import { DBGet } from "../wailsjs/go/main/Tubed.js"

import "./style.css";


const configuredDefaultTabs = [
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

const unconfiguredDefaultTabs = [
    {
        name: "Config",
        group: "App",
        active: true,
        locked: true,
        fallback: true,
        component: Config,
    },
]


let window

function loadWindow([configured]) {
    let defaultTabs = unconfiguredDefaultTabs
    if (configured) {
        // @ts-ignore
        defaultTabs = configuredDefaultTabs
    }

    window = new Window({
        target: document.getElementById("app"),
        props: { defaultTabs }
    });
}

Promise.all([
     DBGet("backend.configured", "bool")
]).then(loadWindow)


export default window;
