import { vitePreprocess } from "@sveltejs/vite-plugin-svelte"
import sveltePreprocess from 'svelte-preprocess'

const config = {
  preprocess: [
    sveltePreprocess({ postcss: true }),
    vitePreprocess({ postcss: true }), 
  ],
  compilerOptions: {
    accessors: true,
  }
};

export default config;
