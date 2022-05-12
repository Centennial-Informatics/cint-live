import axios from 'axios';

export async function get() {
	return {
		body: (await axios.get(process.env.HOST + '/api/v1/collect/advanced')).data
	};
}
