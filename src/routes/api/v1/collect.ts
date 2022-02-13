import axios from 'axios';

export async function post() {
	console.log('collect');
	return {
		body: (await axios.post('http://localhost:8000/api/v1/collect')).data
	};
}
