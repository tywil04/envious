import { WindowFullscreen, WindowUnfullscreen, WindowMaximise, WindowUnmaximise } from "../wailsjs/runtime/runtime"
import { GetConfigured } from "../wailsjs/go/main/Tubed"

import Window from "./components/window/Window.svelte";
import Home from "./tabs/Home.svelte";
import Config from "./tabs/Config.svelte";

import "./style.css";


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
GetConfigured().then(isConfigured => {
    if (isConfigured) {
        tubed = new Window({
            target: document.getElementById("app"),
            props: {
                defaultTabs: [
                    {
                        name: "Home",
                        locked: true,
                        component: Home,
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
                        locked: true,
                        component: Config,
                    },
                ]
            }
        });  
    }
})


export default tubed;
