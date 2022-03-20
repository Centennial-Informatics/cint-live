import axios from 'axios';
import type { ContestData } from 'src/types';

export default async function Collect() {
	const data = (await axios.post('/api/v1/collect')).data;
	return data as ContestData;
}
