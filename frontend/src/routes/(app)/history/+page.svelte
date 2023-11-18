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
	let page = 1;

	let loading = false;

	async function fetchHistories() {
		loading = true;
		const res = await instance.get(`/history?page=${page}`);
		const data = res.data;

		histories = [...histories, ...data.data];
		hasNextPage = data.hasNextPage;
		hasPrevPage = data.hasPrevPage;

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
			<Table.Caption>A list of your recent API history.</Table.Caption>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">Access Key</Table.Head>
					<Table.Head>Created</Table.Head>
					<Table.Head>Full URL</Table.Head>
					<Table.Head>URL</Table.Head>
					<Table.Head class="text-right">User ID</Table.Head>
					<Table.Head class="text-right">ID</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each histories as history, i (i)}
					<Table.Row>
						<Table.Cell class="font-medium">{history.access_key}</Table.Cell>
						<Table.Cell>{history.created}</Table.Cell>
						<Table.Cell class="whitespace-pre-line sm:whitespace-normal group">
							<div class="w-20 sm:w-auto">
								<div
									class="text-ellipsis overflow-hidden whitespace-nowrap w-20 group-hover:hidden"
								>
									{history.fullUrl}
								</div>
								<div class="hidden group-hover:flex">
									{history.fullUrl}
								</div>
							</div>
						</Table.Cell>
						<Table.Cell>{history.url}</Table.Cell>
						<Table.Cell class="text-right">{history.user_id}</Table.Cell>
						<Table.Cell class="text-right">{history._id}</Table.Cell>
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
