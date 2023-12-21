<script>
	import { onMount } from 'svelte';
	import { pb, currentUser, axiosInstance } from '$lib/pocketbase';
	import { Progress } from '$lib/components/ui/progress';
	import { Switch } from '$lib/components/ui/switch';
	import { Label } from '$lib/components/ui/label';
	import { Toaster, toast } from 'svelte-french-toast';
	import * as Table from '$lib/components/ui/table';

	onMount(() => {
		getQuotaScreenshot();
		fetchHistories();
	});
	let quota = {
		screenshots_taken: 0,
		included_screenshots: 0
	};

	let defaultExtra = false;
	let disableExtra = false;
	$: {
		if (defaultExtra !== disableExtra) {
			changeDisableExtra();
		}
	}

	const changeDisableExtra = async () => {
		axiosInstance
			.post('/update_disable_extra', {
				status: disableExtra
			})
			.then((res) => {
				defaultExtra = disableExtra;
				toast.success('Success', {
					duration: 2000,
					position: 'top-right'
				});
			})
			.catch((err) => {
				toast.error('Error', {
					duration: 2000,
					position: 'top-right'
				});
			});
	};

	const getQuotaScreenshot = async () => {
		const userId = $currentUser?.id;
		const quotaCollection = await pb
			.collection('screenshot_usage')
			.getFirstListItem(`user_id = '${userId}'`, {
				expand: 'user_id,user_id.subscription_plan',
				fields:
					'screenshots_taken,expand.user_id.expand.subscription_plan.included_screenshots,disable_extra'
			});

		defaultExtra = quotaCollection.disable_extra;
		disableExtra = quotaCollection.disable_extra;
		quota = {
			screenshots_taken: quotaCollection.screenshots_taken,
			included_screenshots:
				quotaCollection?.expand?.user_id?.expand?.subscription_plan?.included_screenshots
		};
	};

	// ------------------------------ HISTORY ------------------------------
	let histories = [];
	let hasNextPage = false;
	let hasPrevPage = false;
	let totalDocs = 0;
	let page = 1;

	let loading = false;

	async function fetchHistories() {
		loading = true;
		const res = await axiosInstance.get(`/history?page=${page}`);
		const data = res.data;

		histories = [...histories, ...data.data];
		totalDocs = data.totalCount;

		page++;
		loading = false;
	}
	// ------------------------------ HISTORY ------------------------------
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Welcome to Your Dashboard!</h1>
		<p class="text-mute mt-2">
			Quickly access key metrics, recent activities, and manage your settings. All your essential
			tools and insights, streamlined in one place.
		</p>
	</div>
	<div class="flex space-x-3 w-full">
		<div class="bg-white p-5 rounded-md w-full">
			<h2 class="font-bold text-2xl">Your Monthly Quota</h2>
			<p class="text-mute mt-2">
				You have taken <span class="font-bold">{quota.screenshots_taken}</span> screenshots this
				month. Your plan includes
				<span class="font-bold">{quota.included_screenshots}</span> screenshots.
				<a href="/subscription" class="text-red-600 hover:text-red-800 hover:underline"
					>Consider upgrading</a
				> for more!
			</p>
			<div class="my-1">
				<Progress value={quota.screenshots_taken} max={quota.included_screenshots} />
			</div>
		</div>
		<!-- * SWITCH DISABLE OR ENABLE EXTRA SCREENSHOT TAKEND -->
		<div class="bg-white p-5 rounded-md w-4/12">
			<h2 class="font-bold text-2xl">Activate Extra Usage</h2>
			<div class="flex items-center space-x-2 mt-4">
				<Switch id="extra" bind:checked={disableExtra} on:change={changeDisableExtra} />
				<Label for="extra">
					<span class="font-bold">Disable</span> extra screenshot taken
				</Label>
			</div>
		</div>
	</div>
	<div class="bg-white p-5 rounded-md overflow-x-auto">
		<Table.Root >
			<Table.Caption>
				{#if histories.length > 0}
					<a href="/history" class="text-red-600 hover:text-red-800 hover:underline"
						>Showing {histories.length} of {totalDocs} results</a
					>
				{:else if loading}
					Loading...
				{:else}
					No results found
				{/if}
			</Table.Caption>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">Access Key</Table.Head>
					<Table.Head>URL</Table.Head>
					<Table.Head>Created</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each histories as history, i (i)}
					<Table.Row>
						<Table.Cell class="font-medium break-words">{history.access_key}</Table.Cell>
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
<Toaster />
