<script>
	// @ts-nocheck
	import logo from '$lib/assets/image/snapcaptureLogo.png?&format=webp&quality=100&w=25&h=25';
	import Button from '../Button.svelte';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	const menuList = [
		{
			name: 'Features',
			url: '/'
		},
		{
			name: 'Use Cases',
			url: '/about'
		},
		{
			name: 'Pricing',
			url: '/pricing'
		},
		{
			name: 'Docs',
			url: '/docs'
		},
		{
			name: 'Blog',
			url: '/blog'
		}
	];
	let openToggle = false;
	onMount(() => {
		function handleClickOutside(event) {
			const navbar = document.querySelector('#navbar');
			if (navbar && !navbar.contains(event.target)) {
				openToggle = false;
			}
		}

		window.addEventListener('click', handleClickOutside);

		return () => {
			window.removeEventListener('click', handleClickOutside);
		};
	});
</script>

<nav id="navbar" class="shadow-md m-auto fixed bg-white w-full">
	<div class="flex items-center justify-between max-w-7xl px-2 m-auto">
		<a href="/" class="flex space-x-2 items-center">
			<img src={logo} alt="logo" loading="eager" />
			<p class="heading">SnapCapture</p>
		</a>
		<ul class="hidden md:flex space-x-4">
			{#each menuList as menu}
				<li><a href={menu.url}>{menu.name}</a></li>
			{/each}
		</ul>
		<ul class="hidden md:flex space-x-4 items-center">
			<li><a href="/login">Login</a></li>
			<li>
				<a href="/signup">
					<Button class="py-1.5 my-3 px-3 rounded">Get stated for Free</Button>
				</a>
			</li>
		</ul>
		<ul class="md:hidden">
			<div class="my-2">
				<button
					on:click={() => {
						openToggle = !openToggle;
					}}
					class="border rounded p-1"
				>
					<Icon icon="ci:hamburger-md" width="30" />
				</button>
			</div>
			<div
				class:h-screen={openToggle}
				class:h-0={!openToggle}
				class="absolute w-full duration-300 overflow-hidden bg-white left-0 px-3 flex flex-col space-y-1 shadow-xl"
			>
				{#each menuList as menu}
					<li class="py-2 w-full"><a href={menu.url}>{menu.name}</a></li>
				{/each}
				<div class="border-b" />
				<li class="py-2 w-full"><a href="/login">Login</a></li>
				<li class="py-2 w-full"><a href="/signup"> Get started for Free </a></li>
			</div>
		</ul>
	</div>
</nav>
