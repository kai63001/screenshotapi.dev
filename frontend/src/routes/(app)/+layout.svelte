<script>
	//@ts-ignore
	import logo from '$lib/assets/image/snapcaptureLogo.png?w=30&h=30&format=webp&quality=100';
	//@ts-ignore
	import avatar from '$lib/assets/avatar/man.png?w=50&h=50&format=webp&quality=100';
	import { page } from '$app/stores';

	import { onMount } from 'svelte';
	import { currentUser, pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';
	import Icon from '@iconify/svelte';

	let navbarList = [
		{
			name: 'Dashboard',
			icon: 'carbon:dashboard',
			path: '/dashboard',
			active: false
		},
		{
			name: 'Playground',
			icon: 'material-symbols:joystick-outline',
			path: '/playground',
			active: false
		},
		{
			name: 'Custom SET',
			icon: 'icon-park-outline:full-selection',
			path: '/custom-set',
			active: false
		},
		// {
		// 	name: 'Access',
		// 	icon: 'ph:key-bold',
		// 	path: '/access',
		// 	active: false
		// },
		{
			name: 'History',
			icon: 'majesticons:image-multiple-line',
			path: '/history',
			active: false
		},
		{
			name: 'Subscription',
			icon: 'ion:card-outline',
			path: '/subscription',
			active: false
		},
		{
			name: 'Payments',
			icon: 'gg:list',
			path: '/payments',
			active: false
		},
		{
			name: 'Settings',
			icon: 'uil:setting',
			path: '/setting',
			active: false
		}
	];

	$: currentPath = $page.url.pathname;
	onMount(async () => {
		const isAuth = $currentUser;
		if (!isAuth) {
			goto('/login');
		}
	});

	const logout = async () => {
		await localStorage.removeItem('access_key');
		await pb.authStore.clear();
		goto('/login');
	};
</script>

<div class="flex">
	<nav class="w-1/5 flex flex-col justify-between fixed h-screen">
		<div>
			<a href="/" class="p-5 flex items-center space-x-2">
				<img src={logo} alt="logo" width="30px" height="30px" loading="eager" />
				<p class="ml-2 heading text-xl">ScreenshotAPI</p>
			</a>
			<div class="p-5">
				<div class="px-5 py-3 bg-[#faf9fb] rounded">
					<div class="flex items-center space-x-2">
						<div class="hidden xl:block w-10 h-10 rounded-full bg-[#e5e7eb]">
							<img src={avatar} alt="avatar" width="40px" height="40px" loading="eager" />
						</div>
						<div class="flex flex-col">
							<p class="text-sm font-semibold capitalize">
								{($currentUser && $currentUser.username) || ''}
							</p>
							<p class="text-xs text-gray-500">
								{($currentUser && $currentUser.email) || ''}
							</p>
						</div>
					</div>
				</div>
			</div>
			<ul class="pr-5 flex flex-col space-y-1">
				{#each navbarList as nav}
					<li>
						<a
							class:bg-block={nav.path == currentPath}
							href={nav.path}
							class="flex items-center space-x-2 cursor-pointer py-4 rounded-r-md hover:bg-red-100"
						>
							{#if nav.path == currentPath}
								<div class="w-1 h-7 bg-primary rounded-r-2xl" />
							{/if}
							<div class="pl-5 flex items-center">
								<Icon
									class={`${nav.path == currentPath ? 'text-red-600' : 'text-gray-600'}`}
									icon={nav.icon}
									width="20px"
									height="20px"
								/>
								<p class="ml-2 text-sm font-semibold">{nav.name}</p>
							</div>
						</a>
					</li>
				{/each}
			</ul>
		</div>
		<div>
			<ul class="pr-5 pb-5">
				<button
					on:click={logout}
					class="flex w-full items-center space-x-2 cursor-pointer py-4 rounded-r-md hover:bg-red-100"
				>
					<div class="pl-5 flex items-center">
						<Icon icon={'basil:logout-outline'} width="20px" height="20px" />
						<p class="ml-2 text-sm font-semibold">{'Logout'}</p>
					</div>
				</button>
			</ul>
		</div>
	</nav>
	<div class="w-1/5"></div>
	<div class="w-4/5 bg-[#F5F4F6] min-h-screen h-full p-5">
		<slot />
	</div>
</div>

<style scoped>
	.bg-block {
		background-color: #faf9fb;
	}
</style>
