<script>
    // svelte
    import { onMount } from "svelte";


    // javascript
    import { GetSelectedInvidiousInstance } from "../wailsjs/go/main/InvidiousDesktop.js"
    import { newRouter, newRoute, pushRoute, currentRoute } from "./lib/router.js"

    
    // components
    import Home from "./pages/Home.svelte"
    import InstanceSetup from "./pages/InstanceSetup.svelte";
    import Titlebar from "./components/window/Titlebar.svelte";
    

    newRouter([
        newRoute("home", "Home"),
        newRoute("instanceSetup", "Invidious Instance Setup"),
    ])


    onMount(async () => {
        if (await GetSelectedInvidiousInstance() === "") {
            pushRoute("instanceSetup")
        } else {
            pushRoute("home")
        }
    })
</script> 


<Titlebar title={$currentRoute.metadata.title || "Invidious Desktop"}/>

{#if $currentRoute.path == "home"}
    <Home/>
{:else if $currentRoute.path == "instanceSetup"}
    <InstanceSetup/>
{/if}