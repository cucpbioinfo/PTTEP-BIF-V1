module.exports = {
  purge: [
    './src/components/**/*.{js,ts,jsx,tsx}',
    './src/features/**/*.{js,ts,jsx,tsx}',
    './pages/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    fontFamily: {
      sans: ['IBM Plex Sans Thai', 'sans-serif'],
      serif: ['IBM Plex Sans Thai', 'serif'],
    },
    extend: {
      colors: {
        gray: {
          900: '#747474',
          800: '#919191',
          700: '#AEAEAE',
          600: '#CBCBCB',
          500: '#E8E8E8',
          400: '#EDEDED',
          300: '#F1F1F1',
          200: '#F6F6F6',
          100: '#FAFAFA',
        },
        primary: "#126e82",
        subprimary:"#25869c",
        secondary: "#51c4d3",
        tertiary: "#d8e3e7",
        "secondary-dark":"#132c33"
      },
      spacing: {
        '96': '24rem',
        '128': '32rem',
      }
    },
    variants: {
      borderStyle: ['hover'],
      opacity: ['responsive', 'hover', 'focus', 'active', 'group-hover'],
      outline: ['focus'],
    },
    plugins: [],
  }
}
