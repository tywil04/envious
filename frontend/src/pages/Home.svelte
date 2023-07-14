<script>
    // javascript
    import { GetPopular, GetTrending, GetSelectedInstance } from "../../wailsjs/go/main/InvidiousDesktop.js"
    import { link } from "svelte-spa-router"


    // components
    import Main from "../components/window/Main.svelte";
    import VideoRow from "../components/video/VideoRow.svelte";
    import Button from "../components/buttons/Button.svelte";
</script>


<Main>
    <a href="/instanceSetup" use:link>Set or Change Invidious instance used.</a>
    
    {#await GetSelectedInstance()}
        <p>Loading...</p>
    {:then instance} 
        <div class="container">
            {#if instance !== ""}
                <p>Your selected instance is: {instance}.</p>
            {:else}
                <p>You dont have a selected instance.</p>
            {/if}

            <div class="divider"/>

            <div class="buttonContainer">
                <Button>Test Button</Button>

                <div class="seperator"/>

                <Button outline>Test Button</Button>

                <div class="seperator"/>

                <Button minimal>Test Button</Button>
            </div>

            <div class="divider"/>

            <VideoRow dataFunction={GetPopular}/>

            <div class="divider"/>
            
            <VideoRow dataFunction={GetTrending}/>
        </div>
    {:catch error}
        <span class="errorText">{error}</span>
    {/await}
</Main>


<style lang="postcss">
    .container {
        @apply flex flex-col max-w-full;
    }

    .buttonContainer {
        @apply flex flex-row;
    }


    .divider {
        @apply mb-4;
    }

    .seperator {
        @apply mr-2;
    }


    .errorText {
        @apply text-red-600;
    }
</style>