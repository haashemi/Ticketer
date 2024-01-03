import ky from 'ky';
import type { PageServerLoad } from './$types';
import * as z from 'zod';
import { API_HOST } from '$env/static/private';

const schema = z.object({
	id: z.number(),
	name: z.string(),
	movieTime: z.number(),
	genres: z.array(z.string()),
	fromDate: z.coerce.date(), // todo: date
	toDate: z.coerce.date(), // todo: date
	premiereTime: z.number(), // todo: date
});

export const load: PageServerLoad = async ({ params }) => {
	const resp = await ky.get(API_HOST + `/api/public/movies/${params.id}`).json();
	const data = schema.parse(resp);
	return data;
};
