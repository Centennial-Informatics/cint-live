import CLIENT from './client';

export default async function UpdateTeamName(token: string, teamName: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('team_name', teamName);
	return (await CLIENT.post('/api/v1/update', formData)).data as string;
}
