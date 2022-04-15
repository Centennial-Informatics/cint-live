import axios from 'axios';
import type { TeamData } from 'src/types/teamInfo';

export default async function JoinTeam(token: string, teamCode: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('team_code', teamCode);
	return (await axios.post('/api/v1/join', formData)).data as TeamData;
}
