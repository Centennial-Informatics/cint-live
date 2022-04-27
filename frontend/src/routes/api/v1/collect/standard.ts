import axios from 'axios';

export async function post() {
	return {
		body: (await axios.post(process.env.HOST + '/api/v1/collect/standard')).data
	};
}
