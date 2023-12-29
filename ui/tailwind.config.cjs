/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      backgroundImage: {
        hero: "url('./src/assets/leaves.png')",
      },
      width: {
        128: "32rem",
      },
      colors: {
        gridcolor: "rgb(156 163 175)",
      },
    },
  },
  plugins: [],
};
