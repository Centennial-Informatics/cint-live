import type { StandingsData } from 'src/types/contestData';
import CLIENT from './client';

export default async function Standings(): Promise<StandingsData[]> {
	return (await CLIENT.get('/api/v1/standings')).data as StandingsData[];
}
