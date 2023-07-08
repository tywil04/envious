<script>
    import { v4 as uuidv4 } from "uuid"

    import ErrorSpan from "../text/ErrorSpan.svelte";

    export let id = uuidv4()
    export let name = ""
    export let label = ""
    export let required = true
    export let value = ""
    export let errorMessage = ""
    export let validation = undefined
    export let type = "text"
    export let readOnly = false
    export let placeholder = ""

    let innerValue = ""
    let valid = false

    const onInput = (event) => {
        // check type is valid
        if (type.match(/^(text|url|email|password|search)$/)) {
            innerValue = event.target.value || ""
        }

        // check validity
        if (innerValue === "") {
            valid = false
        } else {
            switch (typeof(validation)) {
                case "function": valid = validation(innerValue)
                case "object": {
                    try {
                        valid = validation.test(innerValue)
                    } catch {
                        valid = false
                    }
                }
            }
        }

        if (valid) {
            value = innerValue
        } else {
            value = ""
        }
    }
</script>

<div>
    <label for={id}>{label}</label>

    {#if required}
        <span class="required">(required)</span>
    {/if}

    <input on:input on:click on:input={onInput} readonly={readOnly} value={innerValue} {placeholder} {name} {required} {id} {type}/>
    
    {#if innerValue !== "" && !valid}
        <ErrorSpan>{errorMessage}</ErrorSpan>
    {/if}
</div>

<style lang="postcss">
    div {
        @apply prose space-y-0.5 w-full h-fit max-w-full max-h-fit;
    }

    label {
        @apply font-bold;
    }

    input {
        @apply border-2 border-black rounded-4px px-2 py-0.5 w-full m-0;
    }

    input:focus {
        @apply outline-none ring-0;
    }

    span.required {
        @apply text-sm text-gray-600;
    }
</style>