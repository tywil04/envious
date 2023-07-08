<button {...$$props} on:click>
    <slot/>
</button>

<style lang="postcss">
    button {
        @apply bg-white border-2 border-black text-black rounded-4px px-2 py-0.5 duration-100;
    }

    button:disabled {
        @apply bg-gray-400 border-gray-400 cursor-not-allowed;
    }

    button:hover:not(:active):not(:disabled) {
        @apply bg-gray-200 cursor-pointer;
    }
</style>