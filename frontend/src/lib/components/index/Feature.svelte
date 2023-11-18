<script>
	import Icon from '@iconify/svelte';

	const feature = [
		{
			icon: 'solar:camera-bold-duotone',
			title: 'Snap the Web Instantly',
			description:
				"Harness the power of our optimized capture technology that lets you take full-page screenshots in a flash. Whether it's dynamic content or high-resolution images, our API delivers pixel-perfect snapshots with zero lag. Ideal for real-time applications, monitoring services, or content archiving, our solution ensures you get the visual data you need the moment you need it."
		},
		{
			icon: 'solar:plug-circle-bold-duotone',
			title: 'Easy Integration',
			description:
				'Our API is designed for seamless integration, providing a straightforward, well-documented process that fits effortlessly into your current tech stack. With language-agnostic endpoints and comprehensive guides, you can embed our screenshot capabilities into your application with minimal overhead. From startups to large enterprises, our API scales to meet your needs without a steep learning curve.'
		},
        {
			icon: 'solar:paint-roller-bold-duotone',
			title: 'Tailored to Your Needs',
			description:
				'Flexibility is at the core of our screenshot API. Customize everything from resolution, aspect ratio, to automated capture timings to fit your unique use case. Our API adapts to your requirements, offering options like delayed snapshots for sites with lazy-loaded content, and the ability to capture behind-login screens. Whatever your challenge, our API is engineered to provide a solution'
		}
	];
</script>

<section class="py-4 px-10">
	<h2 class="text-5xl font-semibold text-center heading">Instant Capture</h2>
	<p class="text-center font-semibold text-xl">Your API for Lightning-Fast Screen Snaps</p>
	<div class="grid grid-cols-3 gap-4 mt-5">
		{#each feature as item}
			<div class="shadow-md rounded-md p-7">
				<div class="flex space-x-2 items-center">
					<Icon class="text-5xl text-primary" icon={item.icon} />
				</div>
				<p class="text-xl font-semibold">{item.title}</p>
				<p class="text-muted-foreground mt-1">
					{item.description}
				</p>
			</div>
		{/each}
	</div>
</section>
