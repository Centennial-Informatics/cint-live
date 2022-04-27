import axios from 'axios';

// @ts-expect-error Binding element 'request' implicitly has an 'any' type.
export async function post({ request }) {
	const form = await request.formData();
	const formUrlEncoded = (x: { [key: string]: string }) =>
		Object.keys(x).reduce((p, c) => p + `&${c}=${encodeURIComponent(x[c])}`, '');
	return {
		body: (
			await axios.post(
				process.env.HOST + '/api/v1/join',
				formUrlEncoded({
					token: form.get('token'),
					team_code: form.get('team_code')
				})
			)
		).data
	};
}
