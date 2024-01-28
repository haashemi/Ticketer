<script lang="ts">
	import { page } from '$app/stores';
	import Container from '$lib/Container.svelte';
	import Header from '$lib/Header.svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import ky from 'ky';
	import moment from 'moment';
	import * as z from 'zod';

	const schema = z.object({
		movieId: z.number(),
		movieName: z.string(),
		movieTime: z.number(),
		movieGenres: z.array(z.string()),
		reservedAt: z.string().datetime({ offset: true, precision: 6 }),
		premiereDate: z.string().datetime(),
		premiereTime: z.string().datetime({ offset: true }),
		reservedSeats: z.array(z.number()),
	});

	const query = createQuery({
		queryKey: ['ticket', $page.params.id],
		queryFn: async () => {
			const resp = await ky.get(`/api/profile/ticket/${$page.params.id}`).json();
			return schema.parse(resp);
		},
	});
</script>

<Header />

<div class="fixed left-0 top-0 -z-10 h-screen w-full bg-gradient-to-b from-zinc-950 to-zinc-900" />

<Container class="font-poppins my-16 flex flex-col gap-5">
	<div class="grid gap-5 lg:grid-cols-2">
		{#if $query.isSuccess}
			<div class="flex h-56 w-full items-center overflow-hidden rounded-xl bg-zinc-900 p-3">
				<img src="/static/movie/{$query.data.movieId}.jpg" alt={$query.data.movieName} class="aspect-[3/4] h-full rounded-xl" />
				<div class="flex h-full flex-col justify-between p-5">
					<h2 class="card-title whitespace-nowrap text-2xl italic text-white sm:text-3xl">{$query.data.movieName}</h2>

					<div class="flex flex-col gap-3">
						<p>Time: {$query.data.movieTime} minutes</p>
						<p>Genres: {$query.data.movieGenres.join(', ')}</p>
					</div>
				</div>
			</div>

			<div class="flex h-56 w-full flex-col items-center justify-around rounded-xl bg-zinc-900 p-3">
				<div class="flex flex-col gap-5">
					<p>
						Movie Premieres At: {moment($query.data.premiereDate).format('MMMM Do YYYY')}, {moment($query.data.premiereTime).format('HH:mm')}
					</p>
					<p>Ticket bought at: {moment($query.data.reservedAt).format('MMMM Do YYYY, HH:mm')}</p>
					<p>Money Paid: {($query.data.reservedSeats.length * 65_000).toLocaleString('en')} IRT</p>
				</div>
			</div>
		{:else}
			<div class="flex h-56 w-full animate-pulse rounded-xl bg-zinc-900 p-3" />
			<div class="flex h-56 w-full animate-pulse rounded-xl bg-zinc-900 p-3" />
		{/if}
	</div>

	<div class="flex flex-col items-center justify-center gap-5 rounded-xl bg-zinc-900/55 p-6 backdrop-blur-xl">
		<h5>Reserved Seats:</h5>
		<div class="grid w-full max-w-[40rem] grid-cols-8 grid-rows-10 gap-3">
			{#if $query.isSuccess}
				{#each new Array(120) as _, idx}
					<button class="btn btn-primary h-full w-full {$query.data.reservedSeats.includes(idx) ? '' : 'btn-outline'}">{idx + 1}</button>
				{/each}
			{:else}
				{#each new Array(120) as _, idx}
					<button disabled class="btn btn-primary btn-outline h-full w-full">{idx + 1}</button>
				{/each}
			{/if}
		</div>
	</div>
</Container>
