import axios from 'axios';

export async function get() {
	return {
		body: (await axios.get('http://localhost:8000/api/v1/standings')).data
	};
}
