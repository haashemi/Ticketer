<script lang="ts">
	import Container from '$lib/Container.svelte';
	import Header from '$lib/Header.svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import ky from 'ky';
	import moment from 'moment';
	import * as z from 'zod';

	const schema = z.object({
		tickets: z.array(
			z.object({
				movieID: z.number(),
				ticketID: z.number(),
				movieName: z.string(),
				premiereDate: z.string().datetime(),
				premiereTime: z.string().datetime({ offset: true }),
			}),
		),
	});

	const query = createQuery({
		queryKey: ['tickets'],
		queryFn: async () => {
			const data = await ky.get('/api/profile/tickets').json();
			return schema.parse(data);
		},
	});
</script>

<Header />

<div class="fixed left-0 top-0 -z-10 h-screen w-full bg-gradient-to-b from-zinc-950 to-zinc-900" />

<Container class="font-poppins relative mb-36 mt-20 flex flex-col gap-5">
	<h1 class="mb-10 text-center">Purchased Tickets</h1>

	{#if $query.isLoading}
		<div class="grid grid-cols-1 gap-5 lg:grid-cols-3">
			{#each new Array(10) as _}
				<div class="flex h-20 w-full animate-pulse justify-between rounded-xl border border-zinc-700 bg-zinc-800/75 p-3" />
			{/each}
		</div>
	{:else if $query.isSuccess}
		{@const tickets = $query.data.tickets}
		{#if tickets.length === 0}
			<div class="text-center">You don't have any tickets reserved</div>
		{:else}
			<div class="grid grid-cols-1 gap-5 lg:grid-cols-3">
				{#each tickets as ticket}
					<div class="flex h-20 w-full justify-between rounded-xl border border-zinc-700 bg-zinc-800/75 p-3">
						<div class="flex flex-col justify-around">
							<p>{ticket.movieName}</p>
							<p>{moment(ticket.premiereDate).format('MMMM Do YYYY')}, {moment(ticket.premiereTime).format('HH:mm')}</p>
						</div>

						<a href="/ticket/{ticket.ticketID}" class="btn">View</a>
					</div>
				{/each}
			</div>
		{/if}
	{/if}
</Container>
