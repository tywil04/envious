import { writable } from "svelte/store"


export const newRoute = (path, title) => {
    return {
        path,
        metadata: {
            title,
        }
    }
}


export const newRouter = (allowedRoutes, defaultRoute=newRoute("blank", "")) => {
    currentRoute.set(defaultRoute)

    for (let route of allowedRoutes) {
        storedAllowedRoutes[route.path] = route
    }
}


export const pushRoute = (route) => {
    const current = storedAllowedRoutes[route]

    if (current !== undefined) {
        currentRoute.set(current)
    }
}


export const pushRouteAction = (node, route) => {
    const setRoute = () => pushRoute(route)

    node.addEventListener("click", setRoute)

    return {
        destroy() {
            node.removeEventListener("click", setRoute)
        }
    }
}


export let currentRoute = writable(newRoute("blank", ""))
export let storedAllowedRoutes = {}