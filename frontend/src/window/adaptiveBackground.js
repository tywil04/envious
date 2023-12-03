const adaptiveBackground = {
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
        container.dataset.abc = "true"

        const containerDiv = document.createElement("div")
        containerDiv.style.zIndex = "-11"
        containerDiv.dataset.abb = "true"
        container.append(containerDiv)
    
        const containerImg = document.createElement("img")
        containerImg.src = url
        containerImg.alt = "blurry background image"
        containerImg.crossOrigin = "anonymous"
        containerImg.style.zIndex = "-12"
        containerImg.dataset.abi = "true"
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


export default adaptiveBackground