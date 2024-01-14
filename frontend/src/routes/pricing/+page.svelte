<script>
	import Navbar from '$lib/components/index/Navbar.svelte';
	import PricingBox from '$lib/components/pricing/PricingBox.svelte';
	import Footer from '$lib/components/index/Footer.svelte';
	import PricingFaq from '$lib/components/pricing/PricingFaq.svelte';
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
		<PricingBox />
	</section>
	<section>
		<PricingFaq />
	</section>
</main>

<br />
<br />
<br />

<Footer />
