<script>
    // svelte
    import { onMount } from "svelte";


    // javascript
    import * as runtime from "../../../wailsjs/runtime/runtime"


    // components
    import { Icon } from "@steeze-ui/svelte-icon"
    import { Close, Add, Subtract } from "@steeze-ui/carbon-icons"
    import Home from "../../pages/Home.svelte";
    

    let tabs = [
        {
            name: "Home",
            locked: true,
            component: Home,
        }
    ]
    let currentTab = 0

    onMount(() => {
        for (let i = 0; i < tabs.length; i++) {
            spawnTab(tabs[i], i != currentTab)
        }
    })
</script> 


<script context="module">
    let buttonsContainer
    let viewsContainer

    function selectTab(tabName) {
        for (let child of buttonsContainer.children) {
            if (child.id === `tabbutton::${tabName}`) {
                child.classList.add("active")
            } else {
                child.classList.remove("active")
            }
        }

        for (let child of viewsContainer.children) {
            child.hidden = child.id !== `tab::${tabName}`
        }
    }

    function deleteTab(tabName) {
        const tabButton = document.getElementById(`tabbutton::${tabName}`)
        const wasActive = tabButton.classList.contains("active")

        tabButton.remove()
        document.getElementById(`tab::${tabName}`).remove()

        if (wasActive) {
            buttonsContainer.children[buttonsContainer.children.length - 1].classList.add("active")
            viewsContainer.children[viewsContainer.children.length - 1].hidden = false
        }
    }

    function renderTabButton(tab, active) {
        const button = document.createElement("button")
        button.id = `tabbutton::${tab.name}`
        button.classList.add("tabButton")
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
            for (let child of buttonsContainer.children) {
                child.classList.remove("active")
            }
            button.classList.add("active")
        }

        button.addEventListener("click", () => selectTab(tab.name))
    
        buttonsContainer.appendChild(button)
    }

    function renderTab(tab, hidden=true) {
        const tabRoot = document.createElement("div")
        tabRoot.id = `tab::${tab.name}`
        
        if (!hidden) {
            for (let child of viewsContainer.children) {
                child.hidden = true
            }
            tabRoot.hidden = false
        }

        new tab.component({
            target: tabRoot,
            props: tab.props
        })

        viewsContainer.appendChild(tabRoot)
    }

    export function spawnTab(tab, hidden=true) {
        if (document.getElementById(`tab::${tab.name}`) === null) {
            renderTabButton(tab, !hidden)
            renderTab(tab, hidden)
        } else {
            selectTab(tab.name)
        }
    }
</script>


<nav class="nav" on:dblclick={runtime.WindowToggleMaximise}>
    <div class="navSection left">
        <span class="titlebarText">Invidious Desktop</span>
    </div>

    <div class="navSection right"> 
        <button title="Minimise Window" class="windowButton" on:click={runtime.WindowMinimise}>
            <Icon src={Subtract} size="24"/>
        </button>
        <button title="Maximise Window" class="windowButton" on:click={runtime.WindowToggleMaximise}>
            <Icon src={Add} size="24"/>
        </button>
        <button title="Close Window" class="windowButton close" on:click={runtime.Quit}>
            <Icon src={Close} size="24"/>
        </button>
    </div>
</nav>

<div class="tabButtonsContainer" bind:this={buttonsContainer}></div>

<div class="viewsContainer" bind:this={viewsContainer}></div>


<style lang="postcss">
    .tabButtonsContainer {
        @apply absolute top-[40px] h-[40px] min-w-full w-full bg-black flex flex-row justify-between border-b border-zinc-800 p-1 opacity-100 pointer-events-auto overflow-x-auto overflow-y-hidden;
    }

    :global(.tabButton) {
        @apply bg-black h-[31px] leading-[31px] w-full text-zinc-400 truncate px-2 rounded-4px flex flex-row duration-100 relative select-none min-w-[175px];   
    }

    ::-webkit-scrollbar {
        display: none;
    }

    :global(.tabButton > .text) {
        @apply h-[31px] mr-auto ml-auto px-[24px] truncate;
    }

    :global(.tabButton > .delete) {
        @apply absolute right-2 h-[31px] opacity-0 pointer-events-none duration-100;
    }

    :global(.tabButton:hover > .delete) {
        @apply opacity-100 pointer-events-auto;
    }

    :global(.tabButton:hover:not(.active) > .delete:hover, .tabButton > .delete:hover) {
        @apply text-red-400;
    }

    :global(.tabButton:not(:last-child)) {
        @apply mr-1;
    }

    :global(.tabButton:hover:not(.active)) {
        @apply bg-zinc-900/60 text-zinc-300;
    }

    :global(.tabButton.active) {
        @apply bg-zinc-900 border-b-zinc-900/80 text-zinc-300 cursor-default;
    }

    .tabButtonsContainer:not(:has(:nth-child(2))) {
        @apply pointer-events-none opacity-0 h-0;
    }

    :global(.viewsContainer > *) {
        @apply absolute top-[80px] p-4 max-w-[100%] w-full select-none bg-transparent overflow-y-auto z-0;
        height: calc(100% - 80px);
    }

    :global(.tabButtonsContainer:not(:has(:nth-child(2))) + .viewsContainer > *) {
        @apply top-[40px];
        height: calc(100% - 40px);
    }

    .nav {
        @apply fixed top-0 w-full h-[40px] flex flex-row p-2 border-b border-zinc-800 bg-black text-zinc-200 z-10;
        --wails-draggable:drag;
    }

    .navSection {
        @apply select-none h-6;
        --wails-draggable:no-drag
    }

    .navSection.left {
        @apply mr-auto;
    }

    .navSection.right {
        @apply ml-auto flex flex-row z-50;
    }

    .windowButton {
        @apply flex flex-col justify-center duration-100 rounded-full cursor-auto text-zinc-200;
    }

    .windowButton:hover {
        @apply bg-zinc-800;
    }

    .windowButton.close:hover {
        @apply bg-red-600;
    }

    .titlebarText {
        @apply ml-1.5 font-normal;
    }
</style>