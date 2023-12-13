import { Icon } from "@steeze-ui/svelte-icon"
import { X as XIcon } from "@steeze-ui/lucide-icons"

import adaptiveBackground from "./adaptiveBackground.js"


const tabSystem = {
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
                    theme: "mini",
                    size: "24",
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


export default tabSystem