<script>
	import { pb } from "$lib/pocketbase";
	import axios from "axios";
	import { onMount } from "svelte";

	const instance = axios.create({
		baseURL: import.meta.env.VITE_API_KEY,
		headers: {
			Authorization: 'Bearer ' + pb.authStore.token
		}
	});

	onMount(async () => {
		instance.get('/history').then((res) => {
			console.log(res.data)
		});
	});
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Screenshot API History</h1>
		<p class="text-mute mt-2">
			Review your recent API activity, including timestamps, endpoints, and outcomes. Easily search
			and filter through your history for detailed insights.
		</p>
	</div>
</div>
