import type { ContestData } from 'src/types';
import CLIENT from './client';

export default async function CollectStandard() {
	const data = (await CLIENT.get('/api/v1/collect/standard')).data;
	return data as ContestData;
}
