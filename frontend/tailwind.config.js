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
        "4px": "0.25rem"
      },
      opacity: {
        '15': '0.15',
        '35': '0.35',
        '65': '0.65',
      },
      colors: {
        "glassBlack": "rgba(17, 25, 40, 0.7)",
        "glassWhite": "rgba(255, 255, 255, 0.7)"
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ]
};