import xhook from "xhook";

import Window from "./window/Window.svelte";

import Settings from "./tabs/Settings.svelte";
import Search from "./tabs/Search.svelte";
import Trending from "./tabs/Trending.svelte";
import Config from "./tabs/Config.svelte";
import Subscriptions from "./tabs/Subscriptions.svelte";

import { MagnifyingGlass as MagnifyingGlassIcon, Cog8Tooth as Cog8ToothIcon, Play as PlayIcon, Fire as FireIcon } from "@steeze-ui/heroicons"

import { GetBackendConfigured, GetSelectedInstance } from "../wailsjs/go/main/Tubed.js"

import "./style.css";


const configuredDefaultTabs = [
    {
        name: "Trending",
        group: "App",
        active: true,
        locked: true,
        fallback: true,
        component: Trending,
        icon: FireIcon
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
        icon: MagnifyingGlassIcon,
    },
    {
        name: "Settings",
        group: "App",
        locked: true,
        component: Settings,
        icon: Cog8ToothIcon,
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


// googlevideo doesn't allow cors, invidious can proxy this for us so replace all requests to googlevideo
GetSelectedInstance().then((instance) => {
    xhook.before((request) => {
        if (request.url.includes("googlevideo")) {
            const url = new URL(request.url)
            request.url = instance.apiUrl + url.pathname + url.search.toString()
        }
    })
})


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
    GetBackendConfigured()
]).then(loadWindow)


export default window;
