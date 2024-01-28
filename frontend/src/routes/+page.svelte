<script lang="ts">
	import Container from '$lib/Container.svelte';
	import Header from '$lib/Header.svelte';
	import moment from 'moment';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<Header />

<div
	class="container absolute left-0 right-0 top-0 -z-10 mx-auto h-2/3 w-full bg-cover bg-bottom bg-no-repeat"
	style="
		background-image: 
			linear-gradient(to right, var(--fallback-b1,oklch(var(--b1))), transparent 70%, var(--fallback-b1,oklch(var(--b1)))), 
			linear-gradient(to top, var(--fallback-b1,oklch(var(--b1))), transparent),
			url('/theater.png');
	"
/>

<Container class="font-poppins">
	<!-- Header -->
	<div class="flex w-full flex-col items-start justify-center gap-5 py-24 md:py-40">
		<h2 class="text-3xl italic text-slate-50 md:text-5xl">#1 Ticket marketplace</h2>
		<p class="pl-3 text-sm text-slate-500">Get the newest movies' ticket from Ticketer!</p>
	</div>

	<!-- Movies list -->
	<h3 class="mb-7 text-lg italic">This week's movies:</h3>
	<div class="grid grid-cols-2 gap-8 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6">
		{#each data.movies as movie}
			<div class="transition-transform duration-500 ease-in-out hover:-translate-y-2">
				<a
					href="/movie/{movie.id}"
					class="bg-base-200 flex aspect-[3/4] items-end justify-center overflow-hidden rounded-2xl bg-cover"
					style="background-image: linear-gradient(to top, black, transparent 60%), url('/static/movie/{movie.id}.jpg');"
				>
					<h6 class="overflow-hidden text-ellipsis whitespace-nowrap rounded-xl p-2 text-lg font-medium text-white">
						{movie.name}
					</h6>
				</a>

				<p class="px-3"><span class="text-sm opacity-50">Premieres At:</span> {moment(movie.premiereTime).format('HH:mm')}</p>
			</div>
		{/each}
	</div>
</Container>

<footer class="border-base-300 mt-10 flex h-11 items-center justify-center border-t">Copyright Ticketer, no rights reserved at all.</footer>
