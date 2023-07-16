import "./style.css";
import Window from "./components/window/Window.svelte";

const invidiousDesktop = new Window({
  target: document.getElementById("app"),
});

export default invidiousDesktop;
