<script>
    import { v4 as uuidv4 } from "uuid"

    export let id = uuidv4()
    export let name = ""
    export let required = true
    export let label = ""
    export let options = []
    export let selected = ""
    export let selectedIndex = 0

    if (options[selectedIndex].display !== undefined && options[selectedIndex].value !== undefined) {
        selected = options[selectedIndex].value
    } else if (options[selectedIndex].display === undefined && options[selectedIndex].value === undefined && options[selectedIndex] !== undefined) {
        selected = options[selectedIndex]
    }

    const onInput = (event) => {
        selected = event.target.value

        for (let i = 0; i < options.length; i++) {
            if (options[i].value === selected || options[i] === selected) {
                selectedIndex = i
                break
            }
        }
    }
</script>

<div>
    <label for={id}>{label}</label>

    {#if required}
        <span class="required">(required)</span>
    {/if}

    <select on:input on:click on:input={onInput} {id} {name} {required}>
        {#each options as option, index}
            {#if option.display !== undefined && option.value !== undefined}
                <!-- if option is a map -->
                <option value={option.value} selected={index == selectedIndex}>
                    <!-- if options.display is undefined, use options.value -->
                    {option.display !== undefined ? option.display: option.value}
                </option>
            {:else if option.display === undefined && option.value === undefined && option !== undefined}
                <!-- if option is an array, use the value -->
                <option value={option} selected={index == selectedIndex}>{option}</option>
            {/if}
        {/each}
    </select>
</div>

<style lang="postcss">
    div {
        @apply prose space-y-0.5 w-full h-fit max-w-full max-h-fit
    }

    label {
        @apply font-bold
    }

    select {
        @apply border-2 border-black rounded-4px px-2 py-0.5 w-full m-0
    }

    select:focus {
        @apply outline-none ring-0
    }

    span.required {
        @apply text-sm text-gray-600
    }
</style>