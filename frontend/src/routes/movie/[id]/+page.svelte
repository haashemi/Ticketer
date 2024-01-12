<script lang="ts">
	import { goto } from '$app/navigation';
	import Container from '$lib/Container.svelte';
	import Header from '$lib/Header.svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import ky from 'ky';
	import moment from 'moment';
	import * as z from 'zod';
	import type { PageData } from './$types';
	export let data: PageData;

	let currentDate: Date = getDateAfter(1);
	let selectedSeats: number[] = [];
	let isLoading: boolean = false;

	const schema = z.object({ reservedSeats: z.array(z.number()) });

	const query = createQuery({
		queryKey: ['reservedSeats', data.id, currentDate],
		queryFn: async () => {
			const resp = await ky.get(`/api/public/movies/${data.id}/reserved-seats/${moment(currentDate).format('YYYY/MM/DD')}`).json();
			return schema.parse(resp);
		},
		refetchInterval: 30_000,
	});

	async function reserveSeats() {
		isLoading = true;

		const resp = await ky
			.post('/api/profile/tickets/reserve', {
				json: {
					movieId: data.id,
					forYear: currentDate.getFullYear(),
					forMonth: currentDate.getMonth() + 1,
					forDay: currentDate.getDate(),
					seats: selectedSeats,
				},
			})
			.json<{ ticketId: number }>();

		goto('/ticket/' + resp.ticketId);
	}

	function getDateAfter(day: number): Date {
		const d = new Date();
		d.setDate(new Date().getDate() + day);
		return d;
	}

	function selectDate(date: Date) {
		currentDate = date;
		selectedSeats = [];
		$query.refetch();
	}

	function toggleSeat(seatId: number) {
		const index = selectedSeats.indexOf(seatId);
		console.log(seatId, index);
		if (index > -1) {
			selectedSeats.splice(index, 1);
			selectedSeats = [...selectedSeats]; //just make a copy
		} else {
			selectedSeats = [...selectedSeats, seatId];
		}
	}
</script>

<Header />

<div
	class="absolute left-0 right-0 top-0 -z-10 h-2/3 w-full bg-cover bg-center bg-no-repeat"
	style="
		background-image: 
			linear-gradient(to top, var(--fallback-b1,oklch(var(--b1))) 0%, var(--fallback-b1,oklch(var(--b1) / 0.3))),
			url('/static/movie/{data.id}.jpg');
	"
/>

<Container class="font-poppins my-24 flex flex-col gap-5">
	<div class="flex h-56 w-full items-center overflow-hidden rounded-xl bg-zinc-900/55 p-3 backdrop-blur-xl">
		<img src="/static/movie/{data.id}.jpg" alt={data.name} class="aspect-[3/4] h-full rounded-xl" />
		<div class="flex h-full flex-col justify-between p-5">
			<h2 class="card-title whitespace-nowrap text-2xl italic text-white sm:text-3xl">{data.name}</h2>

			<div class=" flex flex-col gap-3">
				<p>Time: {data.movieTime} minutes</p>
				<p>Genres: {data.genres.join(', ')}</p>
			</div>
		</div>
	</div>

	<fieldset
		disabled={isLoading || $query.fetchStatus === 'fetching'}
		class="flex min-w-0 gap-4 overflow-x-auto rounded-xl bg-zinc-900/55 p-4 backdrop-blur-xl"
	>
		{#each new Array(7) as _, idx}
			{@const day = getDateAfter(idx + 1)}
			{@const isSelected = currentDate.getDate() === day.getDate()}
			<button class="btn btn-primary {isSelected ? '' : 'btn-outline'}" on:click={() => selectDate(day)}>{moment(day).format('MMM Do')}</button>
		{/each}
	</fieldset>

	<div class="flex items-center justify-center rounded-xl bg-zinc-900/55 p-6 backdrop-blur-xl">
		<fieldset disabled={isLoading || $query.fetchStatus === 'fetching'} class="grid w-full max-w-[40rem] grid-cols-8 grid-rows-10 gap-3">
			{#each new Array(120) as _, idx}
				{@const isReserved = $query.data?.reservedSeats.includes(idx)}
				{@const isSelected = selectedSeats.includes(idx)}
				<button disabled={isReserved} on:click={() => toggleSeat(idx)} class="btn btn-primary h-full w-full {isSelected ? '' : 'btn-outline'}">
					{idx + 1}
				</button>
			{/each}
		</fieldset>
	</div>

	{#if selectedSeats.length > 0}
		<div class="sticky bottom-0 flex items-center justify-around rounded-xl bg-zinc-900/55 p-6 backdrop-blur-xl">
			<p>Total Price: {(selectedSeats.length * 65_000).toLocaleString('en')} IRT</p>
			<button disabled={isLoading} on:click={() => reserveSeats()} class="btn btn-accent">Reserve</button>
		</div>
	{/if}
</Container>
