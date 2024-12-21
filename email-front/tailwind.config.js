/** @type {import('tailwindcss').Config} */

import scrollbarPlugin from 'tailwind-scrollbar';
import flowbitePlugin from 'flowbite/plugin';
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    "./node_modules/flowbite/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    scrollbarPlugin,
    flowbitePlugin,
  ],
}

