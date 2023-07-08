<script>
    import Router from "svelte-spa-router"

    import Home from "./pages/Home.svelte"
    import InstanceSetup from "./pages/InstanceSetup.svelte";



    // import { SetInvidiousInstance } from '../wailsjs/go/main/InvidiousDesktop.js'
    // import TextInput from './components/inputs/TextInput.svelte';
    // import InstancesInput from './components/inputs/InstancesInput.svelte';
    // import SelectInput from './components/inputs/SelectInput.svelte';

    // let instance = ""
    // let success = false

    // const submitInstance = async () => {
    //     success = await SetInvidiousInstance(instance)
    // }
</script> 

<Router routes={{
    "/": Home,
    '/instanceSetup': InstanceSetup,
}}/>