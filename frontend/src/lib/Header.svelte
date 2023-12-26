<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query';
	import ky from 'ky';
	import { z } from 'zod';

	const respSchema = z.object({
		name: z.string(),
		email: z.string().email(),
		isAdmin: z.boolean(),
	});

	const profile = createQuery({
		queryKey: ['profile'],
		queryFn: async () => {
			const resp = await ky.get('/api/profile').json();
			return respSchema.parse(resp);
		},
	});
</script>

<header class="from-base-100 sticky top-0 z-20 w-full bg-gradient-to-b from-30% to-transparent">
	<header class="container mx-auto flex h-16 items-center justify-between px-5">
		<a href="/" class="font-silk text-2xl font-bold text-white">Ticketer</a>

		{#if $profile.isSuccess}
			<div class="flex gap-2">
				<a href="/profile" class="btn btn-ghost min-h-10 h-10 font-medium">Welcome {$profile.data.name}</a>
			</div>
		{:else}
			<div class="flex gap-2">
				<a href="/signup" class="btn btn-ghost min-h-10 h-10 font-medium">Sign Up</a>
				<a href="/signin" class="btn btn-primary min-h-10 h-10 font-medium">Sign In</a>
			</div>
		{/if}
	</header>
</header>
