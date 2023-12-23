<script>
	const listFeature = [
		{
			name: 'Format Image',
			link: ''
		},
		{
			name: 'Custom Header',
			link: ''
		}
	];
</script>

<footer class="border-t py-4 mt-10">
	<div class="grid grid-cols-1 md:px-0 px-4 md:grid-cols-4 py-5 max-w-xl md:max-w-7xl m-auto">
		<div>
			<h2 class="heading text-2xl">ScreenshotAPI</h2>
			<p class="text-muted-foreground mt-2">
				Capture Webpages Easily with Our Simple and Powerful ScreenshotAPI
			</p>
		</div>
		<div>
			<strong>
				Features
			</strong>
			<ul>
				{#each listFeature as feature}
					<li><a class="hover:underline text-red-600" href={feature.link}>{feature.name}</a></li>
				{/each}
			</ul>
		</div>
		<div>
			<strong>
				Resources
			</strong>
			<ul>
				{#each listFeature as feature}
					<li><a class="hover:underline text-red-600" href={feature.link}>{feature.name}</a></li>
				{/each}
			</ul>
		</div>
		<div>
			<strong>
				Legal
			</strong>
			<ul>
				{#each listFeature as feature}
					<li><a class="hover:underline text-red-600" href={feature.link}>{feature.name}</a></li>
				{/each}
			</ul>
		</div>
	</div>
	<div class="border-b" />
	<div class="max-w-xl md:max-w-7xl px-2 mt-3 m-auto text-center text-xs">
		<p>&copy; 2024 Screenshot Website API. All rights reserved.</p>
	</div>
</footer>
