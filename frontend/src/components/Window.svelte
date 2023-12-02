<script context="module">
    export const tabSystem = {
        tabsRootElement: null,
        viewsRootElement: null,

        fallbackTab: null,
        tabGroups: {},

        createTabGroup: (groupName) => {
            const div = document.createElement("div")
            div.dataset.tgc = "true"
            
            tabSystem.tabGroups[groupName] = {
                element: div,
                tabs: {},
            }

            const span = document.createElement("span")
            span.innerText = groupName
            span.style.display = "none"
            span.dataset.tgt = "true"
                
            tabSystem.tabGroups[groupName].labelElement = span

            tabSystem.tabsRootElement.appendChild(span)
            tabSystem.tabsRootElement.appendChild(div)
        },

        // groupName, tabName, component, props, locked=false, active=false, fallback=false, icon=null, backgroundUrl=null
        createTab: (tab) => {
            if (!(tab.group in tabSystem.tabGroups)) {
                tabSystem.createTabGroup(tab.group)
            }
            
            if (tab.name in tabSystem.tabGroups[tab.group].tabs) {
                tabSystem.selectTab(tab.group, tab.name)
                return
            }

            const tabElement = document.createElement("button")
            tabElement.title = tab.name
            tabElement.dataset.tsb = "true"

            const tabText = document.createElement("span")
            tabText.innerText = tab.name
            tabText.dataset.tbt = "true"

            if (tab.icon != null) {
                new Icon({
                    target: tabElement,
                    props: {
                        src: tab.icon,
                        size: "24",
                        "stroke-width": "1.5",
                        "data-ti": "true"
                    }
                })
            }

            tabElement.appendChild(tabText)
            tabElement.addEventListener("click", () => {
                tabSystem.selectTab(tab.group, tab.name)
            })
            tabElement.dataset.ta = "false"

            if (!tab.locked) {
                const deleteButton = document.createElement("button")
                deleteButton.dataset.tdb = "true"
                deleteButton.addEventListener("click", (event) => {
                    event.preventDefault()
                    event.stopPropagation()
                    tabSystem.deleteTab(tab.group, tab.name)
                })

                new Icon({
                    target: deleteButton,
                    props: {
                        src: XIcon,
                        size: "24",
                        "stroke-width": "1.5",
                    }
                })

                tabElement.appendChild(deleteButton)
            }

            tabSystem.tabGroups[tab.group].element.appendChild(tabElement)

            tabElement.scrollIntoView({ 
                inline: "end",
                behavior: "smooth", 
            })

            const viewElement = document.createElement("div")
            viewElement.hidden = true
            viewElement.dataset.v = "true"

            new tab.component({
                target: viewElement,
                props: {
                    tab: {
                        groupName: tab.group,
                        tabName: tab.name,
                        backgroundUrl: tab.backgroundUrl
                    },
                    ...tab.props
                }
            })

            tabSystem.viewsRootElement.appendChild(viewElement)

            tabSystem.tabGroups[tab.group].tabs[tab.name] = {
                tabElement: tabElement,
                viewElement: viewElement,
                backgroundUrl: tab.backgroundUrl,
            }

            if (tabSystem.tabGroups[tab.group].labelElement) {
                tabSystem.tabGroups[tab.group].labelElement.style.display = "inline"
            }

            if (tab.fallback) {
                tabSystem.fallbackTab = [tab.group, tab.name]
            }

            if (tab.active) {
                tabSystem.selectTab(tab.group, tab.name)
            }
        },

        selectTab: (groupName, tabName) => {
            if (!(groupName in tabSystem.tabGroups)) {
                tabSystem.createTabGroup(groupName)
            }

            for (const group of Object.values(tabSystem.tabGroups)) {
                for (const [tabName2, tab] of Object.entries(group.tabs)) {
                    if (tabName2 == tabName) {
                        tab.tabElement.dataset.ta = "true"
                        tab.viewElement.hidden = false

                        if (tab.backgroundUrl !== undefined) {
                            adaptiveBackground.setBackground(tab.backgroundUrl)
                        } else {
                            adaptiveBackground.resetBackground()
                        }
                    } else {
                        tab.tabElement.dataset.ta = "false"
                        tab.viewElement.hidden = true
                    }
                }
            }
        },

        deleteTab: (groupName, tabName) => {
            if (groupName in tabSystem.tabGroups && tabName in tabSystem.tabGroups[groupName].tabs) {
                const tab = tabSystem.tabGroups[groupName].tabs[tabName]
                const wasActive = tab.tabElement.dataset.ta === "true"

                tab.tabElement.remove()
                tab.viewElement.remove()

                if (wasActive) {
                    tabSystem.selectTab(...tabSystem.fallbackTab)
                }

                delete tabSystem.tabGroups[groupName].tabs[tabName]

                if (Object.keys(tabSystem.tabGroups[groupName].tabs).length === 0) {
                    tabSystem.tabGroups[groupName].labelElement.style.display = "none"
                }
            }
        },
    }

    export const adaptiveBackground = {
        rootElement: null,

        last: null,

        transitionDuration: 400,

        resetBackground: () => {
            if (adaptiveBackground.last !== null) {
                adaptiveBackground.last.style.opacity = "0"
                setTimeout(() => {
                    adaptiveBackground.last.remove()
                    adaptiveBackground.last = null
                    
                    for (const child of adaptiveBackground.rootElement.children) {
                        child.remove()
                    }
                }, adaptiveBackground.transitionDuration)
            }
        },

        setBackground: (url) => {
            const container = document.createElement("div")
            container.style.transitionDuration = `${adaptiveBackground.transitionDuration}ms`
            container.style.opacity = "0"
            container.style.zIndex = "-9"

            const containerDiv = document.createElement("div")
            containerDiv.style.zIndex = "-11"
            container.append(containerDiv)

            const containerImg = document.createElement("img")
            containerImg.src = url
            containerImg.alt = "blurry background image"
            containerImg.crossOrigin = "anonymous"
            containerImg.style.zIndex = "-12"
            container.append(containerImg)

            adaptiveBackground.rootElement.append(container)

            setTimeout(() => {
                container.style.opacity = "1"
                container.style.zIndex = "-10"
            }, 1)

            if (adaptiveBackground.last === null) {
                adaptiveBackground.last = container
            } else {
                setTimeout(() => {
                    adaptiveBackground.last.remove()
                    adaptiveBackground.last = container
                }, adaptiveBackground.transitionDuration)
            }
        }
    }
</script>


<script>
    import { onMount } from "svelte";

    import { WindowMinimise, WindowToggleMaximise, Quit } from "../../wailsjs/runtime/runtime.js"

    import { Icon } from "@steeze-ui/svelte-icon"
    import { X as XIcon, Square as SquareIcon, Minus as MinusIcon, X } from "@steeze-ui/lucide-icons"

    
    export let defaultTabs = []


    onMount(() => {
        for (const tab of defaultTabs) {
            tabSystem.createTab(tab)
        }

    })


    const doubleClickMaximise = (event) => {
        if (event.target.tagName === "NAV") {
            WindowToggleMaximise()
        }
    }
</script> 


<nav class="wails-drag sticky top-0 w-full h-fit flex flex-row p-2 z-20 *:wails-nodrag *:font-normal *:w-fit *:flex *:flex-row" on:dblclick={doubleClickMaximise}>
    <div class="justify-start">
        <span class="my-auto ml-2 font-semibold">Tubed</span>
    </div>

    <div class="justify-center mx-auto"></div>

    <div class="justify-end *:h-fit *:w-fit *:flex *:flex-row">
        <div class="gap-2 hover:*:bg-white/5 *:h-7 *:w-7 *:duration-75 *:rounded-md *:cursor-pointer"> 
            <button title="Minimise Window" class="p-0.5" on:click={WindowMinimise}>
                <Icon src={MinusIcon} size="24" stroke-width="1.5"/>
            </button>
            <button title="Maximise Window" on:click={WindowToggleMaximise} class="p-1.5">
                <Icon src={SquareIcon} size="16" stroke-width="2"/>
            </button>
            <button title="Close Window" on:click={Quit} class="p-0.5">
                <Icon src={XIcon} size="24" stroke-width="1.5"/>
            </button>
        </div>
    </div>
</nav>


<main class="flex flex-row gap-2 p-2 pt-0 w-full h-[calc(100%-44px)]">
    <div class="sticky z-10 h-full overflow-auto resize-x min-w-[256px] w-[256px] flex flex-col decendant-data-[tgc]:flex decendant-data-[tgc]:flex-col gap-4 decendant-data-[tgc]:gap-1 decendant-data-[tsb]:flex decendant-data-[tsb]:flex-row decendant-data-[tsb]:truncate decendant-data-[tsb]:rounded-md hover:decendant-data-[tsb]:bg-white/5 decendant-data-[tsb]:h-8 decendant-data-[tsb]:px-2 decendant-data-[tsb]:py-1 decendant-data-[tsb]:text-left decendant-data-[tsb]:w-full decendant-data-[ta=true]:!bg-white/10 hover:decendant-data-[tsb]:duration-100 decendant-data-[ti]:py-[3px] decendant-data-[ti]:pr-[3px] [&_svg[data-ti]+span[data-tbt]]:ml-[5px] decendant-data-[tbt]:truncate decendant-data-[tdb]:absolute decendant-data-[tdb]:mr-1 decendant-data-[tdb]:rounded-md decendant-data-[tdb]:right-0 decendant-data-[tdb]:hidden [&_*[data-tsb]:hover>span[data-tbt]]:mr-[20px] [&_*[data-tsb]:hover>button[data-tdb]]:block decendant-data-[tgt]:text-sm decendant-data-[tgt]:mx-2 decendant-data-[tgt]:-mb-3 decendant-data-[tgt]:opacity-70 decendant-data-[tgt]:duration-75" bind:this={tabSystem.tabsRootElement}></div>

    <div class="decendant-data-[v]:overflow-y-auto decendant-data-[v]:p-4 decendant-data-[v]:min-w-full w-full h-full decendant-data-[v]:w-full decendant-data-[v]:h-full decendant-data-[v]:min-h-full decendant-data-[v]:rounded-xl decendant-data-[v]:rounded-br-4px decendant-data-[v]:bg-black/40" bind:this={tabSystem.viewsRootElement}></div>
</main>


<div class="pointer-events-none absolute top-0 w-full h-full opacity-20 -z-10 duration-500 *:absolute *:w-full *:h-full *:*:absolute *:*:w-full *:*:h-full *:*:backdrop-blur-[100px]" bind:this={adaptiveBackground.rootElement}></div>