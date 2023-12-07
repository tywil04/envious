<script>
    import { v4 as uuidv4 } from "uuid"

    
    export let id = uuidv4()
    export let name = ""
    export let required = true
    export let label = ""
    export let options = []
    export let selected = null
    export let selectedIndex = 0


    let option = options[selectedIndex]
    if (option.display !== undefined && option.value !== undefined) {
        selected = option.value
    } else if (option.display === undefined && option.value === undefined && option !== undefined) {
        selected = option
    }

    
    const onInput = (event) => {
        for (let i = 0; i < options.length; i++) {
            if (i == event.target.value) {
                selectedIndex = i
                selected = options[i].value
                break
            }
        }
    }
</script>


<div class="root">
    <label class="label" for={id}>{label}</label>

    {#if required}
        <span class="notice required">(required)</span>
    {/if}

    <select class="input" on:input on:click on:input={onInput} {id} {name} {required}>
        {#each options as option, index}
            <option value={index} selected={index == selectedIndex}>
                {option.display}
            </option>
        {/each}
    </select>
</div>


<style lang="postcss">
    .root {
        @apply space-y-0.5 w-full h-fit max-w-full max-h-fit;
    }

    .required {
        @apply text-sm text-zinc-400;
    }

    .label {
        @apply font-semibold text-zinc-300;
    }

    .input {
        @apply border bg-zinc-800 border-zinc-700 text-zinc-300 rounded-md px-2 py-0.5 w-full m-0;
    }

    .input:focus {
        @apply outline-none ring-0;
    }
</style>