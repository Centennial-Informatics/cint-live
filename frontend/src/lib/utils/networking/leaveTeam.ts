import axios from 'axios';
import type { TeamData } from 'src/types/teamInfo';

export default async function LeaveTeam(token: string) {
	const formData = new FormData();
	formData.append('token', token);
	return (await axios.post('/api/v1/leave', formData)).data as TeamData;
}
