import xhook from "xhook";
import Window from "./Window.svelte";
import Settings from "./tabs/Settings.svelte";
import Search from "./tabs/Search.svelte";
import Trending from "./tabs/Trending.svelte";
import Config from "./tabs/Config.svelte";
import Subscriptions from "./tabs/Subscriptions.svelte";
import * as heroIcons from "@steeze-ui/heroicons"
import * as go from "../wailsjs/go/main/Tubed.js"
import "./style.css";


const configuredDefaultTabs = [
    {
        name: "Trending",
        group: "App",
        active: true,
        locked: true,
        fallback: true,
        component: Trending,
        icon: heroIcons.Fire,
    },
    {
        name: "Subscriptions",
        group: "App",
        locked: true,
        component: Subscriptions,
        icon: heroIcons.Play
    },
    {
        name: "Search",
        group: "App",
        locked: true,
        component: Search,
        icon: heroIcons.MagnifyingGlass,
    },
    {
        name: "Settings",
        group: "App",
        locked: true,
        component: Settings,
        icon: heroIcons.Cog8Tooth,
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

go.WaitForBackend().then(() => {
    // googlevideo doesn't allow cors, invidious can proxy this for us so replace all requests to googlevideo
    go.GetSelectedInstance().then((instance) => {
        xhook.before((request) => {
            if (request.url.includes("googlevideo")) {
                const url = new URL(request.url)
                request.url = instance.apiUrl + url.pathname + url.search.toString()
            }
        })
    })

    const backendData = Promise.all([
        go.GetBackendConfigured()
    ])

    backendData.then(([configured]) => {
        let defaultTabs = unconfiguredDefaultTabs
        if (configured) {
            // @ts-ignore
            defaultTabs = configuredDefaultTabs
        }
    
        window = new Window({
            target: document.getElementById("app"),
            props: { defaultTabs }
        });
    })
})


export default window;
