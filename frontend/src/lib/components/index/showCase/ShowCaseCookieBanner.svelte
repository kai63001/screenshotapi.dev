<script>
	import { onMount } from 'svelte';

	import Typewriter from 'typewriter-effect/dist/core';
    //@ts-ignore
    import bg from '$lib/assets/image/index/cookiebot_with_banner.jpg?format=webp&quality=100&w=573&h=456';

	function displayApiResponse() {
		const targetElement = document.getElementById('showcase2-response');

		// Create a new instance of the Typewriter
		const typewriter = new Typewriter(targetElement, {
			delay: 50, // Delay between each character
			cursor: '|',
			loop: true
		});

		// Start the typewriter effect
		typewriter
			.pauseFor(1000) // Pause for 1 second before applying animations
			.deleteAll() // Delete the existing text
			.typeString('&no_cookie_banner=true') // Type the width query parameter
			.pauseFor(1500) // Pause for 0.5 seconds before applying the next animation
			.start()
			.pauseFor(2000) // Pause for 2 seconds before restarting the animation
			.deleteAll() // Delete the existing text
			.start();
	}
	onMount(() => {
		displayApiResponse();
	});
</script>

<div class="flex justify-between gap-10 mt-10 items-center">
	<div class="w-6/12 order-2">
		<div class="shadow-md rounded-md p-5">
			<img
				src={bg}
                loading="lazy"
				alt="Show Case Cookie Banner API SCREENSHOT"
				width="100%"
				height="100%"
				class="w-full inline-block mr-2 border rounded-md"
			/>
			<div class="bg-red-50 p-2 rounded-md flex -ml-[1px] flex-wrap mt-2">
				<span>https://screenshotapi.dev/api/screenshot</span>
				<span class="text-red-500">?url=https://screenshotapi.dev</span>
				<span class="text-red-500" id="showcase2-response" />
			</div>
		</div>
	</div>
	<div class="w-6/12 order-1">
		<p class="font-semibold text-lg text-red-500">Seamless Capture Experience</p>
		<h2 class="text-4xl font-semibold heading">
			Eliminate Distractions with Cookie Banner Removal
		</h2>
		<p class="mt-2 text-muted-foreground">
			Our screenshot API offers a clutter-free solution with our intuitive Cookie Banner Removal
			feature. Capture clean, focused snapshots without the intrusive pop-ups. Ideal for creating
			professional-quality images, our API ensures:
		</p>
		<ul class="list-disc mt-5 ml-10 text-muted-foreground">
			<li>
				<strong>Uninterrupted Views</strong>:Automatically detect and dismiss cookie consent banners
				for a pristine screenshot.
			</li>
			<li>
				<strong>User Experience Preservation</strong>: Maintain the authenticity of the user
				experience by removing temporary overlays and pop-ups.
			</li>
		</ul>
		<p class="mt-5 text-muted-foreground">
			Activate this feature with a simple parameter in your API call, and let our system handle the
			rest, delivering screenshots that represent the true content of your pages, free from
			transient obstructions.
		</p>
	</div>
</div>
