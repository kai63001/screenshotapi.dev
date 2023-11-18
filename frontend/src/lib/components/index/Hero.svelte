<script lang="ts">
	//@ts-ignore
	import bg from '$lib/assets/image/index/881fcd8b-8f01-4d8d-8656-10544c588efd.png?format=webp&quality=100&w=600&h=317';
	//@ts-ignore
	import { onMount } from 'svelte';

	import Typewriter from 'typewriter-effect/dist/core';

	function displayApiResponse() {
		const targetElement = document.getElementById('api-response');

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
			.typeString('&width=1920') // Type the width query parameter
			.pauseFor(1500) // Pause for 0.5 seconds before applying the next animation
			.typeString('&height=1080') // Type the height query parameter
			.start()
			.pauseFor(2000) // Pause for 2 seconds before restarting the animation
			.deleteAll() // Delete the existing text
			.start();
	}
	onMount(() => {
		displayApiResponse();
	});
</script>

<div class="shadow-xl rounded-md duration-500">
	<div class="p-3 duration-500">
		<img
			loading="eager"
			id="imager"
            src={bg}
			class="object-contain shadow-lg rounded-md transition-all duration-500 ease-in-out transform hover:scale-105 hover:shadow-xl cursor-pointer"
			alt="ScreenshotAPI example screenshot"
		/>
	</div>
	<div class="p-3">
		<div class="bg-red-50 p-2 rounded-md flex -ml-[1px] flex-wrap">
			<span>https://screenshotapi.dev/api/screenshot</span>
			<span class="text-red-500">?url=https://screenshotapi.dev</span>
			<span class="text-red-500" id="api-response" />
		</div>
	</div>
</div>
