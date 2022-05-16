import CLIENT from '../client';

export default async function CreateUser(
	token: string,
	email: string,
	name: string,
	division: string
) {
	const formData = new FormData();
	formData.append('token', token);
	formData.append('email', email);
	formData.append('name', name);
	formData.append('division', division);
	return (await CLIENT.post('/api/v1/admin/create', formData)).data;
}
