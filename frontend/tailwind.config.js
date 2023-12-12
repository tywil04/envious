const plugin = require('tailwindcss/plugin')

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,svelte}",
  ],
  
  darkMode: "class",

  theme: {
    extend: {
      borderRadius: {
        "4px": "4px",
      },

      screens: {
        "video-grid-1": "360px",
        "video-grid-2": "720px",
        "video-grid-3": "1080px",
        "video-grid-4": "1440px",
        "video-grid-5": "1800px",
        "video-grid-6": "2160px",
      }
    }
  },

  plugins: [
    require('@tailwindcss/forms'),
    plugin(({ addVariant, addUtilities, matchVariant, matchUtilities }) => {
      addVariant("child", "& > *")
      matchVariant("child-data", (value) => `& > *[data-${value}]`)
      matchVariant("child-type", (value) => `& > ${value}`)

      matchVariant("parent-data", (value) => `&:has(*[data-${value}])`)

      addVariant("decendant", "& *")
      matchVariant("decendant-data", (value) => `& *[data-${value}]`)
      matchVariant("decendant-type", (value) => `& ${value}`)

      addUtilities({
        ".wails-drag": {
          "--wails-draggable": "drag"
        },
        ".wails-nodrag": {
          "--wails-draggable": "nodrag"
        },
      })

      addUtilities({
        ".word-break": {
          "word-break": "break-word"
        },
      })
    })
  ]
};