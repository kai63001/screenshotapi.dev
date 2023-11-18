<script>
	import { onMount } from 'svelte';

	import Typewriter from 'typewriter-effect/dist/core';

    //@ts-ignore
    import bg from '$lib/assets/image/index/cookiebot_with_banner.jpg?format=webp&quality=100&w=573&h=456';

	function displayApiResponse() {
		const targetElement = document.getElementById('showcase1-response');

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
			.typeString('&profile=custom') // Type the width query parameter
			.pauseFor(1500) // Pause for 0.5 seconds before applying the next animation
            .deleteChars(6)
			.typeString('loginbypass') // Type the height query parameter
			.start()
			.pauseFor(2000) // Pause for 2 seconds before restarting the animation
			.deleteAll() // Delete the existing text
			.start();
	}
	onMount(() => {
		displayApiResponse();
	});
</script>

<div class="flex justify-between space-x-10 mt-10 items-center">
	<div class="w-6/12">
		<div class="shadow-md rounded-md p-5">
			<img
				src={bg}
                loading="lazy"
				alt="Show Case Profile API SCREENSHOT"
				width="100%"
				height="100%"
				class="w-full inline-block mr-2 border rounded-md"
			/>
            <div class="bg-red-50 p-2 rounded-md flex -ml-[1px] flex-wrap mt-2">
                <span>https://screenshotapi.dev/api/screenshot</span>
                <span class="text-red-500">?url=https://screenshotapi.dev</span>
                <span class="text-red-500" id="showcase1-response" />
            </div>
		</div>
	</div>
	<div class="w-6/12">
		<p class="font-semibold text-lg text-red-500">Tailored Browsing Environments</p>
		<h2 class="text-4xl font-semibold heading">Craft Your Perfect Snapshot Context</h2>
		<p class="mt-2 text-muted-foreground">
			Fine-tune your captures with our Custom Set feature. Define a 'custom' profile that lets you
			manipulate JavaScript, CSS, cookies, localStorage, and more directly through the URL. Just
			append <code
				class="relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold"
			>
				&profile=custom
			</code> to your API call, and our service will render the page according to your specifications.
			Perfect for:
		</p>
		<ul class="list-disc mt-5 ml-10 text-muted-foreground">
			<li>
				<strong>Personalized Styling</strong>: Apply custom CSS to alter the appearance of the
				webpage for your screenshot.
			</li>
			<li>
				<strong>Behavioral Scripting</strong>: Run specific JavaScript to capture dynamic states of
				the site.
			</li>
			<li>
				<strong>Session Simulation</strong>: Set cookies and localStorage values to replicate user
				sessions or test specific scenarios.
			</li>
		</ul>
	</div>
</div>
