import type { ContestData } from 'src/types';
import CLIENT from './client';

export default async function CollectAdvanced() {
	const data = (await CLIENT.get('/api/v1/collect/advanced')).data;
	return data as ContestData;
}
