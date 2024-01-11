<script lang="ts">
	import Container from '$lib/Container.svelte';
	import Header from '$lib/Header.svelte';
	import moment from 'moment';

	const test = {
		movieId: 1,
		movieName: 'Residency',
		movieTime: 65,
		movieGenres: ['Documentary'],
		reservedAt: new Date(),
		premiereDate: new Date(),
		premiereTime: new Date(),
		reservedSeats: [6, 7, 8],
	};
</script>

<Header />

<div class="fixed left-0 top-0 -z-10 h-screen w-full bg-gradient-to-b from-zinc-950 to-zinc-900" />

<Container class="font-poppins my-16 flex flex-col gap-5">
	<div class="grid gap-5 lg:grid-cols-2">
		<div class="flex h-56 w-full items-center overflow-hidden rounded-xl bg-zinc-900 p-3">
			<img src="/static/movie/{test.movieId}.jpg" alt={test.movieName} class="aspect-[3/4] h-full rounded-xl" />
			<div class="flex h-full flex-col justify-between p-5">
				<h2 class="card-title whitespace-nowrap text-2xl italic text-white sm:text-3xl">{test.movieName}</h2>

				<div class="flex flex-col gap-3">
					<p>Time: {test.movieTime} minutes</p>
					<p>Genres: {test.movieGenres.join(', ')}</p>
				</div>
			</div>
		</div>

		<div class="flex h-56 w-full flex-col items-center justify-around rounded-xl bg-zinc-900 p-3">
			<div class="flex flex-col gap-5">
				<p>Movie Premieres At: {moment(test.premiereDate).format('MMMM Do YYYY')}, {moment(test.premiereTime).format('HH:mm')}</p>
				<p>Ticket bought at: {moment(test.reservedAt).format('MMMM Do YYYY, HH:mm')}</p>
				<p>Money Paid: {(test.reservedSeats.length * 65_000).toLocaleString('en')} IRT</p>
			</div>
		</div>
	</div>

	<div class="flex flex-col items-center justify-center gap-5 rounded-xl bg-zinc-900/55 p-6 backdrop-blur-xl">
		<h5>Reserved Seats:</h5>
		<div class="grid w-full max-w-[40rem] grid-cols-8 grid-rows-10 gap-3">
			{#each new Array(120) as _, idx}
				<button class="btn btn-primary h-full w-full {test.reservedSeats.includes(idx) ? '' : 'btn-outline'}">{idx + 1}</button>
			{/each}
		</div>
	</div>
</Container>
