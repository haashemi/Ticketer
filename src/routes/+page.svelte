<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<header class="sticky top-0 z-50 w-full bg-black/10 backdrop-blur-lg">
	<header class="container mx-auto flex h-16 items-center justify-between px-5">
		<div class="flex gap-2">
			<a href="/signin" class="btn btn-secondary">Login</a>
		</div>

		<a href="/" class="font-silk text-2xl font-bold text-white">Ticketer</a>
	</header>
</header>

<main class="container mx-auto px-5">
	<!-- Header -->
	<div class="my-14 flex w-full flex-col items-center gap-3">
		<h2 class="text-4xl font-semibold italic text-slate-50">#1 Event ticket store</h2>
		<p class="text-sm text-slate-500">Get the newest events' ticket from Ticketer!</p>
	</div>

	<!-- Events list -->
	<h3 class="font-poppins my-3 text-lg italic">Latest Events:</h3>
	<div class="grid grid-cols-1 gap-8 lg:grid-cols-2 xl:grid-cols-3">
		{#each data.events as event}
			<a
				href="/event/{event.id}"
				class="card aspect-video overflow-hidden bg-base-200 bg-cover shadow-md transition-transform ease-in-out hover:-translate-y-2"
				style="background-image: url('https://cataas.com/cat/cute?id={event.id}');"
			>
				<div class="card-body flex flex-col justify-end gap-4 bg-gradient-to-b from-transparent to-black text-white">
					<h2 class="font-poppins card-title whitespace-nowrap text-2xl italic sm:text-3xl">{event.name}</h2>

					<div class="card-actions items-center justify-between">
						<p class="text-base text-accent sm:text-lg">{event.price} IRT</p>
						<button class="btn btn-accent btn-sm">Book a Ticket</button>
					</div>
				</div>
			</a>
		{/each}
	</div>
</main>

<footer class="mt-10 flex h-11 items-center justify-center border-t border-base-300">Copyright Ticketer, no rights reserved at all.</footer>
