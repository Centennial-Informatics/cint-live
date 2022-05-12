import CLIENT from './client';

interface LoginInfo {
	error?: string;
	token: string;
	team_id: string;
}

export default async function Login(googleIdToken: string) {
	const formData = new FormData();
	formData.append('id_token', googleIdToken);
	return (await CLIENT.post('/api/v1/login', formData)).data as LoginInfo;
}
