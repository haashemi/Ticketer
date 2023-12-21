<script>
	import Container from '$lib/Container.svelte';
	import Header from '$lib/Header.svelte';
</script>

<Header />

<Container class="flex py-10">
	<div
		class="flex h-72 w-full items-end justify-center overflow-hidden rounded-2xl bg-base-200 bg-cover"
		style="background-image: url('https://cataas.com/cat/cute');"
	>
		<h2 class="font-poppins card-title whitespace-nowrap py-10 text-2xl italic text-white sm:text-3xl">My Long As Shit Event Name</h2>
	</div>
</Container>
