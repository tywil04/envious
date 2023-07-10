<script>
    // svelte
    // import { onMount } from "svelte";


    // javascript
    import { location } from "svelte-spa-router"

    
    // components
    import Home from "./pages/Home.svelte"
    import InstanceSetup from "./pages/InstanceSetup.svelte";
    import Titlebar from "./components/window/Titlebar.svelte";
    import Router from "svelte-spa-router"


    // onMount(async () => {
    //     if (await GetSelectedInvidiousInstance() === "") {
    //         pushRoute("instanceSetup")
    //     } else {
    //         pushRoute("home")
    //     }
    // })


    // use hooks to determin what page is seen


    const routes = {
        "/": Home,
        "/instanceSetup": InstanceSetup
    }

    const routesTitleMap = {
        "/": "Home",
        "/instanceSetup": "Invidious Instance Configuration"
    }
</script> 


<Titlebar title={routesTitleMap[$location]}/>
<Router routes={routes}/>