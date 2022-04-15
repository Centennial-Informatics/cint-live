import axios from 'axios';
import type { TeamData } from 'src/types/teamInfo';

export default async function GetTeam(token: string) {
	const formData = new FormData();
	formData.append('token', token);
	return (await axios.post('/api/v1/team', formData)).data as TeamData;
}
