import axios from 'axios';

export default async function UpdateTeamName(token: string, teamName: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('team_name', teamName);
	return (await axios.post('/api/v1/update', formData)).data as string;
}
