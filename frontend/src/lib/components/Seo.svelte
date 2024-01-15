<script>
	import SvelteSeo from 'svelte-seo';

	export let title = 'ScreenshotAPI: The Best Screenshot API for Developers & Businesses';
	export let description =
		'Introducing ScreenshotAPI - the ultimate tool for developers and businesses to easily capture high-quality webpage screenshots. Perfect for efficient content management and web archiving. Start with our user-friendly, customizable API today.';

	export let path = '';
</script>

<SvelteSeo
	{title}
	{description}
	canonical={`https://screenshotapi.dev${path}`}
	openGraph={{
		title: title,
		description: description,
		url: `https://screenshotapi.dev${path}`,
		type: 'website',
		images: [
			{
				url: 'https://www.primal-movement.com/images/jumping.jpg'
			}
		],
		site_name: 'ScreenshotAPI.DEV'
	}}
	twitter={{
		card: 'summary_large_image',
		site: '@screenshotapi.dev',
		title: title,
		description: description,
		image: 'https://www.primal-movement.com/images/squatting.jpg'
	}}
/>
