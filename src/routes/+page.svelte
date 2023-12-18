<header class="sticky top-0 z-50 w-full bg-black/10 backdrop-blur-lg">
	<header class="container mx-auto flex h-16 items-center justify-between px-5">
		<div class="flex gap-2">
			<a href="/signin" class="btn btn-secondary">Login</a>
		</div>

		<h1 class="text-xl font-bold text-white">Ticketer</h1>
	</header>
</header>

<main class="container mx-auto">
	<div class="my-14 flex w-full flex-col items-center gap-3">
		<h2 class="text-4xl font-semibold italic">#1 Event ticket store</h2>
		<p class="text-sm text-slate-500">Get the newest events' ticket from Ticketer!</p>
	</div>

	<h3 class="my-3 text-lg italic">Latest Events:</h3>

	<div class="grid grid-cols-3 gap-5">
		{#each Array(20) as idx}
			<div class="card bg-base-200 shadow-md">
				<figure><img src="https://cataas.com/cat/cute" class="aspect-video" alt="Cat" /></figure>

				<div class="card-body">
					<h2 class="card-title">کنسرت عمه فاحشه ات</h2>
					<p>Your عمه جنده's concert begins at 1402/11/22 20:30 in Qazvin</p>

					<div class="card-actions justify-end pt-2">
						<button class="btn btn-accent">Book a Ticket</button>
					</div>
				</div>
			</div>
		{/each}
	</div>
</main>

<footer class="border-base-300 mt-10 flex h-11 items-center justify-center border-t">Copyright Ticketer, no rights reserved at all.</footer>
