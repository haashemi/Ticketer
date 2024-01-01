import ky from 'ky';
import type { PageServerLoad } from './$types';
import * as z from 'zod';
import { API_HOST } from '$env/static/private';

const schema = z.object({
	movies: z.array(
		z.object({
			id: z.number(),
			name: z.string(),
			premiereTime: z.number(),
		}),
	),
});

export const load: PageServerLoad = async () => {
	const resp = await ky.get(API_HOST + '/api/public/movies').json();
	const data = schema.parse(resp);

	return data;
};
