<script lang="ts">
	// @ts-nocheck
	import bg from '$lib/assets/image/snapcaptureLogo.png?w=50&h=50&format=webp&quality=100';
	import { pb, currentUser } from '$lib/pocketbase';
	import InputField from '$lib/components/InputField.svelte';
	import Icon from '@iconify/svelte';
	import toast, { Toaster } from 'svelte-french-toast';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	onMount(async () => {
		const isAuth = $currentUser;
		if (isAuth) {
			goto('/playground');
		}
	});

	let loading = false;

	const login = async (e: any) => {
		const email = e.target.email.value;
		const password = e.target.password.value;

		loading = true;
		toast.promise(
			loginUser({
				email,
				password
			}),
			{
				loading: 'Logging in...',
				success: 'Login successfully!',
				error: 'Login failed!'
			},
			{
				duration: 3000,
				position: 'top-right'
			}
		);
	};
	const loginUser = async (user) => {
		return await pb
			.collection('users')
			.authWithPassword(user.email, user.password)
			.then(() => {
				loading = false;
				setTimeout(() => {
					goto('/playground');
				}, 1500);
			})
			.catch((e) => {
				loading = false;
				throw new Error(e);
			});
	};
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-white">
	<a href="/" class="absolute top-0 left-0 m-4 text-gray-800 text-sm font-bold">Home</a>
	<h1 class="text-4xl font-bold text-gray-800 mb-6">
		<img src={bg} width="50px" height="50px" alt="ScreenshotAPI logo" class="w-10 h-10 inline-block mr-2" />
		Sign In
	</h1>
	<form
		class="w-full md:w-[420px] flex-col space-y-6 rounded-lg p-6"
		on:submit|preventDefault={login}
	>
		<InputField label="Email" name="email" type="email" id="email" required autofocus />

		<InputField label="Password" name="password" type="password" id="password" required />

		<p class="block text-xs relative -top-5 h-0">
			Don't have an account? <a href="/signup" class="text-primary hover:underline">Sign Up</a>.
		</p>

		<button
			type="submit"
			class="w-full bg-black text-white rounded py-3.5 text-sm group flex items-center justify-center"
			class:opacity-50={loading}
			disabled={loading}
		>
			Login
			<Icon
				icon="akar-icons:arrow-right"
				class="w-5 h-5 ml-2 transform duration-300 group-hover:translate-x-1"
			/>
		</button>
		<!-- term and policy -->
		<p class="text-xs text-gray-500 relative -top-4">
			By Sign In, you agree to our <a href="/terms" class="text-primary hover:underline"
				>Terms of Service</a
			>
			and <a href="/privacy" class="text-primary hover:underline">Privacy Policy</a>.
		</p>
	</form>
</div>
<Toaster />
