import CLIENT from '../client';

export default async function UpdateStart(token: string, time: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('start_time', time);
	console.log(time);
	return (await CLIENT.post('/api/v1/admin/updatestart', formData)).data;
}
