import "./style.css";
import InvidiousDesktop from "./InvidiousDesktop.svelte";

const app = new InvidiousDesktop({
  target: document.getElementById("app"),
});

export default app;
