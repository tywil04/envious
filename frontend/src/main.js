import { WindowFullscreen, WindowUnfullscreen, WindowMaximise, WindowUnmaximise } from "../wailsjs/runtime/runtime"

import Window from "./components/window/Window.svelte";
import Home from "./tabs/Home.svelte";

import "./style.css";


const onFullscreen = (event) => {
    if (document.fullscreenElement !== null) {
        WindowMaximise()
        // setTimeout(WindowFullscreen, 1)
        WindowFullscreen()
    } else {
        WindowUnfullscreen()
        // setTimeout(WindowUnmaximise, 1)
        WindowUnmaximise()
    }
}
document.addEventListener("fullscreenchange", onFullscreen)


const invidiousDesktop = new Window({
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


export default invidiousDesktop;
