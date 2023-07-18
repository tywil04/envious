import { WindowFullscreen, WindowUnfullscreen } from "../wailsjs/runtime/runtime"

import Window from "./components/window/Window.svelte";

import "./style.css";


const onFullscreen = (event) => {
    if (document.fullscreenElement === event.target) {
        WindowFullscreen()
    } else {
        WindowUnfullscreen()
    }
}
document.addEventListener("fullscreenchange", onFullscreen)


const invidiousDesktop = new Window({
  target: document.getElementById("app"),
});


export default invidiousDesktop;
