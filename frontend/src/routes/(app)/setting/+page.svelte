<script lang="ts">
	import InputField from '$lib/components/InputField.svelte';
	import { axiosInstance, currentUser, pb } from '$lib/pocketbase';
	import { onMount } from 'svelte';
	import Icon from '@iconify/svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import { goto } from '$app/navigation';
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
	let deleteConfirm = '';

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

	const logout = async () => {
		await localStorage.removeItem('access_key');
		await pb.authStore.clear();
		//refresh page

		window.location.reload();
	};

	// Delete Account Function
	const deleteAccount = async () => {
		await axiosInstance
			.delete('/delete_account')
			.then((res) => {
				pb.collection('users').authRefresh();
				logout();
			})
			.catch((err) => {
				console.log(err);
			});
	};
</script>

<Seo
	title="Settings - ScreenshotAPI.dev"
	description="Get a comprehensive overview of your projects and manage your screenshot capturing effortlessly with ScreenshotAPI.dev's powerful dashboard. Explore our documentation and enhance your web development workflow."
	path="/setting"
/>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Access Settings</h1>
		<p class="text-mute mt-2">
			Manage and reset access keys for enhanced security. Ensure secure and appropriate access for
			users.
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

	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Delete Account</h1>
		<p class="text-mute mt-2">
			Deleting your account is irreversible. Please enter your password to confirm.
		</p>
		<Dialog.Root>
			<Dialog.Trigger>
				<button class="bg-primary text-white px-7 py-2 rounded-md mt-4 flex items-center space-x-2">
					<Icon
						class="text-white mr-1"
						icon="material-symbols:delete-outline"
						width="20px"
						height="20px"
					/>
					Delete Account
				</button>
			</Dialog.Trigger>
			<Dialog.Content class="sm:max-w-[425px]">
				<Dialog.Header>
					<Dialog.Title>Confirm Delete Account</Dialog.Title>
					<Dialog.Description>
						Are you sure you want to delete your account? This action is irreversible.
					</Dialog.Description>
				</Dialog.Header>
				<div>
					<InputField
						type="text"
						label={'Confirm'}
						placeholder="Type 'delete' to delete"
						bind:value={deleteConfirm}
					/>
				</div>
				<Dialog.Footer>
					<button
						class="bg-primary text-white px-7 py-2 rounded-md mt-4"
						disabled={deleteConfirm !== 'delete'}
						class:bg-red-200={deleteConfirm !== 'delete'}
						on:click={deleteAccount}
					>
						Confirm Delete
					</button>
				</Dialog.Footer>
			</Dialog.Content>
		</Dialog.Root>
	</div>
</div>
