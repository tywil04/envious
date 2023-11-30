<script context="module">
    import ColorThief from "colorthief"
    import * as StackBlur from 'stackblur-canvas';

    export const tabSystem = {
        tabsRootElement: null,
        viewsRootElement: null,

        fallbackTab: null,
        tabGroups: {},

        createTabGroup: (groupName) => {
            const div = document.createElement("div")
            
            tabSystem.tabGroups[groupName] = {
                element: div,
                tabs: {},
            }

            const span = document.createElement("span")
            span.innerText = groupName
            span.style.display = "none"
                
            tabSystem.tabGroups[groupName].labelElement = span

            tabSystem.tabsRootElement.appendChild(span)
            tabSystem.tabsRootElement.appendChild(div)
        },

        createTab: (groupName, tabName, component, props, locked=false, active=false, fallback=false) => {
            if (!(groupName in tabSystem.tabGroups)) {
                tabSystem.createTabGroup(groupName)
            }
            
            if (tabName in tabSystem.tabGroups[groupName].tabs) {
                tabSystem.selectTab(groupName, tabName)
                return
            }

            const tabElement = document.createElement("button")
            tabElement.title = tabName

            const tabText = document.createElement("span")
            tabText.innerText = tabName

            tabElement.appendChild(tabText)
            tabElement.addEventListener("click", () => tabSystem.selectTab(groupName, tabName))
            tabElement.dataset.tabActive = "false"

            if (!locked) {
                const deleteButton = document.createElement("button")
                deleteButton.addEventListener("click", (event) => {
                    event.preventDefault()
                    event.stopPropagation()
                    tabSystem.deleteTab(groupName, tabName)
                })

                new Icon({
                    target: deleteButton,
                    props: {
                        src: X,
                        size: "24",
                        "stroke-width": "1.5",
                    }
                })

                tabElement.appendChild(deleteButton)
            }

            tabSystem.tabGroups[groupName].element.appendChild(tabElement)

            tabElement.scrollIntoView({ 
                inline: "end",
                behavior: "smooth", 
            })

            const viewElement = document.createElement("div")
            viewElement.hidden = true

            new component({
                target: viewElement,
                props: props
            })

            tabSystem.viewsRootElement.appendChild(viewElement)

            tabSystem.tabGroups[groupName].tabs[tabName] = {
                tabElement: tabElement,
                viewElement: viewElement,
            }

            if (tabSystem.tabGroups[groupName].labelElement) {
                tabSystem.tabGroups[groupName].labelElement.style.display = "inline"
            }

            if (fallback) {
                tabSystem.fallbackTab = [groupName, tabName]
            }

            if (active) {
                tabSystem.selectTab(groupName, tabName)
            }
        },

        selectTab: (groupName, tabName) => {
            if (!(groupName in tabSystem.tabGroups)) {
                tabSystem.createTabGroup(groupName)
            }

            for (const group of Object.values(tabSystem.tabGroups)) {
                for (const [tabName2, tab] of Object.entries(group.tabs)) {
                    if (tabName2 == tabName) {
                        tab.tabElement.dataset.tabActive = "true"
                        tab.viewElement.hidden = false
                    } else {
                        tab.tabElement.dataset.tabActive = "false"
                        tab.viewElement.hidden = true
                    }
                }
            }
        },

        deleteTab: (groupName, tabName) => {
            if (groupName in tabSystem.tabGroups && tabName in tabSystem.tabGroups[groupName].tabs) {
                const tab = tabSystem.tabGroups[groupName].tabs[tabName]
                // const wasActive = tab.tabElement.dataset.tabActive === "true"

                tab.tabElement.remove()
                tab.viewElement.remove()

                tabSystem.selectTab(...tabSystem.fallbackTab)

                delete tabSystem.tabGroups[groupName].tabs[tabName]

                if (Object.keys(tabSystem.tabGroups[groupName].tabs).length === 0) {
                    tabSystem.tabGroups[groupName].labelElement.style.display = "none"
                }
            }
        }
    }

    export const adaptiveBackground = {
        rootElement: null,

        lastCanvasElement: null,
        canvasesInTransition: {},
        visibleCanvasId: null,

        // @ts-ignore
        colorThief: new ColorThief(),

        transitionDuration: 400,

        setBackgroundFromImage: async (id, imgElement) => {
            if (id in adaptiveBackground.canvasesInTransition || adaptiveBackground.visibleCanvasId === id) {
                return
            }

            if (adaptiveBackground.lastCanvasElement === null) {
                const canvas = document.createElement("canvas")
                canvas.style.zIndex = "-9"
                canvas.style.transitionDuration = `${adaptiveBackground.transitionDuration}ms`
                canvas.style.opacity = "0"
                adaptiveBackground.canvasesInTransition[id] = true
                
                adaptiveBackground.rootElement.appendChild(canvas)

                StackBlur.image(imgElement, canvas, 150, false);

                canvas.style.opacity = "1"
                canvas.style.zIndex = "-10"

                adaptiveBackground.lastCanvasElement = canvas
                adaptiveBackground.visibleCanvasId = id
                delete adaptiveBackground.canvasesInTransition[id]
            } else {
                const canvas = document.createElement("canvas")
                canvas.style.zIndex = "-9"
                canvas.style.transitionDuration = `${adaptiveBackground.transitionDuration}ms`
                canvas.style.opacity = "0"
                adaptiveBackground.canvasesInTransition[id] = true

                adaptiveBackground.rootElement.appendChild(canvas)

                StackBlur.image(imgElement, canvas, 150, false);

                canvas.style.opacity = "1"
                canvas.style.zIndex = "-10"

                setTimeout(() => {
                    adaptiveBackground.lastCanvasElement.remove()
                    adaptiveBackground.lastCanvasElement = canvas
                    adaptiveBackground.visibleCanvasId = id
                    delete adaptiveBackground.canvasesInTransition[id]
                }, adaptiveBackground.transitionDuration + 1)
            }
        }
    }
</script>


<script>
    import { onMount } from "svelte";

    import { WindowMinimise, WindowToggleMaximise, Quit } from "../../wailsjs/runtime/runtime.js"

    import { Icon } from "@steeze-ui/svelte-icon"
    import { X, Square, Minus } from "@steeze-ui/lucide-icons"

    
    export let defaultTabs = []


    onMount(() => {
        for (const tab of defaultTabs) {
            tabSystem.createTab(
                tab.group,
                tab.name,
                tab.component,
                tab.props,
                tab.locked,
                tab.active,
                tab.fallback
            )
        }
    })


    const doubleClickMaximise = (event) => {
        if (event.target.tagName === "NAV") {
            WindowToggleMaximise()
        }
    }
</script> 


<nav class="wails-drag sticky top-0 w-full h-fit flex flex-row p-2 text-white z-20 *:wails-nodrag *:font-normal *:w-fit *:flex *:flex-row" on:dblclick={doubleClickMaximise}>
    <div class="justify-start">
        <span class="my-auto ml-2 font-semibold text-white">Tubed</span>
    </div>

    <div class="justify-center mx-auto"></div>

    <div class="justify-end *:h-fit *:w-fit *:flex *:flex-row">
        <div class="gap-2 hover:*:bg-black/40 *:h-7 *:w-7 *:duration-75 *:rounded-md *:cursor-pointer *:text-white"> 
            <button title="Minimise Window" class="p-0.5" on:click={WindowMinimise}>
                <Icon src={Minus} size="24" stroke-width="1.5"/>
            </button>
            <button title="Maximise Window" on:click={WindowToggleMaximise} class="p-1.5">
                <Icon src={Square} size="16" stroke-width="2"/>
            </button>
            <button title="Close Window" on:click={Quit} class="p-0.5">
                <Icon src={X} size="24" stroke-width="1.5"/>
            </button>
        </div>
    </div>
</nav>


<main class="flex flex-row gap-2 p-2 pt-0 w-full h-[calc(100%-44px)]">
    <div class="sticky z-10 h-full min-w-[256px] w-64 flex flex-col *:flex *:flex-col gap-4 *:gap-1 *:*:flex *:*:flex-row *:*:truncate *:*:rounded-md hover:*:*:bg-white/5 *:*:h-8 *:*:px-2 *:*:py-1 *:*:text-left *:*:w-full data-[tab-active=true]:*:*:!bg-white/10 hover:*:*:duration-75 [&>span]:*:*:max-w-[240px] [&>span]:hover:*:*:max-w-[220px] [&>span]:*:*:truncate [&>button]:*:*:absolute [&>button]:*:*:mr-1 [&>button]:*:*:rounded-md [&>button]:*:*:right-0 [&>button]:*:*:hidden [&>button]:hover:*:*:block [&>span]:text-sm [&>span]:mx-2 [&>span]:-mb-3 [&>span]:opacity-70 [&>span]:duration-75" bind:this={tabSystem.tabsRootElement}></div>
    
    <div class="p-2 rounded-md bg-black/40 max-w-[calc(100%-256px-8px)] w-full h-full" bind:this={tabSystem.viewsRootElement}></div>
</main>


<div class="opacity-[15%] absolute top-0 !w-full !h-full -z-10 *:absolute *:top-0 *:!w-full *:!h-full *:duration-500 *:-z-10 *:pointer-events-none" bind:this={adaptiveBackground.rootElement}></div>