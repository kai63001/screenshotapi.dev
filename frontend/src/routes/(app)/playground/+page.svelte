<script lang="ts">
	import InputField from '$lib/components/InputField.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { Switch } from '$lib/components/ui/switch';
	import Icon from '@iconify/svelte';



	$: url = 'https://unclelife.co';
	$: isFullScreen = false;
    $: innerWidth = 1280;
    $: innerHeight = 1024;
    $: delay = 2;
    let isCapturing = false;

	let screenshot =
		'https://underconstructionpage.com/wp-content/uploads/screenshots-article/screenshotapi.png';

	const takeScreenshot = async () => {
        isCapturing = true;
		const apiUrl = new URL(`${import.meta.env.VITE_API_KEY}/screenshot`);
		apiUrl.searchParams.append('url', url);
		if (isFullScreen) apiUrl.searchParams.append('fullScreen', 'true');
        if (innerWidth != 0) apiUrl.searchParams.append('innerWidth', innerWidth.toString())
        if (innerHeight != 0 && !isFullScreen) apiUrl.searchParams.append('innerHeight', innerHeight.toString())
        if (delay != 2) apiUrl.searchParams.append('delay', delay.toString())
		const response = await fetch(apiUrl.toString());
		const blob = await response.blob();
		const blobUrl = URL.createObjectURL(blob);
		screenshot = blobUrl;
        isCapturing = false;
	};

	$: APITextConverter = () => {
		const apiUrl = new URL(`${import.meta.env.VITE_API_KEY}/screenshot`);
		apiUrl.searchParams.append('url', url);
		if (isFullScreen) apiUrl.searchParams.append('fullScreen', 'true');
        if (innerWidth != 0) apiUrl.searchParams.append('innerWidth', innerWidth.toString())
        if (innerHeight != 0 && !isFullScreen) apiUrl.searchParams.append('innerHeight', innerHeight.toString())
        if (delay != 2 && delay) apiUrl.searchParams.append('delay', delay.toString())
		return apiUrl.toString();
	};
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">API Playground</h1>
		<p class="text-mute mt-2">
			Experiment with our API and see how easily you can capture high-quality screenshots of web
			pages.
		</p>
	</div>
	<div class="grid gap-4 grid-cols-1 lg:grid-cols-2">
		<div class="flex flex-col space-y-4">
			<div class="bg-white p-5 rounded-md">
				<InputField
					bind:value={url}
					icon="mdi:web"
					label="URL"
					required={true}
					placeholder="https://example.com"
				/>
				<p class="text-xs text-gray-500 py-1">Enter a webpage URL here</p>
				<div class="flex justify-end -mt-3">
					<button
						on:click|preventDefault={takeScreenshot}
                        disabled={isCapturing}
                        class:opacity-50={isCapturing}
						class="flex items-center bg-primary hover:bg-red-600 duration-200 text-white px-6 py-3 rounded-md"
					>
						<Icon class="text-white mr-1" icon="tabler:capture-filled" width="20px" height="20px" />
						Capture</button
					>
				</div>
			</div>
			<div class="bg-white p-5 rounded-md">
				<div class="grid grid-cols-2 gap-4">
					<InputField
						icon="material-symbols:width"
						label="Width"
						help="The browser window width."
						type="number"
						placeholder="1280"
                        bind:value={innerWidth}
					/>
					<InputField
						icon="material-symbols:height"
						label="Height"
						disabled={isFullScreen}
						help="The browser window height."
						type="number"
                        bind:value={innerHeight}
						placeholder="1024"
					/>
				</div>
				<div class="flex space-x-2 items-center mt-1">
					<Switch bind:checked={isFullScreen} id="full-screen" />
					<Label for="full-screen" class="text-gray-500">Full Screen</Label>
				</div>
			</div>
			<div class="bg-white p-5 rounded-md">
				<div class="grid grid-cols-2 gap-4">
					<InputField
						icon="mingcute:time-fill"
						label="Delay"
						help="Specify the delay in seconds before capturing the screenshot."
						type="number"
                        bind:value={delay}
						placeholder="2"
					/>
					<InputField
						icon="material-symbols:avg-time"
						label="Timeout"
						help="Should the site fail to respond within the set timeframe, the API request will be unsuccessful."
						type="number"
						placeholder="30"
					/>
				</div>
			</div>
		</div>
		<div class="bg-white p-5 rounded-md">
			<textarea
				rows="5"
				disabled
				class="text-mute mt-2 w-full overflow-auto bg-[#E4E9EC] p-2 text-sm cursor-text rounded"
				>{APITextConverter()}</textarea
			>
			<h2 class="text-xl font-semibold">Screenshot</h2>
            <div class="p-5 rounded-md bg-[#E4E9EC] mt-2">
                <img src={screenshot} alt="screenshot" class="w-full rounded-md" />
            </div>
		</div>
	</div>
</div>
