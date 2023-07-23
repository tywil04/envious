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
        } else {
            tabRoot.hidden = true
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

    import { WindowMinimise, WindowToggleMaximise, Quit } from "../../../wailsjs/runtime/runtime.js"

    import { Icon } from "@steeze-ui/svelte-icon"
    import { Close, Add, Subtract, Settings as Cog, Search } from "@steeze-ui/carbon-icons"

    import Settings from "../../tabs/Settings.svelte"

    
    export let defaultTabs = []


    onMount(() => {
        for (let i = 0; i < defaultTabs.length; i++) {
            renderTab(defaultTabs[i], i === 0)
            renderTabView(defaultTabs[i], i !== 0)
        }
    })


    const doubleClickMaximise = (event) => {
        if (event.target.tagName === "NAV") {
            WindowToggleMaximise()
        }
    }


    const openSettingsTab = (event) => {
        spawnTab({
            name: "Settings",
            component: Settings, 
        }, false)
    }
</script> 


<nav class="navigation" on:dblclick={doubleClickMaximise}>
    <div class="left">
        <span class="title">Tubed</span>
    </div>

    <div class="centre"></div>

    <div class="right">
        <div class="buttons">
            <button class="button settings" on:click={openSettingsTab}>
                <Icon src={Cog} size="18"/>
            </button>
        </div>

        <div class="divider"></div>

        <div class="controls"> 
            <button title="Minimise Window" class="button" on:click={WindowMinimise}>
                <Icon src={Subtract} size="24"/>
            </button>
            <button title="Maximise Window" class="button" on:click={WindowToggleMaximise}>
                <Icon src={Add} size="24"/>
            </button>
            <button title="Close Window" class="button close" on:click={Quit}>
                <Icon src={Close} size="24"/>
            </button>
        </div>
    </div>
</nav>


<div class="tabs" bind:this={tabsElement}></div>


<div class="views" bind:this={viewsElement}></div>


<style lang="postcss">
    ::-webkit-scrollbar {
        display: none;
    }


    /* for tabs */
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
        @apply text-red/90;
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

    :global(.views > *) {
        @apply absolute top-[80px] p-4 max-w-[100%] w-full select-none bg-transparent overflow-y-auto z-0;
        height: calc(100% - 80px);
    }


    .navigation {
        @apply fixed top-0 w-full h-[40px] flex flex-row p-2 border-b border-zinc-800 bg-black text-zinc-200 z-10;
        --wails-draggable:drag;
    }

    .navigation > .left {
        @apply font-normal w-fit flex flex-row justify-start select-none;
        --wails-draggable:no-drag;
    }

    .navigation > .left > .title {
        @apply ml-1 text-zinc-200;
    }

    .navigation > .centre {
        @apply font-normal w-fit flex flex-row mx-auto justify-center select-none;
        --wails-draggable:no-drag;
    }

    .navigation > .right {
        @apply font-normal w-fit flex flex-row justify-end select-none;
        --wails-draggable:no-drag;
    }

    .navigation > .right > .buttons {
        @apply flex flex-row h-6;
    }

    .navigation > .right > .divider {
        @apply mt-1 mx-2.5 h-4 w-[1px] bg-zinc-800;
    }

    .navigation > .right > .buttons > .button {
        @apply flex flex-col justify-center h-6 w-6 rounded-full;
    }

    .navigation > .right > .buttons > .button.settings {
        @apply pl-[3px];
    }

    .navigation > .right > .buttons > .button:hover {
        @apply bg-zinc-800;
    }

    .navigation > .right > .controls {
        @apply select-none h-6 w-fit flex flex-row z-50;
    }

    .navigation > .right > .controls > .button {
        @apply flex flex-col justify-center duration-100 rounded-full cursor-auto text-zinc-200;
    }

    .navigation > .right > .controls > .button:hover {
        @apply bg-zinc-800;
    }

    .navigation > .right > .controls > .button.close:hover {
        @apply bg-red-600;
    }
</style>