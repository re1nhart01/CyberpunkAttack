import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        "space-armor": ["var(--font-space-armor)"],
        orbitron: ["var(--font-orbitron)"],
      },
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        black: "#000000",
      },
    },
  },
  plugins: [],
};
export default config;
