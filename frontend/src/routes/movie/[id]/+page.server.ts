import { API_HOST } from '$env/static/private';
import ky from 'ky';
import * as z from 'zod';
import type { PageServerLoad } from './$types';

const schema = z.object({
	id: z.number(),
	name: z.string(),
	movieTime: z.number(),
	genres: z.array(z.string()),
	fromDate: z.coerce.date(), // todo: date
	toDate: z.coerce.date(), // todo: date
	premiereTime: z.string().datetime({ offset: true }), // todo: date
});

export const load: PageServerLoad = async ({ params }) => {
	const resp = await ky.get(API_HOST + `/api/public/movies/${params.id}`).json();
	const data = schema.parse(resp);
	return data;
};
