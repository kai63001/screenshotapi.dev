<script lang="ts">
	import { onMount } from 'svelte';
	import { pb, currentUser } from '$lib/pocketbase';
	import InputField from '$lib/components/InputField.svelte';
	import Icon from '@iconify/svelte';

	const register = async (e: any) => {
		console.log(e);
		const name = e.target.name.value;
		console.log(name);
		// await pb.collection('users').create({ name, email, password, passwordConfirm: password });
	};

	onMount(() => {
		console.log('Register page mounted');
	});
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-white">
	<h1 class="text-4xl font-bold text-gray-800 mb-6">Register</h1>
	<form class="w-full md:w-[420px] flex-col space-y-6 rounded-lg p-6" on:submit|preventDefault={register}>
		<InputField label="Email" name="email" type="email" id="email" required autofocus />

		<InputField label="Password" name="password" type="password" id="password" required />

		<InputField
			label="Confirm Password"
			name="confirmPassword"
			type="password"
			id="confirmPassword"
			required
		/>


		<button
			type="submit"
			class="w-full bg-black text-white rounded py-3.5 text-sm group flex items-center justify-center"
		>
			Register
			<Icon
				icon="akar-icons:arrow-right"
				class="w-5 h-5 ml-2 transform group-hover:translate-x-1"
			/>
		</button>
	</form>
</div>
