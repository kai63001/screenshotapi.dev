<script>
	import { onMount } from 'svelte';

	import Typewriter from 'typewriter-effect/dist/core';

	function displayApiResponse() {
		const targetElement = document.getElementById('showcase3-response');

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
			.typeString('&store_s3=true&storage_path=snap') // Type the width query parameter
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

<div class="flex justify-between space-x-10 mt-10 items-center">
	<div class="w-6/12">
		<div class="shadow-md rounded-md p-5">
			<img
				src={'https://screenshotone.com/cookiebot_with_banner.jpg'}
				alt="Show Case 1"
				width="100%"
				height="100%"
				class="w-full inline-block mr-2 border rounded-md"
			/>
			<div class="bg-red-50 p-2 rounded-md flex -ml-[1px] flex-wrap mt-2">
				<span>https://screenshotapi.dev/api/screenshot</span>
				<span class="text-red-500">?url=https://apple.com/mac</span>
				<span class="text-red-500" id="showcase3-response" />
			</div>
		</div>
	</div>
	<div class="w-6/12">
		<p class="font-semibold text-lg text-red-500">Direct S3 Bucket Integration</p>
		<h2 class="text-4xl font-semibold heading">Seamless Storage with One Click</h2>
		<p class="mt-2 text-muted-foreground">
			With our screenshot API, you can effortlessly save your captures directly to an Amazon S3
			bucket. This feature streamlines your workflow by automating the storage process, ensuring
			that your images are securely archived without additional steps. Here's how you benefit:
		</p>
		<ul class="list-disc mt-5 ml-10 text-muted-foreground">
			<li>
				<strong>Hassle-Free Configuration</strong>: Set up once with your S3 bucket details, and our
				API will automatically deposit screenshots right into your cloud storage.
			</li>
			<li>
				<strong>Scalable Storage Solution</strong>: Leverage S3's vast storage capacity and
				durability to manage and store large volumes of screenshots without compromising on
				performance.
			</li>
			<li>
				<strong>Secure Transfer</strong>: Screenshots are transferred using secure methods, ensuring
				that your data remains confidential and intact during transit.
			</li>
			<li>
				<strong>Immediate Accessibility</strong>: Access your screenshots from anywhere, at any
				time, thanks to S3's reliable and globally available infrastructure.
			</li>
		</ul>
		<p class="mt-5 text-muted-foreground">
			Activate this feature with a simple addition to your API call, and rest assured that your
			screenshots are saved instantly to your designated bucket, ready for retrieval or distribution
			according to your needs.
		</p>
	</div>
</div>
