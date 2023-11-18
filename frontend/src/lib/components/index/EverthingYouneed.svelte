<script>
	import Icon from '@iconify/svelte';

	const list = [
		{
			title: 'Competitive Pricing',
			description:
				'Start with our complimentary plan, which includes 1,000 screenshots monthly. Our scalable pricing model ensures the more you use, the less you pay—with rates dropping to as low as $0.50 per 1,000 screenshots.'
		},
		{
			title: 'Exceptional Image Fidelity',
			description:
				'Our service guarantees high-definition captures, with the flexibility to set the browser resolution to meet your exact specifications.'
		},
		{
			title: 'Cross-Platform Compatibility',
			description:
				'Manipulate the browser viewport to emulate devices of all sizes—from mobile phones and tablets to desktop computers.'
		},
		{
			title: 'Dynamic Resizing',
			description:
				'Effortlessly adjust the size of your images on-the-fly, ensuring perfect thumbnails or custom dimensions while preserving the original aspect ratio.'
		},
		{
			title: 'Versatile Image Formats',
			description:
				'Flexibility is key; choose from JPG, PNG, GIF, or WebP formats for your images, instantly prepared by our robust system.'
		},
		{
			title: 'Integrated CDN Option',
			description:
				'Optimize delivery and reduce your storage footprint with our CDN option, designed for swift and reliable distribution of your screenshots.'
		},
		{
			title: 'Tailored Developer Feedback',
			description:
				'Receive data in the format that works for you, be it the direct image or a JSON object detailing image IDs, CDN URLs, and more.'
		},
		{
			title: 'Adjustable Capture Timing',
			description:
				'Gain control over capture timing with customizable delays, ensuring even the most resource-intensive pages are rendered flawlessly.'
		}
	];
</script>

<section class="py-4 flex flex-col md:flex-row md:space-x-10 justify-between">
	<div class="w-full md:w-4/12">
		<p class="font-semibold text-lg text-red-500">ALL YOU REQUIRE</p>
		<h2 class="text-4xl font-semibold heading">All-in-one platform</h2>
		<p class="mt-3">
			Unlock the full potential of our screenshot API with a free account, ready to deliver 100 snapshots monthly at no cost
		</p>
	</div>
	<div class="grid grid-cols-2 gap-4 mt-5 w-full md:w-8/12">
		{#each list as item}
			<div>
				<div class="flex items-center space-x-2">
                    <Icon class="text-3xl text-primary" icon="solar:verified-check-bold" />
					<p class="text-xl font-semibold">{item.title}</p>
				</div>
				<p class="text-muted-foreground mt-1">
					{item.description}
				</p>
			</div>
		{/each}
	</div>
</section>
