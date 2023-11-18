<script>
	import Button from '$lib/components/ButtonCustom.svelte';
	import Icon from '@iconify/svelte';
	import Navbar from '$lib/components/index/Navbar.svelte';
	import Hero from '$lib/components/index/Hero.svelte';
	import EverthingYouneed from '$lib/components/index/EverthingYouneed.svelte';
	import ShowCase from '$lib/components/index/ShowCase.svelte';
	import PricingBox from '$lib/components/pricing/PricingBox.svelte';
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
</header>

<main class="py-8 max-w-xl md:max-w-7xl px-2 m-auto pt-16">
    <section>
        <PricingBox/>
    </section>
</main>

<!-- <footer class="bg-gray-200 py-4">
	<p>&copy; 2022 Screenshot Website API. All rights reserved.</p>
</footer> -->
