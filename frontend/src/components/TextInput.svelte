<script>
    import { v4 as uuidv4 } from "uuid"

    
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
  

    const validate = (event) => {
        // check type is valid
        if (type.match(/^(text|url|email|password|search)$/)) {
            innerValue = event.target.value || ""
        }

        // check validity
        if (validation === undefined) {
            valid = true
        } else {
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
        }

        value = valid ? innerValue: ""
    }
</script>


<div class="root">
    <label class="label" for={id}>{label}</label>

    {#if required}
        <span class="required">(required)</span>
    {/if}

    <input class="input" on:input on:click on:input={validate} readonly={readOnly} value={innerValue} {placeholder} {name} {required} {id} {type}/>
    
    {#if innerValue !== "" && !valid}
        <span class="error">{errorMessage}</span>
    {/if}
</div>


<style lang="postcss">
    .root {
        @apply space-y-0.5 w-full h-fit max-w-full max-h-fit;
    }

    .required {
        @apply text-sm text-zinc-400;
    }

    .error {
        @apply text-sm text-red-600;
    }

    .label {
        @apply font-semibold text-zinc-300;
    }

    .input {
        @apply border bg-zinc-800 border-zinc-700 text-zinc-300 rounded-4px px-2 py-0.5 w-full m-0;
    }

    .input::placeholder {
        @apply text-zinc-500;
    }

    .input:focus {
        @apply outline-none ring-0;
    }
</style>