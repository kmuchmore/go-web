/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./view/**/*.templ",
  ],
  darkMode: "class",
  theme: {
    // colors: {},
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('daisyui')
  ],
}

