/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,ts}"],
  theme: {
    extend: {
      height: {
        "90vh": "90vh",
      },
      minHeight: {
        "9/10": "90%",
      },
      boxShadow: {
        paper:
          "0 -1px 1px rgba(0,0,0,0.15), 0 -10px 0 -5px rgb(249 250 251), 0 -10px 1px -4px rgba(0,0,0,0.15), 0 -20px 0 -10px rgb(249 250 251), 0 -20px 1px -9px rgba(0,0,0,0.15);",
      },
    },
  },
  plugins: [],
};
