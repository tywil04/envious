<script>
    import * as go from "../../wailsjs/go/main/Envious.js"
</script>


{#await go.KVIsSet("background.type") then set} 
    {#if set}
        {#await go.KVGet("background.type") then backgroundType}
            <p>{backgroundType}</p>     
        {/await}
    {:else}
        <p>Color</p>
    {/if}
{/await}