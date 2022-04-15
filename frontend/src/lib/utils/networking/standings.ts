import axios from 'axios';
import type { StandingsData } from 'src/types/contestData';

export default async function Standings(): Promise<StandingsData[]> {
	return (await axios.get('/api/v1/standings')).data as StandingsData[];
}
