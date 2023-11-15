<script>
	import { pb } from '$lib/pocketbase';
	import { onMount } from 'svelte';
	//@ts-ignore
	import successPage from '$lib/assets/image/subscription/search.png?format=webp&quality=100&w=300&h=300';

	onMount(async () => {
		await pb.collection('users').authRefresh();
	});
</script>

<div class="bg-[#f7f8fc] h-screen flex m-auto">
	<div class="bg-[#edf0fa] m-auto p-3 rounded-md w-[40%]">
		<div class="bg-white p-5 rounded-md w-full flex justify-center">
			<div class="text-center py-20">
				<img src={successPage} width="300px" alt="cancel" class="m-auto mb-5" />
				<h1 class="font-semibold text-2xl py-2">Subscription Cancellation</h1>
				<p class="max-w-md">
					We're sorry to see you go. Your subscription has been successfully canceled. If you have a moment, please let us know why you're leaving. Your feedback is valuable to us or Looking for different options? Consider pausing your subscription or switching to a different plan.
				</p>
				<div class="flex space-x-2 my-2 justify-center">
					<a href="/feedback" class="border text-[#161d2b] px-10 py-3 rounded text-sm">
						Feedback Form
					</a>
					<a href="/subscription" class="bg-[#161d2b] text-white px-10 py-3 rounded text-sm">
						Change Plan
					</a>
				</div>
			</div>
		</div>
	</div>
</div>
