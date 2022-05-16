import CLIENT from '../client';

export default async function AdminStandings(token: string) {
	const formData = new FormData();
	formData.append('token', token);
	return (await CLIENT.post('/api/v1/admin/standings', formData)).data;
}
