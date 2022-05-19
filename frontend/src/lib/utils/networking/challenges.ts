import CLIENT from './client';

export default async function CollectChallenges() {
	const data = (await CLIENT.get('/api/v1/challenges')).data;
	return data as string[];
}
