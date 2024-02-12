import xhook from "xhook";
import Window from "./Window.svelte";
import Settings from "./tabs/Settings.svelte";
import Search from "./tabs/Search.svelte";
import Trending from "./tabs/Trending.svelte";
import Config from "./tabs/Config.svelte";
import Subscriptions from "./tabs/Subscriptions.svelte";
import * as heroIcons from "@steeze-ui/heroicons"
import * as go from "../wailsjs/go/main/Envious.js"
import "./style.css";


let window

go.GetInvidiousInstance().then((instance) => {
    console.log(instance)

    let configured = instance.apiUrl !== ""

    if (configured) {
        xhook.before((request) => {
            if (request.url.includes("googlevideo")) {
                const url = new URL(request.url)
                request.url = instance.apiUrl + url.pathname + url.search.toString()
            }
        })

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
            }
        })
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
        })
    }
})

export default window;
