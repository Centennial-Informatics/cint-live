import CLIENT from '../client';

export default async function UpdateCache(token: string) {
	const formData = new FormData();
	formData.append('token', token);
	return (await CLIENT.post('/api/v1/admin/updatecache', formData)).data;
}
