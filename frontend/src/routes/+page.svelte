<script>
	import Button from '$lib/components/ButtonCustom.svelte';
	import Icon from '@iconify/svelte';
	import Navbar from '$lib/components/index/Navbar.svelte';
	import Hero from '$lib/components/index/Hero.svelte';
	import EverthingYouneed from '$lib/components/index/EverthingYouneed.svelte';
	import ShowCase from '$lib/components/index/ShowCase.svelte';
	import PricingBox from '$lib/components/pricing/PricingBox.svelte';
	import Footer from '$lib/components/index/Footer.svelte';
	import LangMain from '$lib/components/index/langsupport/LangMain.svelte';
	import FaqIndex from '$lib/components/index/FaqIndex.svelte';
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
				<a class="text-red-800 underline" href="/docs/introduction"> Learn More. </a>
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
			<Hero />
		</div>
	</div>
</header>

<main class="py-8 max-w-xl md:max-w-7xl px-2 m-auto mt-10">
	<EverthingYouneed />
	<!-- <Feature /> -->
	<ShowCase />
	<LangMain />
	<PricingBox />
	<FaqIndex />
</main>

<Footer />
