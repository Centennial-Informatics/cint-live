import type { TeamData } from 'src/types/teamInfo';
import CLIENT from './client';

export default async function JoinTeam(token: string, teamCode: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('team_code', teamCode);
	return (await CLIENT.post('/api/v1/join', formData)).data as TeamData;
}
