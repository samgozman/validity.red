/* eslint-disable no-undef */
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{vue,js,ts}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      {
        "validity-light": {
          primary: "#e11d48",
          secondary: "#fb7185",
          accent: "#4b5563",
          neutral: "#191D24",
          "base-100": "#FFFFFF",
          "base-200": "#EBF1F4",
          info: "#3ABFF8",
          success: "#36D399",
          warning: "#FBBD23",
          error: "#F87272",
        },
      },
      {
        "validity-dark": {
          primary: "#e11d48",
          secondary: "#fb7185",
          accent: "#4b5563",
          neutral: "#191D24",
          "base-100": "#2e3b51",
          "base-200": "#252e41",
          info: "#3ABFF8",
          success: "#36D399",
          warning: "#FBBD23",
          error: "#F87272",
        },
      },
    ],
    darkTheme: "validity-dark",
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
};
