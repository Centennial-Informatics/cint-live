import { BASE_URL } from '$lib/data/constants/url';
import axios from 'axios';

const CLIENT = axios.create({
	baseURL: BASE_URL[process.env.NODE_ENV]
});

export default CLIENT;
