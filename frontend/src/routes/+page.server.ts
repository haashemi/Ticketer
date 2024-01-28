import { API_HOST } from '$env/static/private';
import ky from 'ky';
import * as z from 'zod';
import type { PageServerLoad } from './$types';

const schema = z.object({
	movies: z.array(
		z.object({
			id: z.number(),
			name: z.string(),
			premiereTime: z.string().datetime({ offset: true }),
		}),
	),
});

export const load: PageServerLoad = async () => {
	const resp = await ky.get(API_HOST + '/api/public/movies').json();
	const data = schema.parse(resp);

	return data;
};
