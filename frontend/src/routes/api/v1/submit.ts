import axios from 'axios';

// @ts-expect-error Binding element 'request' implicitly has an 'any' type.
export async function post({ request }) {
	const form = await request.formData();
	const formUrlEncoded = (x: { [key: string]: string }) =>
		Object.keys(x).reduce((p, c) => p + `&${c}=${encodeURIComponent(x[c])}`, '');
	return {
		body: (
			await axios.post(
				'http://localhost:8000/api/v1/submit',
				formUrlEncoded({
					token: form.get('token'),
					problem: form.get('problem'),
					language: form.get('language'),
					submission: form.get('submission'),
					file: form.get('file')
				})
			)
		).data
	};
}
