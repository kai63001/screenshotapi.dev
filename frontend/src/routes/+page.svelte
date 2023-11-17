<script>
	// @ts-nocheck

	import Button from '$lib/components/ButtonCustom.svelte';
	import Icon from '@iconify/svelte';
	import Navbar from '$lib/components/index/Navbar.svelte';
	import Hero from '$lib/components/index/Hero.svelte';
</script>

<svelte:head>
	<title>ScreenshotAPI: Easy Webpage Screenshot API for Developers & Businesses</title>
	<meta
		name="description"
		content="Discover ScreenshotAPI, the user-friendly screenshot API designed for seamless integration. Capture high-quality webpage screenshots with our fast, customizable tool. Ideal for developers and businesses, SnapCapture enhances your project with efficient web capture technology. Start simplifying your content management and web archiving today."
	/>
</svelte:head>
<header>
	<Navbar />
	<div
		class="max-w-xl md:max-w-7xl px-2 m-auto flex flex-col lg:flex-row space-y-3 lg:space-y-0 lg:space-x-3 justify-between pt-28"
	>
		<div class="w-full lg:w-6/12">
			<h1 class="text-5xl font-bold heading leading-tight text-gray-900">
				Capture Webpages Easily with Our Simple <span
					class="bg-gradient-to-r from-rose-400 to-red-500 bg-clip-text text-transparent"
					>Screenshot API</span
				>.
			</h1>
			<p class="text-lg text-gray-500 mt-5">
				Effortless, High-Quality Webpage Screenshots in Seconds {'–'}
				<strong>Fully Customizable to Fit Your Unique Needs.</strong>
				Tailored for Developers and Businesses.
				<a class="text-primary hover:underline" href="/docs/introduction"> Learn More. </a>
			</p>
			<a href="/playground">
				<Button class="mt-5 flex items-center space-x-2 shadow-md"
					>Get started for Free <Icon
						class="ml-2"
						icon="material-symbols:arrow-right-alt-rounded"
						width="25"
					/></Button
				>
			</a>
			<p class="text-xs mt-1 text-gray-500">No credit card required</p>
		</div>
		<div class="w-full lg:w-6/12">
			<Hero/>
		</div>
	</div>
</header>

<!-- <main class="py-8">
	<section class="py-4">
		<h2 class="text-2xl font-bold">How it works</h2>
		<p>
			Simply make a request to our API with the URL of the website you want to capture, and we'll
			return a screenshot image for you to use.
		</p>
	</section>

	<section class="py-4">
		<h2 class="text-2xl font-bold">Features</h2>
		<ul class="list-disc pl-8">
			<li>High-quality screenshots</li>
			<li>Fast and reliable API</li>
			<li>Customizable options</li>
			<li>Secure and private</li>
		</ul>
	</section>

	<section class="py-4">
		<h2 class="text-2xl font-bold">Get started</h2>
		<p>Sign up for an API key and start generating screenshots in minutes.</p>
		<a
			href="https://api.screenshotwebsite.com/signup"
			class="bg-blue-500 text-white py-2 px-4 rounded-md">Sign up</a
		>
	</section>
</main>

<footer class="bg-gray-200 py-4">
	<p>&copy; 2022 Screenshot Website API. All rights reserved.</p>
</footer> -->
