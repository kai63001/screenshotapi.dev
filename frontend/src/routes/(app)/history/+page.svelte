<script>
	import { pb } from '$lib/pocketbase';
	import axios from 'axios';
	import { onMount } from 'svelte';
	import * as Table from '$lib/components/ui/table';

	const instance = axios.create({
		baseURL: import.meta.env.VITE_API_KEY,
		headers: {
			Authorization: 'Bearer ' + pb.authStore.token
		}
	});

	let histories = [];
	let hasNextPage = false;
	let hasPrevPage = false;
	let totalDocs = 0;
	let page = 1;

	let loading = false;

	async function fetchHistories() {
		loading = true;
		const res = await instance.get(`/history?page=${page}`);
		const data = res.data;

		histories = [...histories, ...data.data];
		hasNextPage = data.hasNextPage;
		hasPrevPage = data.hasPrevPage;
		totalDocs = data.totalCount;

		page++;
		loading = false;
	}

	onMount(fetchHistories);
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Screenshot API History</h1>
		<p class="text-mute mt-2">
			Review your recent API activity, including timestamps, endpoints, and outcomes. Easily search
			and filter through your history for detailed insights.
		</p>
	</div>
	<div class="bg-white p-5 rounded-md overflow-x-auto">
		<Table.Root>
			<Table.Caption>
				{#if histories.length > 0}
					Showing {histories.length} of {totalDocs} results
				{:else if loading}
					Loading...
				{:else}
					No results found
				{/if}
			</Table.Caption>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">Access Key</Table.Head>
					<Table.Head>Full URL</Table.Head>
					<Table.Head>URL</Table.Head>
					<Table.Head>Created</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each histories as history, i (i)}
					<Table.Row>
						<Table.Cell class="font-medium">{history.access_key}</Table.Cell>
						<Table.Cell class="w-[300px] break-words group block">
							{history.fullUrl}
						</Table.Cell>
						<Table.Cell class="break-words">{history.url}</Table.Cell>
						<Table.Cell class="break-words">{history.created}</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
		{#if hasNextPage || hasPrevPage}
			<div class="flex justify-center mt-4">
				<button
					class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
					on:click={fetchHistories}
					disabled={loading}
					class:opacity-50={loading}
				>
					Load More
				</button>
			</div>
		{/if}
	</div>
</div>
