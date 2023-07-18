<script context="module">
    let tabsElement
    let viewsElement


    function selectTab(tabName) {
        for (let child of tabsElement.children) {
            if (child.id === `tab::${tabName}`) {
                child.classList.add("active")
            } else {
                child.classList.remove("active")
            }
        }

        for (let child of viewsElement.children) {
            child.hidden = child.id !== `view::${tabName}`
        }
    }


    function deleteTab(tabName) {
        const tabButton = document.getElementById(`tab::${tabName}`)
        const wasActive = tabButton.classList.contains("active")

        tabButton.remove()
        document.getElementById(`view::${tabName}`).remove()

        if (wasActive) {
            tabsElement.children[tabsElement.children.length - 1].classList.add("active")
            viewsElement.children[viewsElement.children.length - 1].hidden = false
        }
    }


    function renderTab(tab, active) {
        const button = document.createElement("button")
        button.id = `tab::${tab.name}`
        button.classList.add("tab")
        button.title = tab.name

        const span = document.createElement("span")
        span.innerText = tab.name
        span.classList.add("text")

        button.appendChild(span)
        
        if (!tab.locked) {
            const deleteButton = document.createElement("button")
            deleteButton.classList.add("delete")

            deleteButton.addEventListener("click", (event) => {
                event.stopPropagation()
                deleteTab(tab.name)
            })

            new Icon({
                target: deleteButton,
                props: {
                    src: Close,
                    size: "24",
                }
            })

            button.appendChild(deleteButton)
        }

        if (active) {
            for (let child of tabsElement.children) {
                child.classList.remove("active")
            }
            button.classList.add("active")
        }

        button.addEventListener("click", () => selectTab(tab.name))
    
        tabsElement.appendChild(button)
    }


    function renderTabView(tab, hidden=true) {
        const tabRoot = document.createElement("div")
        tabRoot.id = `view::${tab.name}`
        
        if (!hidden) {
            for (let child of viewsElement.children) {
                child.hidden = true
            }
            tabRoot.hidden = false
        }

        new tab.component({
            target: tabRoot,
            props: tab.props
        })

        viewsElement.appendChild(tabRoot)
    }


    export function spawnTab(tab, hidden=true) {
        if (document.getElementById(`tab::${tab.name}`) === null) {
            renderTab(tab, !hidden)
            renderTabView(tab, hidden)
        } else {
            selectTab(tab.name)
        }
    }
</script>


<script>
    import { onMount } from "svelte";

    import * as runtime from "../../../wailsjs/runtime/runtime"

    import { Icon } from "@steeze-ui/svelte-icon"
    import { Close, Add, Subtract } from "@steeze-ui/carbon-icons"

    import Home from "../../pages/Home.svelte";
    

    let tabs = [
        {
            name: "Home",
            locked: true,
            component: Home,
        },
    ]
    let currentTab = 0


    onMount(() => {
        for (let i = 0; i < tabs.length; i++) {
            spawnTab(tabs[i], i !== currentTab)
        }
    })
</script> 


<nav class="navigation" on:dblclick={runtime.WindowToggleMaximise}>
    <div class="title">
        <span class="text">Invidious Desktop</span>
    </div>

    <div class="controls"> 
        <button title="Minimise Window" class="button" on:click={runtime.WindowMinimise}>
            <Icon src={Subtract} size="24"/>
        </button>
        <button title="Maximise Window" class="button" on:click={runtime.WindowToggleMaximise}>
            <Icon src={Add} size="24"/>
        </button>
        <button title="Close Window" class="button close" on:click={runtime.Quit}>
            <Icon src={Close} size="24"/>
        </button>
    </div>
</nav>


<div class="tabs" bind:this={tabsElement}></div>


<div class="views" bind:this={viewsElement}></div>


<style lang="postcss">
    ::-webkit-scrollbar {
        display: none;
    }


    /* tabs container and tab */
    .tabs {
        @apply absolute top-[40px] h-[40px] min-w-full w-full bg-black flex flex-row justify-between border-b border-zinc-800 px-[5px] py-1 opacity-100 pointer-events-auto overflow-x-auto overflow-y-hidden;
    }

    .tabs:not(:has(:nth-child(2))) {
        @apply pointer-events-none opacity-0 h-0;
    }

    :global(.tabs:not(:has(:nth-child(2))) + .views > *) {
        @apply top-[40px];
        height: calc(100% - 40px);
    }

    :global(.tab) {
        @apply bg-black h-[31px] leading-[31px] w-full text-zinc-400 truncate px-2 rounded-4px flex flex-row duration-100 relative select-none min-w-[175px];   
    }

    :global(.tab > .text) {
        @apply h-[31px] mr-auto ml-auto px-[24px] truncate;
    }

    :global(.tab > .delete) {
        @apply absolute right-2 h-[31px] opacity-0 pointer-events-none duration-100;
    }

    :global(.tab:hover > .delete) {
        @apply opacity-100 pointer-events-auto;
    }

    :global(
        .tab:hover:not(.active) > .delete:hover, 
        .tab > .delete:hover
    ) {
        @apply text-red-400;
    }

    :global(.tab:not(:last-child)) {
        @apply mr-1;
    }

    :global(.tab:hover:not(.active)) {
        @apply bg-zinc-900/60 text-zinc-300;
    }

    :global(.tab.active) {
        @apply bg-zinc-900 border-b-zinc-900/80 text-zinc-300 cursor-default;
    }


    /* views container */
    :global(.views > *) {
        @apply absolute top-[80px] p-4 max-w-[100%] w-full select-none bg-transparent overflow-y-auto z-0;
        height: calc(100% - 80px);
    }


    /* navigation */
    .navigation {
        @apply fixed top-0 w-full h-[40px] flex flex-row p-2 border-b border-zinc-800 bg-black text-zinc-200 z-10;
        --wails-draggable:drag;
    }

    .navigation .title {
        @apply select-none h-6 mr-auto;
        --wails-draggable:no-drag  
    }

    .navigation > .title > .text {
        @apply ml-1.5 font-normal;
    }

    .navigation > .controls {
        @apply select-none h-6 ml-auto flex flex-row z-50;
        --wails-draggable:no-drag
    }

    .navigation > .controls > .button {
        @apply flex flex-col justify-center duration-100 rounded-full cursor-auto text-zinc-200;
    }

    .navigation > .controls > .button:hover {
        @apply bg-zinc-800;
    }

    .navigation > .controls > .button.close:hover {
        @apply bg-red-600;
    }
</style>