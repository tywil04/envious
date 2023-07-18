<script>
    import { GetPopular, GetTrending, GetSelectedInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"

    import VideoRow from "../components/video/VideoRow.svelte";
</script>

    
{#await GetSelectedInstance()}
    <p>Loading...</p>
{:then instance} 
    <div class="root">
        {#if instance !== ""}
            <p>Your selected instance is: {instance}.</p>
        {:else}
            <p>You dont have a selected instance.</p>
        {/if}

        <div class="divider"/>

        <p>Popular:</p>
        <VideoRow dataFunction={GetPopular}/>

        <div class="divider"/>
        
        <p>Trending:</p>
        <VideoRow dataFunction={GetTrending}/>
    </div>
{:catch error}
    <span class="errorText">{error}</span>
{/await}


<style lang="postcss">
    .root {
        @apply flex flex-col max-w-full;
    }

    .divider {
        @apply mb-4;
    }

    .errorText {
        @apply text-red-600;
    }
</style>