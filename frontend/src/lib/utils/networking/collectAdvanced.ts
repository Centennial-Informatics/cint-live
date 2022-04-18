import axios from 'axios';
import type { ContestData } from 'src/types';

export default async function CollectAdvanced() {
	const data = (await axios.post('/api/v1/collect/advanced')).data;
	return data as ContestData;
}
