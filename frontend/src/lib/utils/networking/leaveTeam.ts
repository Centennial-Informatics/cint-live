import type { TeamData } from 'src/types/teamInfo';
import CLIENT from './client';

export default async function LeaveTeam(token: string) {
	const formData = new FormData();
	formData.append('token', token);
	return (await CLIENT.post('/api/v1/leave', formData)).data as TeamData;
}
