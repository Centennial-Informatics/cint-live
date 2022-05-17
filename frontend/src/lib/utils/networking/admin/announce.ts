import CLIENT from '../client';

export default async function MakeAnnouncement(token: string, title: string, details: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('title', title);
	formData.append('details', details);
	return (await CLIENT.post('/api/v1/admin/announce', formData)).data;
}
