<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query';
	import ky from 'ky';
	import { z } from 'zod';
	//@ts-ignore
	import Icon from 'svelte-icons-pack/Icon.svelte';
	import FiLogOut from 'svelte-icons-pack/fi/FiLogOut';

	let loading = false;

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
		retry: 0,
		refetchOnMount: true,
	});

	function logout() {
		loading = true;
		ky.post('/api/auth/sign-out')
			.then(() => $profile.refetch())
			.finally(() => (loading = false));
	}
</script>

<header class="from-base-100 sticky top-0 z-20 w-full bg-gradient-to-b from-30% to-transparent">
	<header class="container mx-auto flex h-16 items-center justify-between px-5">
		<a href="/" class="font-silk text-2xl font-bold text-white">Ticketer</a>

		{#if $profile.isLoading}
			<div class="flex animate-pulse gap-2">
				<div class="h-10 w-40 rounded-lg bg-zinc-800" />
				<div class="h-10 w-10 rounded-lg bg-zinc-800" />
			</div>
		{:else if $profile.isSuccess}
			<div class="flex gap-2">
				<a href="/profile" class="btn btn-primary h-10 rounded-lg font-medium {loading && 'btn-disabled'}">{$profile.data.name}</a>
				<button disabled={loading} class="btn btn-outline btn-error h-10 rounded-lg font-medium" on:click={logout}>
					<Icon src={FiLogOut} size="1.25rem" title="Logout icon" /></button
				>
			</div>
		{:else if $profile.isError}
			<div class="flex gap-2">
				<a href="/signup" class="btn btn-ghost h-10 min-h-10 font-medium">Sign Up</a>
				<a href="/signin" class="btn btn-primary h-10 min-h-10 font-medium">Sign In</a>
			</div>
		{/if}
	</header>
</header>
