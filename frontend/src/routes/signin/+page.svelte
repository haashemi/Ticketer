<script lang="ts">
	import { goto } from '$app/navigation';
	import { ValidationMessage, reporter } from '@felte/reporter-svelte';
	import { validator } from '@felte/validator-zod';
	import { createForm } from 'felte';
	import ky, { HTTPError } from 'ky';
	import * as z from 'zod';

	let loading: boolean = false;
	let error: string | null = null;

	const schema = z.object({
		email: z.string().email(),
		password: z.string().min(6).max(72),
	});

	const { form } = createForm({
		extend: [validator({ schema }), reporter],
		onSubmit: async (v: z.infer<typeof schema>) => {
			loading = true;
			ky.post('/api/auth/sign-in', { json: v })
				.then(() => goto('/'))
				.catch((e) => (e as HTTPError).response.json().then((v) => (error = v.message)))
				.finally(() => (loading = false));
		},
	});
</script>

<div class="absolute left-1/2 top-1/2 -z-10 h-96 w-96 -translate-x-1/2 -translate-y-1/2 rounded-full bg-blue-900 blur-3xl" />

<main class="flex min-h-screen items-center justify-center bg-black/80 backdrop-blur-xl">
	<div class="flex w-96 flex-col items-center justify-center gap-6 rounded-xl border border-slate-700 p-10">
		<h1 class="text-2xl font-semibold">Welcome Back!</h1>

		{#if error}
			<div class="w-full text-wrap rounded-xl border border-red-500 bg-red-500/20 p-5">{error}</div>
		{/if}

		<form use:form class="w-full">
			<fieldset disabled={loading} class="flex flex-col gap-3">
				<label class="form-control w-full max-w-xs">
					<div class="label">
						<span class="label-text">Email</span>
					</div>
					<input name="email" id="email" required type="email" placeholder="username@gmail.com" class="input input-bordered w-full max-w-xs" />

					<ValidationMessage for="email" let:messages>
						<ul class="label" aria-live="polite">
							{#each messages ?? [] as message}
								<li class="label-text text-red-500">{message}</li>
							{/each}
						</ul>
					</ValidationMessage>
				</label>

				<label class="form-control w-full max-w-xs">
					<div class="label">
						<span class="label-text">Password</span>
					</div>
					<input name="password" id="password" required type="password" placeholder="********" class="input input-bordered w-full max-w-xs" />

					<ValidationMessage for="password" let:messages>
						<ul class="label" aria-live="polite">
							{#each messages ?? [] as message}
								<li class="label-text text-red-500">{message}</li>
							{/each}
						</ul>
					</ValidationMessage>
				</label>

				<div class="flex justify-end">
					<button class="btn btn-primary">Sign In</button>
				</div>
			</fieldset>
		</form>

		<p>Don't have an account? <a class="text-primary" href="/signup">Sign Up!</a></p>
	</div>
</main>
