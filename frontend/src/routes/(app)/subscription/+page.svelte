<script>
	import { pb, currentUser } from '$lib/pocketbase';
	import axios from 'axios';

	const instance = axios.create({
		baseURL: import.meta.env.VITE_API_KEY,
		headers: {
			Authorization: 'Bearer ' + pb.authStore.token
		}
	});

	const subscription = async () => {
		const data = await instance.post(`/subscription`, {});
		console.log(data);
	};
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Manage Your Subscription</h1>
		<p class="text-mute mt-2">
			View your current plan and explore various upgrade options. Manage your billing information
			and customize your subscription to fit your needs.
		</p>
	</div>
	<div class="bg-white p-5 rounded-md">
		<button class="p-5 bg-red-500" on:click={subscription}> test sub </button>
	</div>
	<script async src="https://js.stripe.com/v3/pricing-table.js">
	</script>
    {#if $currentUser}
	<stripe-pricing-table
		pricing-table-id="prctbl_1OC7MwH2Tv3zxv6JdzPUO7O8"
		publishable-key="pk_test_51OC3rdH2Tv3zxv6JEKhyAi3DDtDS13YdMN5Wjgp1ZFoPso3zQsbUOXBuUxnAkHg1yCIualRY7TimSWnZai3Zk0Ll004qO1noLL"
        client-reference-id="{$currentUser.id}"
        customer-email="{$currentUser.email}"
	/>
    {/if}
</div>
