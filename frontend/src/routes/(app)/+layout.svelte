<script>
	import { onMount } from 'svelte';
	import { currentUser } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	onMount(async () => {
		const isAuth = $currentUser;
		if (!isAuth) {
			goto('/login');
		}
	});
</script>


<slot />