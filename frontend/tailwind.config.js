const plugin = require('tailwindcss/plugin')

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,svelte}",
  ],
  darkMode: "class",
  theme: {},
  plugins: [
    require('@tailwindcss/forms'),
    plugin(({ addVariant, addUtilities }) => {
      addVariant("*", "& > *")

      addUtilities({
        ".wails-drag": {
          "--wails-draggable": "drag"
        },
        ".wails-nodrag": {
          "--wails-draggable": "nodrag"
        },
      })
    })
  ]
};