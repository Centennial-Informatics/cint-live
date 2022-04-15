import axios from 'axios';

interface LoginInfo {
	token: string;
	team_id: string;
}

export default async function Login(googleIdToken: string) {
	const formData = new FormData();
	formData.append('id_token', googleIdToken);
	return (await axios.post('/api/v1/login', formData)).data as LoginInfo;
}
