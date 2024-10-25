/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        gray: '#282828',
        border: '#333333',
        text: '#DDDDDD'
      },
      backgroundColor: {
        gray: '#282828',
        border: '#333333',
        text: '#DDDDDD'
      }
    }
  },
  plugins: []
}
