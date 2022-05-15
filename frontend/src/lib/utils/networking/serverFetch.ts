/* Typed fetch api wrapper. Via https://www.carlrippon.com/fetch-with-async-await-and-typescript */
export default async function fetchType<T>(request: RequestInfo): Promise<T> {
	return await (await fetch(request)).json();
}
