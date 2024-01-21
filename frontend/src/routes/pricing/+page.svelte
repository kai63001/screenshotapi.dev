<script>
	import Navbar from '$lib/components/index/Navbar.svelte';
	import PricingBox from '$lib/components/pricing/PricingBox.svelte';
	import Footer from '$lib/components/index/Footer.svelte';
	import PricingFaq from '$lib/components/pricing/PricingFaq.svelte';

	import Seo from '$lib/components/Seo.svelte';
</script>

<Seo
	title="Affordable ScreenshotAPI Plans - Pricing for Developers & Businesses"
	description="Explore ScreenshotAPI's affordable pricing for developers & businesses. Efficient, high-quality solutions for web screenshots, content management, and archiving."
	path="/pricing"
/>
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
