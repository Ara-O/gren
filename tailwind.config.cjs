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
        gren: "#0F2D14",
        bestSellingProductItem: "#E8E8E8",
      },
    },
  },
  plugins: [],
};
