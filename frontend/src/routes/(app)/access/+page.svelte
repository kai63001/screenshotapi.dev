<script lang="ts">
	import InputField from '$lib/components/InputField.svelte';
	import { axiosInstance, currentUser, pb } from '$lib/pocketbase';
	import { onMount } from 'svelte';
	import Icon from '@iconify/svelte';
	import Seo from '$lib/components/Seo.svelte';


	onMount(async () => {
		pb.collection('access_keys')
			.getFirstListItem(`user_id = '${$currentUser.id}'`)
			.then((res) => {
				accessKey = res.access_key;
			});
	});

	let accessKey = '';
	let loading = false;

	const resetAccessKey = () => {
		loading = true;
		axiosInstance
			.patch('/access_key', {
				user_id: $currentUser.id
			})
			.then((res) => {
				accessKey = res.data.access_key;
				pb.collection('users').authRefresh();
			})
			.finally(() => {
				loading = false;
			});
	};
</script>

<Seo
	title="Access - ScreenshotAPI.dev"
	description="Get a comprehensive overview of your projects and manage your screenshot capturing effortlessly with ScreenshotAPI.dev's powerful dashboard. Explore our documentation and enhance your web development workflow."
	path="/access"
/>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Access Management</h1>
		<p class="text-mute mt-2">
			Quickly manage and reset access keys for enhanced security. Ensure secure and appropriate
			access for users.
		</p>
	</div>
	<div class="bg-white p-5 rounded-md">
		<InputField
			type="text"
			label={'Access Key'}
			placeholder="Access Key"
			bind:value={accessKey}
			disabled
		/>
		<div class="mt-2">
			<button
				class="bg-primary text-white px-7 py-2 rounded-md flex space-x-2 items-center"
				disabled={loading}
				class:opacity-50={loading}
				on:click={resetAccessKey}
			>
				<Icon
					class="text-white mr-1"
					icon="fluent:key-reset-20-filled"
					width="20px"
					height="20px"
				/>
				Reset Access Key</button
			>
		</div>
	</div>
</div>
