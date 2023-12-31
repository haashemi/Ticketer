import ky from 'ky';
import type { PageServerLoad } from './$types';
import * as z from 'zod';
import { API_HOST } from '$env/static/private';

const schema = z.object({
	movies: z.array(
		z.object({
			id: z.number(),
			name: z.string(),
			showTime: z.number(),
		}),
	),
});

export const load: PageServerLoad = async () => {
	console.log(API_HOST + '/api/public/movies');
	const resp = await ky.get(API_HOST + '/api/public/movies').json();
	console.log(resp);
	const data = schema.parse(resp);

	return data;
};
