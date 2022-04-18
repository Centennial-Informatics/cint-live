import axios from 'axios';
import type { ContestData } from 'src/types';

export default async function CollectStandard() {
	const data = (await axios.post('/api/v1/collect/standard')).data;
	return data as ContestData;
}
