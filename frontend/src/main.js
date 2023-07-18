import Window from "./components/window/Window.svelte";

import "./style.css";


const invidiousDesktop = new Window({
  target: document.getElementById("app"),
});


export default invidiousDesktop;
