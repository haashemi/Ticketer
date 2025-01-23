<script lang="ts">
	import { browser } from '$app/environment';
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
	import '../app.css';
	interface Props {
		children?: import('svelte').Snippet;
	}

	let { children }: Props = $props();

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				enabled: browser,
				refetchOnMount: false,
				refetchOnWindowFocus: false,
				retry: 0,
			},
		},
	});
</script>

<svelte:head>
	<title>Ticketer</title>
	<meta name="description" content="A simple movie ticket e-commerce website" />
</svelte:head>

<QueryClientProvider client={queryClient}>
	<div class="min-h-screen">
		{@render children?.()}
	</div>
</QueryClientProvider>
