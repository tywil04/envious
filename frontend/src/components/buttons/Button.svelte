<button {...$$props} on:click>
    <slot/>
</button>

<style lang="postcss">
    button {
        @apply border-2 border-black bg-black text-white rounded-4px px-2 py-0.5 duration-100;
    }

    button:disabled {
        @apply bg-gray-500 border-gray-500 cursor-not-allowed;
    }

    button:hover:not(:active):not(:disabled) {
        @apply bg-gray-700 border-gray-700 cursor-pointer;
    }
</style>