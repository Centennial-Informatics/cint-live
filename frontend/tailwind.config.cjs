module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	darkMode: 'class', // or 'media' or 'class'
	theme: {
		extend: {
			colors: {
				// brand: {
				// 	lightest: '#b3cde1',
				// 	light: '#4d89b9',
				// 	DEFAULT: '#00579b',
				// 	dark: '#003c7f',
				// 	darkest: '#002c6d'
				// },
				brand: {
					lightest: '#BBBEEE',
					light: '#8E93FF',
					DEFAULT: '#4A5DFE',
					dark: '#3814C5',
					darkest: '#271190'
				},
				alt: {
					lightest: '#FFD156',
					light: '#F5A727',
					DEFAULT: '#ED8B3B',
					dark: '#C5591C',
					darkest: '#8F371B'
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
				},
				good: {
					0: '#86efac',
					50: '#4ade80',
					100: '#166534'
				},
				bad: {
					0: '#fca5a5',
					50: '#f87171',
					100: '#991b1b'
				},
				neutral: {
					0: '#e2e8f0',
					50: '#64748b',
					100: '#1f2937'
				},
				warning: {
					0: '#fde68a',
					50: '#f59e0b',
					100: '#92400e'
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
			},
			transitionProperty: {
				width: 'width'
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
