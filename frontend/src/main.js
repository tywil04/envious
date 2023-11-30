import { WindowFullscreen, WindowUnfullscreen, WindowMaximise, WindowUnmaximise } from "../wailsjs/runtime/runtime.js"
import { GetConfigured } from "../wailsjs/go/main/Tubed.js"

import Window from "./components/Window.svelte";
import Home from "./tabs/Home.svelte";
import Config from "./tabs/Config.svelte";

import "./style.css";
import Settings from "./tabs/Settings.svelte";
import Search from "./tabs/Search.svelte";


// make fullscreen
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


let tubed 
GetConfigured().then((isConfigured) => {
    if (isConfigured) {
        tubed = new Window({
            target: document.getElementById("app"),
            props: {
                defaultTabs: [
                    {
                        name: "Home",
                        group: "App",
                        active: true,
                        locked: true,
                        fallback: true,
                        component: Home,
                        
                    },
                    {
                        name: "Search",
                        group: "App",
                        locked: true,
                        component: Search,
                    },
                    {
                        name: "Settings",
                        group: "App",
                        locked: true,
                        component: Settings,
                    },
                ]
            }
        });
    } else {
        tubed = new Window({
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

export default tubed;
