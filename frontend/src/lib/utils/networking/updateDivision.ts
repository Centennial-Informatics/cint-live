import CLIENT from './client';

export default async function UpdateTeamDivision(token: string, teamDivision: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('division', teamDivision);
	return (await CLIENT.post('/api/v1/division', formData)).data as string;
}
