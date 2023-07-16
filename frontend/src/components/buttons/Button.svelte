<script>
    // exports
    export let minimal = false
    export let outline = false
    

    let normal
    $: normal = !minimal && !outline 
</script>


<button {...$$props} on:click class="button" class:normal={normal} class:minimal={minimal} class:outline={outline}>
    <slot/>
</button>


<style lang="postcss">
    .button {
        @apply border outline-none rounded-4px px-2 py-0.5 duration-100;
    }

    .normal {
        @apply bg-blue-700 border-blue-600 text-zinc-100;
    }

    .minimal {
        @apply border-0 bg-zinc-900 text-zinc-300;
    }

    .outline {
        @apply bg-zinc-900 border-zinc-700;
    }

    .normal:disabled,
    .minimal:disabled, 
    .outline:disabled {
        @apply brightness-75 cursor-not-allowed;
    }

    .normal:hover:not(:active):not(:disabled),
    .minimal:hover:not(:active):not(:disabled), 
    .outline:hover:not(:active):not(:disabled) {
        @apply brightness-125 cursor-pointer;
    }
</style>