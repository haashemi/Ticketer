import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	return {
		movies: [
			{ id: 1, name: 'Movie Name' },
			{ id: 2, name: 'Movie Name' },
			{ id: 3, name: 'Movie Name' },
			{ id: 4, name: 'Movie Name' },
			{ id: 5, name: 'Movie Name' },
			{ id: 6, name: 'Movie Name' },
			{ id: 7, name: 'Movie Name' },
			{ id: 8, name: 'Movie Name' },
			{ id: 9, name: 'Movie Name' },
			{ id: 10, name: 'Movie Name' },
			{ id: 11, name: 'Movie Name' },
			{ id: 12, name: 'Movie Name' },
			{ id: 13, name: 'Movie Name' },
			{ id: 14, name: 'Movie Name' },
			{ id: 15, name: 'Movie Name' },
		],
	};
};
