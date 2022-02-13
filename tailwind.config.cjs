module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	darkMode: 'class', // or 'media' or 'class'
	theme: {
		extend: {
			colors: {
				brand: {
					lightest: '#b3cde1',
					light: '#4d89b9',
					DEFAULT: '#00579b',
					dark: '#003c7f',
					darkest: '#002c6d'
				},
				gray: {
					50: '#f8fafc',
					100: '#f1f5f9',
					200: '#e2e8f0',
					300: '#cbd5e1',
					400: '#94a3b8',
					500: '#64748b',
					600: '#475569',
					700: '#314156',
					800: '#1f2937',
					900: '#0f172a'
				}
				// brand: {
				//   lightest: "#f9c4c7",
				//   light: "#f59da2",
				//   DEFAULT: "#ea3b44",
				//   dark: "#e1262d",
				//   darkest: "#db191f",
				// },
			},
			minHeight: {
				64: '16rem'
			}
		},
		fontFamily: {
			sans: ['Inter', 'Roboto', 'sans-serif'],
			body: ['Roboto', 'Open Sans', 'sans-serif'],
			mono: ['Fira Mono', 'monospace']
		}
	},
	variants: {
		extend: {}
	},
	plugins: []
};
