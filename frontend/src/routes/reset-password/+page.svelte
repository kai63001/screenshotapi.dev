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

    const sendResetPassword = async (e: any) => {
        const email = e.target.email.value;
        loading = true;
        toast.promise(
            resetPassword({
                email
            }),
            {
                loading: 'Sending reset password email...',
                success: 'Reset password email sent!',
                error: 'Reset password email failed!'
            },
            {
                duration: 3000,
                position: 'top-right'
            }
        );
    };

    const resetPassword = async (user) => {
        return await pb
            .collection('users')
            .requestPasswordReset(user.email)
            .then(() => {
                loading = false;
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
        Forgot Password
	</h1>
	<form
		class="w-full md:w-[420px] flex-col space-y-6 rounded-lg p-6"
		on:submit|preventDefault={sendResetPassword}
	>
		<InputField label="Email" name="email" type="email" id="email" required autofocus />

		<p class="block text-xs relative -top-5 h-0">
			 <a href="/login" class="text-primary hover:underline">Sign In</a>
             ·
			 <a href="/signup" class="text-primary hover:underline">Sign Up</a>
		</p>

		<button
			type="submit"
			class="w-full bg-black text-white rounded py-3.5 text-sm group flex items-center justify-center"
			class:opacity-50={loading}
			disabled={loading}
		>
			Reset Password
			<Icon
				icon="streamline:send-email-solid"
				class="w-5 h-5 ml-2 transform duration-300 group-hover:translate-x-1"
			/>
		</button>
	</form>
</div>
<Toaster />
