<script>
    import { onMount } from "svelte";

    import { Icon } from "@steeze-ui/svelte-icon"
    import { X as XIcon, Square as SquareIcon, Minus as MinusIcon } from "@steeze-ui/lucide-icons"

    import { WindowToggleMaximise, WindowMinimise, Quit } from "../../wailsjs/runtime/runtime.js"
    import tabSystem from "./tabSystem.js";
    import adaptiveBackground from "./adaptiveBackground.js";

    
    export let defaultTabs = []


    onMount(() => defaultTabs.forEach(tabSystem.createTab))
</script> 


<nav on:dblclick|self={WindowToggleMaximise} class="flex sticky top-0 z-20 flex-row p-2 w-full h-[44px] wails-drag child:wails-nodrag child:font-normal child:w-fit child:flex child:flex-row">
    <div class="justify-start">
        <span class="my-auto ml-2 font-semibold">Tubed</span>
    </div>

    <div class="justify-center mx-auto"></div>

    <div class="justify-end child:h-fit child:w-fit child:flex child:flex-row">
        <div class="gap-2 hover:child:bg-white/5 child:h-7 child:w-7 child:duration-75 child:rounded-md child:cursor-pointer child:p-0.5"> 
            <button title="Minimise Window" on:click={WindowMinimise}>
                <Icon src={MinusIcon} size="24" stroke-width="1.5"/>
            </button>
            <button title="Maximise Window" on:click={WindowToggleMaximise} class="!p-1.5">
                <Icon src={SquareIcon} size="16" stroke-width="2"/>
            </button>
            <button title="Close Window" on:click={Quit}>
                <Icon src={XIcon} size="24" stroke-width="1.5"/>
            </button>
        </div>
    </div>
</nav>


<div class="flex flex-row gap-2 p-2 pt-0 w-full h-[calc(100%-44px)]">
    <aside bind:this={tabSystem.tabsRootElement} class="z-10 sticky h-full overflow-auto resize-x min-w-[256px] w-[256px] flex flex-col decendant-data-[stt]:duration-100 [&_*[data-st]:hover>*[data-stt]]:text-white/60 decendant-data-[stl]:pointer-events-none decendant-data-[stl]:-z-50 decendant-data-[st]:!py-0.5 decendant-data-[st]:!bg-transparent decendant-data-[st]:h-[28px] decendant-data-[stt]:ml-[29px] decendant-data-[stl]:absolute [&_:first-child>*[data-stl]]:!h-[14px] [&_:first-child>*[data-stl]]:!-mt-0 decendant-data-[stl]:h-[30px] decendant-data-[stl]:-mt-[16px] decendant-data-[stl]:border-white/40 decendant-data-[stl]:w-3 decendant-data-[stl]:mr-1 decendant-data-[stl]:border-l decendant-data-[stl]:ml-2 decendant-data-[stl]:border-b decendant-data-[stl]:rounded-bl-sm decendant-data-[tgc]:flex decendant-data-[tgc]:flex-col gap-4 decendant-data-[tgc]:gap-1 decendant-data-[st]:text-white/75 decendant-data-[tsb]:flex decendant-data-[tsb]:flex-row decendant-data-[tsb]:truncate decendant-data-[stc]:rounded-md decendant-data-[tsb]:rounded-md hover:decendant-data-[tsb]:bg-white/5 decendant-data-[tsb]:h-8 decendant-data-[tsb]:px-2 decendant-data-[tsb]:py-1 decendant-data-[tsb]:text-left decendant-data-[tsb]:w-full decendant-data-[ta=true]:!bg-white/10 hover:decendant-data-[tsb]:duration-100 decendant-data-[ti]:py-[3px] decendant-data-[ti]:pr-[3px] [&_svg[data-ti]+span[data-tbt]]:ml-[5px] decendant-data-[tbt]:truncate decendant-data-[tdb]:absolute decendant-data-[tdb]:mr-1 decendant-data-[tdb]:rounded-md decendant-data-[tdb]:right-0 decendant-data-[tdb]:hidden [&_*[data-tsb]:hover>span[data-tbt]]:mr-[20px] [&_*[data-tsb]:hover>button[data-tdb]]:block decendant-data-[tgt]:text-sm decendant-data-[tgt]:mx-2 decendant-data-[tgt]:-mb-3 decendant-data-[tgt]:opacity-70 decendant-data-[tgt]:duration-75"></aside>

    <main bind:this={tabSystem.viewsRootElement} class="decendant-data-[v]:overflow-y-auto decendant-data-[v]:p-4 decendant-data-[v]:min-w-full w-full h-full decendant-data-[v]:w-full decendant-data-[v]:h-full decendant-data-[v]:min-h-full decendant-data-[v]:rounded-xl decendant-data-[v]:rounded-br-4px decendant-data-[v]:bg-black/40"></main>
</div>


<div bind:this={adaptiveBackground.rootElement} class="pointer-events-none absolute top-0 w-full h-full opacity-20 -z-10 duration-500 decendant-data-[abc]:absolute decendant-data-[abc]:w-full decendant-data-[abc]:h-full decendant-data-[abb]:absolute decendant-data-[abi]:absolute decendant-data-[abb]:w-full decendant-data-[abi]:w-full decendant-data-[abb]:h-full decendant-data-[abi]:h-full decendant-data-[abb]:backdrop-blur-[100px]"></div>