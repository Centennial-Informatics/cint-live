import CLIENT from '../client';

export default async function UpdateEnd(token: string, time: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('end_time', time);
	return (await CLIENT.post('/api/v1/admin/updateend', formData)).data;
}
