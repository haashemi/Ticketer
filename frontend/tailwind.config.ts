/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				silk: ['Silkscreen', 'sans-serif'],
				poppins: ['Poppins', 'sans-serif'],
			},
		},
	},

	daisyui: {
		themes: [
			{
				light: {
					primary: '#302470',
					secondary: '#8776db',
					accent: '#3f26c0',
					neutral: '#1b1538',
					'base-100': '#eeecf8',
				},
				dark: {
					primary: '#9a8eda',
					secondary: '#342489',
					accent: '#5940d9',
					neutral: '#1b1538',
					'base-100': '#070708',
				},
			},
		],

		darkTheme: 'dark', // name of one of the included themes for dark mode
		base: true, // applies background color and foreground color for root element by default
		styled: true, // include daisyUI colors and design decisions for all components
		utils: true, // adds responsive and modifier utility classes
		logs: true, // Shows info about daisyUI version and used config in the console when building your CSS
	},

	plugins: [require('daisyui')],
};
