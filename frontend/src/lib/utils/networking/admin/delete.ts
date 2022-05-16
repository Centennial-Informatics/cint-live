import CLIENT from '../client';

export default async function DeleteUser(token: string, email: string) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('email', email);
	return (await CLIENT.post('/api/v1/admin/dangerous/delete', formData)).data;
}
