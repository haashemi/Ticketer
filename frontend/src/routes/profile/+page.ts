import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	return {
		tickets: [
			{ id: 1, movieName: 'The Lost Of The Whale Shark', premieresAt: new Date() },
			{ id: 1, movieName: 'The Lost Of The Whale Shark', premieresAt: new Date() },
			{ id: 1, movieName: 'The Lost Of The Whale Shark', premieresAt: new Date() },
			{ id: 1, movieName: 'The Lost Of The Whale Shark', premieresAt: new Date() },
			{ id: 1, movieName: 'The Lost Of The Whale Shark', premieresAt: new Date() },
			{ id: 1, movieName: 'The Lost Of The Whale Shark', premieresAt: new Date() },
			{ id: 1, movieName: 'The Lost Of The Whale Shark', premieresAt: new Date() },
			{ id: 1, movieName: 'The Lost Of The Whale Shark', premieresAt: new Date() },
		],
	};
};
