<script>
	import { pb } from '$lib/pocketbase';
	import { onMount } from 'svelte';
	//@ts-ignore
	import successPage from '$lib/assets/image/subscription/love-message.png?format=webp&quality=100&w=300&h=300';

	onMount(async () => {
		await pb.collection('users').authRefresh();
	});
</script>

<div class="bg-[#f7f8fc] h-screen flex m-auto">
	<div class="bg-[#edf0fa] m-auto p-3 rounded-md w-[40%]">
		<div class="bg-white p-5 rounded-md w-full flex justify-center">
			<div class="text-center py-20">
				<img src={successPage} width="300px" alt="success" class="m-auto mb-5" />
				<h1 class="font-semibold text-2xl py-2">Thank you for choosing ScreenshotAPI!</h1>
				<p class="max-w-md">
					Your subscription is now active. You can now use all the features of the app.
				</p>
				<div class="flex mt-1 justify-center">
					<a href="/subscription" class="bg-[#161d2b] text-white px-10 py-3 rounded mt-5 text-sm">
						Back to screenshotAPI.dev
					</a>
				</div>
			</div>
		</div>
	</div>
</div>
