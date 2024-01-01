<script lang="ts">
	import Container from '$lib/Container.svelte';
	import Header from '$lib/Header.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<Header />

<div
	class="absolute left-0 right-0 top-0 -z-10 h-screen w-screen bg-cover bg-center bg-no-repeat"
	style="
		background-image: 
			linear-gradient(to top, var(--fallback-b1,oklch(var(--b1))) 5%, var(--fallback-b1,oklch(var(--b1) / 0.3))),
			url('/static/movie/{data.id}.jpg');
	"
/>

<Container class="font-poppins mt-24 flex flex-col gap-5">
	<div class="flex h-56 w-full items-center overflow-hidden rounded-2xl bg-black/55 p-3 backdrop-blur-xl">
		<img src="/static/movie/{data.id}.jpg" alt={data.name} class="aspect-[3/4] h-full rounded-xl" />
		<div class="flex h-full flex-col justify-between p-5">
			<h2 class="card-title whitespace-nowrap text-2xl italic text-white sm:text-3xl">{data.name}</h2>

			<div class=" flex flex-col gap-3">
				<p>Time: {data.movieTime} minutes</p>
				<p>Genres: {data.genres.join(', ')}</p>
			</div>
		</div>
	</div>

	<div role="tablist" class="tabs tabs-boxed">
		<button role="tab" class="tab">Tab 1</button>
		<button role="tab" class="tab tab-active">Tab 2</button>
		<button role="tab" class="tab">Tab 3</button>
	</div>
</Container>
