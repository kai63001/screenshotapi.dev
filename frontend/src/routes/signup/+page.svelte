<script lang="ts">
	// @ts-nocheck
	import bg from '$lib/assets/image/snapcaptureLogo.png?w=50&h=50&format=webp&quality=100';
	import { pb, currentUser } from '$lib/pocketbase';
	import InputField from '$lib/components/InputField.svelte';
	import Icon from '@iconify/svelte';
	import toast, { Toaster } from 'svelte-french-toast';
	import { goto } from '$app/navigation';
    import {onMount} from 'svelte';

    onMount(async () => {
        const isAuth = $currentUser;
        if(isAuth){
            goto('/dashboard');
        }
    })

	let loading = false;

	const register = async (e: any) => {
		const email = e.target.email.value;
		const password = e.target.password.value;
		const confirmPassword = e.target.confirmPassword.value;

		if (password !== confirmPassword) {
			toast.error('Passwords do not match', {
				color: 'red',
				duration: 3000,
				position: 'top-right'
			});
			return;
		}
		loading = true;
		toast.promise(
			saveUser({
				email,
				password,
				passwordConfirm: confirmPassword
			}),
			{
				loading: 'Registering...',
				success: 'Registered successfully!',
				error: 'Registration failed!'
			},
			{
				duration: 3000,
				position: 'top-right'
			}
		);
	};
	const saveUser = async (user) => {
		return new Promise((resolve, reject) => {
			setTimeout(() => {
				pb.collection('users')
					.create(user)
					.then(() => {
						resolve(true);
						loading = false;
						goto('/login');
					})
					.catch((err) => {
						Object.keys(err.data.data).forEach((key) => {
							toast.error(err.data.data[key].message, {
								color: 'red',
								duration: 3000,
								position: 'top-right'
							});
						});
						loading = false;
						reject(new Error('Could not save'));
					});
			}, 0);
		});
	};
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-white">
    <a href="/" class="absolute top-0 left-0 m-4 text-gray-800 text-sm font-bold">Home</a>
	<h1 class="text-4xl font-bold text-gray-800 mb-6">
		<img src={bg} alt="snapcapture logo" width="50px" height="50px" class="w-10 h-10 inline-block mr-2" />
		Register
	</h1>
	<form
		class="w-full md:w-[420px] flex-col space-y-6 rounded-lg p-6"
		on:submit|preventDefault={register}
	>
		<InputField label="Email" name="email" type="email" id="email" required autofocus />

		<InputField label="Password" name="password" type="password" id="password" required />

		<InputField
			label="Confirm Password"
			name="confirmPassword"
			type="password"
			id="confirmPassword"
			required
		/>

		<p class="block text-xs relative -top-5 h-0">
			Already have an account? <a href="/login" class="text-secondary hover:underline">Log In</a>.
		</p>

		<button
			type="submit"
			class="w-full bg-black text-white rounded py-3.5 text-sm group flex items-center justify-center"
			class:opacity-50={loading}
			disabled={loading}
		>
			Register
			<Icon
				icon="akar-icons:arrow-right"
				class="w-5 h-5 ml-2 transform duration-300 group-hover:translate-x-1"
			/>
		</button>
		<!-- term and policy -->
		<p class="text-xs text-gray-500 relative -top-4">
			By registering, you agree to our <a href="/terms" class="text-secondary hover:underline"
				>Terms of Service</a
			>
			and <a href="/privacy" class="text-secondary hover:underline">Privacy Policy</a>.
		</p>
	</form>
</div>
<Toaster />
