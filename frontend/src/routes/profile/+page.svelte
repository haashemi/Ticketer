<script lang="ts">
	import Container from '$lib/Container.svelte';
	import moment from 'moment';
	import type { PageData } from './$types';
	import Header from '$lib/Header.svelte';

	export let data: PageData;
</script>

<Header />

<div class="fixed left-0 top-0 -z-10 h-screen w-full bg-gradient-to-b from-zinc-950 to-zinc-900" />

<Container class="font-poppins relative mb-36 mt-20 flex flex-col gap-5">
	<h1 class="mb-10 text-center">Purchased Tickets</h1>

	<div class="grid grid-cols-1 gap-5 lg:grid-cols-3">
		{#each data.tickets as ticket}
			<div class="flex w-full justify-between rounded-xl border border-zinc-700 bg-zinc-800/75 p-3">
				<div class="flex flex-col justify-around">
					<p>{ticket.movieName}</p>
					<p>{moment(ticket.premieresAt).format('MMMM Do YYYY, HH:mm')}</p>
				</div>

				<a href="/ticket/{ticket.id}" class="btn">View</a>
			</div>
		{/each}
	</div>
</Container>
