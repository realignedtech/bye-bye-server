/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte}', './index.html'],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#f0f7fd',
          100: '#e0eefa',
          200: '#c7e1f5',
          300: '#a0cdee',
          400: '#5ba3d9',
          500: '#3b8fce',
          600: '#2e7dbd',
          700: '#013373',
          800: '#012d66',
          900: '#012255',
        },
      },
      fontFamily: {
        sans: [
          '"Inter"',
          '-apple-system',
          'BlinkMacSystemFont',
          '"Segoe UI"',
          'Roboto',
          'sans-serif',
        ],
      },
    },
  },
};
