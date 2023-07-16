// components
import Window from "./components/window/Window.svelte";


// styles
import "./style.css";


const invidiousDesktop = new Window({
  target: document.getElementById("app"),
});

export default invidiousDesktop;
