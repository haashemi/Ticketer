<script lang="ts">
	import Container from '$lib/Container.svelte';
	import Header from '$lib/Header.svelte';
	import moment from 'moment';
	import type { PageData } from './$types';

	export let data: PageData;

	function getDateAfter(day: number): Date {
		const d = new Date();
		d.setDate(new Date().getDate() + day);
		return d;
	}

	let currentDate: Date = getDateAfter(1);
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

	<div class="flex gap-4 overflow-x-auto rounded-xl bg-zinc-900/55 p-4 backdrop-blur-xl">
		{#each new Array(7) as _, idx}
			{@const day = getDateAfter(idx + 1)}
			{@const isSelected = currentDate.getDate() === day.getDate()}
			<button class="btn btn-primary {isSelected ? '' : 'btn-outline'}" on:click={() => (currentDate = day)}
				>{moment(day).format('MMM Do')}</button
			>
		{/each}
	</div>

	<div class="flex items-center justify-center rounded-xl bg-zinc-900/55 p-6 backdrop-blur-xl">
		<div class="grid w-full max-w-[40rem] grid-cols-8 grid-rows-10 gap-3">
			{#each new Array(120) as _}
				<button class="btn btn-primary btn-outline h-full w-full">💺</button>
			{/each}
		</div>
	</div>
</Container>
