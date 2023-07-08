import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

const config = {
  preprocess: [vitePreprocess({})],
  compilerOptions: {
    accessors: true,
  }
};

export default config;
