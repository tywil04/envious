<script context="module">
    class Background {
        enabled = true

        backgroundElement = null
    
        transitionDuration = 400
        
        #last = null
    
        reset() {
            if (this.enabled) {
                if (this.#last !== null) {
                    this.#last.style.opacity = "0"
                    setTimeout(() => {
                        this.#last.remove()
                        this.#last = null
                    }, this.transitionDuration)
                }
            }
        }
    
        setBackgroundUrl(url) {
            console.log(this.enabled)
            if (this.enabled) {
                const container = document.createElement("div")
                container.style.transitionDuration = `${this.transitionDuration}ms`
                container.style.opacity = "0"
                container.style.zIndex = "-9"
                container.classList.add("part")
        
                const filter = document.createElement("div")
                filter.style.zIndex = "-11"
                filter.classList.add("filter")
                container.append(filter)
            
                const img = document.createElement("img")
                img.src = url
                img.alt = "blurry background image"
                img.crossOrigin = "anonymous"
                img.style.zIndex = "-12"
                img.classList.add("image")
                container.append(img)
        
                this.backgroundElement.append(container)
        
                setTimeout(() => {
                    container.style.opacity = "1"
                    container.style.zIndex = "-10"
                }, 1)
        
                if (this.#last === null) {
                    this.#last = container
                } else {
                    setTimeout(() => {
                        this.#last.remove()
                        this.#last = container
                    }, this.transitionDuration)
                }
            }
        }
    }

    class Tabs {
        tabbarElement = null
        viewsElement = null

        #fallback = null
        #active = null
        #groups = {}

        createGroup(groupName) {
            if (groupName in this.#groups) {
                return
            }

            const container = document.createElement("div")
            container.classList.add("tabgroup")

            const label = document.createElement("span")
            label.innerText = groupName
            label.style.display = "none"
            label.classList.add("tabgroup:label")

            this.#groups[groupName] = {
                label: label,
                container: container,
                tabs: {},
            }

            this.tabbarElement.appendChild(label)
            this.tabbarElement.appendChild(container)
        }

        create(...newTabs) {
            for (let tab of newTabs) {
                if (!(tab.group in this.#groups)) {
                    tabs.createGroup(tab.group)
                }

                if (tab.name in this.#groups[tab.group].tabs) {
                    this.select(tab.group, tab.name)
                    break
                }

                const button = document.createElement("button")
                button.title = tab.name
                button.classList.add("tab")

                const buttonText = document.createElement("span")
                buttonText.innerText = tab.name
                buttonText.classList.add("label")

                if (tab.icon != null) {
                    new Icon({
                        target: button,
                        props: {
                            src: tab.icon,
                            theme: "mini",
                            size: "24",
                            class: "icon",
                        }
                    })
                }

                button.appendChild(buttonText)
                button.addEventListener("click", () => {
                    this.select(tab.group, tab.name)
                })
                button.classList.remove("active")

                if (!tab.locked) {
                    const deleteButton = document.createElement("button")
                    deleteButton.classList.add("delete")
                    deleteButton.addEventListener("click", (event) => {
                        event.preventDefault()
                        event.stopPropagation()
                        this.delete(tab.group, tab.name)
                    })

                    new Icon({
                        target: deleteButton,
                        props: {
                            src: lucideIcons.X,
                            size: "24",
                            "stroke-width": "1.5",
                        }
                    })

                    button.appendChild(deleteButton)
                }

                this.#groups[tab.group].container.appendChild(button)

                button.scrollIntoView({ 
                    inline: "end",
                    behavior: "smooth", 
                })

                const view = document.createElement("div")
                view.hidden = true
                view.classList.add("view")

                let svelteComponent = new tab.component({
                    target: view,
                    props: tab.props,
                })

                this.viewsElement.appendChild(view)

                const tabCopy = Object.assign({}, tab)
                delete tabCopy.name 

                this.#groups[tab.group].tabs[tab.name] = {
                    button,
                    view,
                    svelteComponent,
                    ...tabCopy
                }

                if (this.#groups[tab.group].label) {
                    this.#groups[tab.group].label.style.display = "inline"
                }

                if (tab.fallback && this.#fallback === null) {
                    this.#fallback = {
                        groupName: tab.group,
                        tabName: tab.name,
                    }
                }

                if (tab.active) {
                    this.select(tab.group, tab.name)
                }
            }
        }

        select(groupName, tabName) {
            if (!(groupName in this.#groups)) {
                return
            }

            if (!(tabName in this.#groups[groupName].tabs)) {
                return
            }

            if (this.#active !== null) {
                this.#groups[this.#active.groupName].tabs[this.#active.tabName].button.classList.remove("active")
                this.#groups[this.#active.groupName].tabs[this.#active.tabName].view.hidden = true
                this.#groups[this.#active.groupName].tabs[this.#active.tabName].active = false
            }
            
            this.#groups[groupName].tabs[tabName].button.classList.add("active")
            this.#groups[groupName].tabs[tabName].view.hidden = false
            this.#groups[groupName].tabs[tabName].active = true

            if (this.#groups[groupName].tabs[tabName].backgroundUrl !== undefined) {
                background.setBackgroundUrl(this.#groups[groupName].tabs[tabName].backgroundUrl)
            } else {
                background.reset()
            }

            this.#active = {
                groupName,
                tabName,
            }
        }

        delete(groupName, tabName) {
            if (!(groupName in this.#groups)) {
                return
            }

            if (!(tabName in this.#groups[groupName].tabs)) {
                return
            }

            if (this.#groups[groupName].tabs[tabName].button.classList.contains("active")) {
                if (this.#fallback !== null) {
                    tabs.select(this.#fallback.groupName, this.#fallback.tabName)
                }
            }

            this.#groups[groupName].tabs[tabName].svelteComponent.$destroy()
            this.#groups[groupName].tabs[tabName].button.remove()
            this.#groups[groupName].tabs[tabName].view.remove()

            delete this.#groups[groupName].tabs[tabName]

            if (Object.keys(this.#groups[groupName].tabs).length === 0) {
                this.#groups[groupName].label.style.display = "none"
            }
        }

        isElementFromActiveTab(element) {
            let parent = element
            while (parent != null && parent.tagName !== "HTML") {
                if (parent.classList.contains("view") && parent.parentElement.classList.contains("views")) {
                    return !parent.hidden
                } else {
                    parent = parent.parentElement
                }
            }
            return false
        }
    }

    export const background = new Background()
    export const tabs = new Tabs()
</script>


<script>
    import { onMount } from "svelte";
    import * as wails from "../wailsjs/runtime/runtime.js"
    import { Icon } from "@steeze-ui/svelte-icon"
    import * as lucideIcons from "@steeze-ui/lucide-icons"

    export let defaultTabs = []

    onMount(() => tabs.create(...defaultTabs))
</script> 


<nav on:dblclick|self={wails.WindowToggleMaximise} class="titlebar">
    <div class="start">
        <span class="title">Envious</span>
    </div>

    <div class="end">
        <div class="window-controls"> 
            <button class="minimise" title="Minimise Window" on:click={wails.WindowMinimise}>
                <Icon src={lucideIcons.Minus} size="24" stroke-width="1.5"/>
            </button>
            <button class="maximise" title="Maximise Window" on:click={wails.WindowToggleMaximise}>
                <Icon src={lucideIcons.Square} size="16" stroke-width="2"/>
            </button>
            <button class="close" title="Close Window" on:click={wails.Quit}>
                <Icon src={lucideIcons.X} size="24" stroke-width="1.5"/>
            </button>
        </div>
    </div>
</nav>

<div class="content">
    <aside bind:this={tabs.tabbarElement} class="tabbar"></aside>
    <main bind:this={tabs.viewsElement} class="views"></main>
</div>

<div bind:this={background.backgroundElement} class="background">
    <div class="noise"></div>
</div>


<style>
    .titlebar {
        display: flex;
        position: sticky;
        top: 0;
        z-index: 20;
        flex-direction: row;
        padding: 0.5rem;
        width: 100%;
        height: 2.75rem;
        --wails-draggable: drag;

        & > .start,  
        & > .end {
            --wails-draggable: nodrag;
            font-weight: 400;
            width: fit-content;
            display: flex;
            flex-direction: row;
        }

        & > .start {
            justify-content: flex-start;

            & > .title {
                margin-top: auto;
                margin-bottom: auto;
                padding: 0.15rem 0.5rem 0.15rem 0.5rem;
                background: rgba(255, 255, 255, 0.05);
                border-radius: 4px;
                font-weight: 600;
            }
        }

        & > .end {
            justify-content: end;
            margin-left: auto;

            & > .window-controls {
                height: fit-content;
                width: fit-content;
                display: flex;
                flex-direction: row;
                gap: 0.5rem;

                & > .minimise:hover, 
                & > .maximise:hover, 
                & > .close:hover {
                    background-color: rgba(255, 255, 255, 0.05)
                }

                & > .minimise, 
                & > .maximise, 
                & > .close {
                    height: 1.75rem;
                    width: 1.75rem;
                    transition-duration: 100ms;
                    border-radius: 0.375rem;
                    cursor: pointer;
                    padding: 0.125rem;
                }

                & > .maximise {
                    padding: 0.375rem !important;
                }
            }
        }
    } 
    
    .content {
        display: flex;
        flex-direction: row;
        gap: 0.5rem;
        padding: 0.5rem;
        padding-top: 0;
        width: 100%;
        height: calc(100% - 2.75rem);
    }

    .tabbar {
        z-index: 10;
        position: sticky;
        height: 100%;
        overflow: auto;
        resize: horizontal;
        min-width: 16rem;
        width: 16rem;
        display: flex;
        flex-direction: column;
        gap: 1rem;

        & > .tabgroup {
            display: flex;
            flex-direction: column;
            gap: 0.25rem;

            & > .tab {
                display: flex;
                flex-direction: row;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                border-radius: 0.375rem;
                padding: 0.25rem 0.5rem 0.25rem 0.5rem;
                text-align: left;
                height: 2rem;
                widows: 100%;

                &:hover {
                    background-color: rgba(255, 255, 255, 0.05);
                    transition-duration: 100ms;

                    & .label {
                        margin-right: 1.25rem;
                    }

                    & > .delete {
                        display: block;
                    }
                }

                &.active {
                    background-color: rgba(255, 255, 255, 0.1) !important;
                }

                & > .icon {
                    padding: 0.188rem 0.188rem 0.188rem 0;

                    & + .label {
                        margin-left: 0.313rem;
                    }
                }

                & > .label {
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                }

                & > .delete {
                    position: absolute;
                    margin-right: 0.25rem;
                    border-radius: 0.375rem;
                    right: 0;
                    display: none;
                }
            }
        }

        & > .tabgroup\:label {
            font-size: 0.875rem;
            line-height: 1.25rem;
            margin: 0 0.5rem -0.75rem 0.5rem;
            opacity: 0.7;
        }
    }

    .views {
        height: 100%;
        width: 100%;
        clip-path: inset(0 0 0 0 round 0.75rem 0.75rem 0.25rem 0.75rem);

        & > .view {
            border-radius: 0.75rem;
            border-bottom-right-radius: 0.25rem;
            box-shadow: inset 0 0 0.313rem 0.125rem rgba(0, 0, 0, 0.15);
            height: 100%;
            min-height: 100%;
            min-width: 100%;
            overflow-y: auto;
            padding: 1rem;
            width: 100%;
            background: rgba(0, 0, 0, 0.4);
        }
    }

    .background {
        pointer-events: none;
        position: absolute;
        top: 0;
        width: 100%;
        height: 100%;
        opacity: 0.2;
        z-index: -10;
        will-change: contents;
        transform: translate3d(0, 0, 0);

        & > .noise {
            background-repeat: repeat;
            background-image: url(/assets/noise.png);
            opacity: 0.01;
            position: absolute;
            width: 100%;
            height: 100%;
        }

        & > .part {
            position: absolute;
            width: 100%;
            height: 100%;
            will-change: opacity;
            transform: translate3d(0, 0, 0);

            & > .filter, 
            & > .image {
                position: absolute;
                width: 100%;
                height: 100%;
            }

            & > .filter {
                backdrop-filter: blur(1000px) brightness(0.99);
            }
        }
    }
</style>