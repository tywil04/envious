<script>
    // svelte
    import { blur } from "svelte/transition"


    const animationDuration = 150
    const outBlur = {
        duration: animationDuration
    }
    const inBlur = {
        delay: animationDuration,
        duration: animationDuration,
    }
</script> 


<main class="container" out:blur={outBlur} in:blur={inBlur}>
    <slot/>
</main>


<style lang="postcss">
    .container {
        @apply p-4 max-w-[100%] w-full max-h-full select-none bg-transparent;
    }


    /* @media (prefers-color-scheme: dark) {
        button.windowControl:hover {
            @apply bg-blue-50/20 text-white;
        }
    } */
</style>