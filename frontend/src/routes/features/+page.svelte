<script>
	import Navbar from '$lib/components/index/Navbar.svelte';
	import Footer from '$lib/components/index/Footer.svelte';
	import Seo from '$lib/components/Seo.svelte';

	const featureLists = [
		{
			name: 'Format Images',
			description:
				'Convert your screenshots to various formats including PNG, JPEG, or PDF, ensuring compatibility and ease of use across different platforms.'
		},
		{
			name: 'Custom Set',
			description:
				'Tailor your screenshots to your needs using our API, allowing for personalized settings and adjustments.'
		},
		{
			name: 'Custom Header',
			description:
				'Modify the header of your screenshots with custom elements and styles through our API for a unique and branded look.'
		},
		{
			name: 'Custom CSS',
			description:
				'Enhance your screenshots with custom CSS styling, giving you the freedom to create visually appealing and distinct images.'
		},
		{
			name: 'Custom Javascript',
			description:
				'Incorporate custom JavaScript functionalities into your screenshots, enabling dynamic elements and interactivity.'
		},
		{
			name: 'Custom Quality',
			description:
				'Adjust the quality of your screenshots to your preference, balancing between clarity and file size, using our API.'
		},
		{
			name: 'Custom Width and Height',
			description:
				'Specify the exact width and height for your screenshots, ensuring they fit perfectly in your intended layout or design.'
		},
		{
			name: 'Custom Delay',
			description:
				'Set a custom delay for capturing screenshots, allowing web pages to fully load or animations to complete, ensuring an accurate representation.'
		},
		{
			name: 'Custom User Agent',
			description:
				'Modify the user agent in your screenshot requests to test how websites appear on different browsers or devices.'
		},
		{
			name: 'Save to S3',
			description:
				'Directly save your customized screenshots to Amazon S3, providing a seamless integration for easy storage and access.'
		}
	];
</script>

<Seo
	title="ScreenshotAPI Features - Capture and Manage Web Screenshots"
	description="Discover the powerful features of ScreenshotAPI and enhance your web development workflow. Capture and manage web screenshots with efficiency and high-quality solutions."
	path="/features"
/>


<header>
	<Navbar />
</header>

<main class="py-8 max-w-xl md:max-w-7xl px-2 m-auto pt-16">
	<section>
		<div class="text-center mt-10">
			<h2 class="text-3xl font-extrabold text-gray-900 sm:text-4xl">Features</h2>
			<p class="mt-4 text-lg text-gray-500">Discover the powerful features of ScreenshotAPI</p>
		</div>
	</section>
    <section>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4">
            {#each featureLists as featureList}
                <div class="bg-gray-100 p-4">
                    <h3 class="text-lg font-semibold">{featureList.name}</h3>
                    <p class="mt-2 text-gray-600">{featureList.description}</p>
                </div>
            {/each}
        </div>
    </section>
</main>

<br />
<br />
<br />

<Footer />
