export const BASE_URL = {
	production: 'http://' + import.meta.env.VITE_BACKEND_HOST, //'https://live-judge.herokuapp.com',
	development: 'http://localhost:8000'
};

export const WS_URL = {
	production: 'wss://' + import.meta.env.VITE_BACKEND_HOST + '/ws',
	development: 'ws://localhost:8000/ws'
};
