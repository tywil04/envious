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
        @apply border-2 rounded-4px px-2 py-0.5 duration-100;
    }


    .normal {
        @apply border-2 border-black bg-black text-white;
    }

    .minimal {
        @apply border-transparent text-black;
    }

    .outline {
        @apply bg-white border-black text-black;
    }


    .normal:disabled {
        @apply bg-gray-600 border-gray-600 cursor-not-allowed;
    }

    .minimal:disabled, .outline:disabled {
        @apply bg-gray-400 border-gray-400 cursor-not-allowed;
    }


    .normal:hover:not(:active):not(:disabled) {
        @apply bg-gray-700 border-gray-700 cursor-pointer;
    }

    .minimal:hover:not(:active):not(:disabled), .outline:hover:not(:active):not(:disabled) {
        @apply bg-gray-200 cursor-pointer;
    }
</style>