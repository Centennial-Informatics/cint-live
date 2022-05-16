import CLIENT from '../client';

export default async function DownloadStandings(token: string, division: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('division', division);
	return (await CLIENT.post('/api/v1/admin/download', formData)).data;
}
