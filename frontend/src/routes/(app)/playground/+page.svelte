<script lang="ts">
	import InputField from '$lib/components/InputField.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { Switch } from '$lib/components/ui/switch';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { Toaster, toast } from 'svelte-french-toast';
	import * as Select from '$lib/components/ui/select';
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	let access_key = '';

	let url = 'https://unclelife.co';
	let isFullScreen = false;
	let scrollDelay = 1;
	let innerWidth = 1280;
	let innerHeight = 1024;
	let delay = 0;
	let timeout = 60;
	let noAds = false;
	let noCookie = false;
	let blockTracker = false;
	let async = false;
	let saveToS3 = false;
	let path_file_name = '';
	let quality = 100;
	let formatImage = {
		value: 'png',
		label: 'PNG'
	};

	let formatImageList = [
		{
			value: 'png',
			label: 'PNG'
		},
		{
			value: 'jpg',
			label: 'JPG'
		},
		{
			value: 'jpeg',
			label: 'JPEG'
		},
		{
			value: 'webp',
			label: 'WEBP'
		},
		{
			value: 'pdf',
			label: 'PDF'
		},
		{
			value: 'svg',
			label: 'SVG'
		}
	];

	let isCapturing = false;

	let screenshot = '';

	let customList = [];
	let selectedCustomSet: any = {};
	let fullCustomData = [];

	let dataRespnse = {};

	let selectedResponse = {
		value: 'image',
		label: 'Image'
	};

	$: if (selectedCustomSet.value === 'custom') {
		goto('/custom-set');
	} else if (selectedCustomSet.value === 'none') {
		selectedCustomSet = {};
	}

	const takeScreenshot = async () => {
		dataRespnse = {};
		isCapturing = true;
		const apiUrl = apiText;
		const response = await fetch(apiUrl.toString());
		const blob = await response.blob();
		if (blob.type === 'application/json') {
			const json = await blob.text();
			const data = JSON.parse(json);
			console.log(data);
			if (data.status) {
				toast.success(data.status, {
					duration: 3000,
					position: 'top-right'
				});
				dataRespnse = data;
				isCapturing = false;
				return;
			}
			toast.error(data.message, {
				duration: 3000,
				position: 'top-right'
			});
		} else {
			const blobUrl = URL.createObjectURL(blob);
			screenshot = blobUrl;
		}
		isCapturing = false;
	};

	$: apiText = APITextConverterDuplicate();

	$: APITextConverterDuplicate = () => {
		const apiUrl = new URL(`${import.meta.env.VITE_API_KEY}/screenshot`);
		apiUrl.searchParams.append('access_key', access_key);
		apiUrl.searchParams.append('url', url);
		if (isFullScreen) apiUrl.searchParams.append('full_screen', 'true');
		if (scrollDelay != 1 && isFullScreen)
			apiUrl.searchParams.append('scroll_delay', scrollDelay.toString());
		if (innerWidth != 0) apiUrl.searchParams.append('v_width', innerWidth.toString());
		if (innerHeight != 0 && !isFullScreen)
			apiUrl.searchParams.append('v_height', innerHeight.toString());
		if (delay != 0 && delay) apiUrl.searchParams.append('delay', delay.toString());
		if (timeout != 0 && timeout != 60) apiUrl.searchParams.append('timeout', timeout.toString());
		if (noAds) apiUrl.searchParams.append('no_ads', 'true');
		if (noCookie) apiUrl.searchParams.append('no_cookie_banner', 'true');
		if (blockTracker) apiUrl.searchParams.append('block_tracker', 'true');
		if (async) apiUrl.searchParams.append('async', 'true');
		if (selectedCustomSet.label) apiUrl.searchParams.append('custom', selectedCustomSet.label);
		if (selectedResponse.value && selectedResponse.value != 'image')
			apiUrl.searchParams.append('response_type', selectedResponse.value);
		if (saveToS3) apiUrl.searchParams.append('save_to_s3', 'true');
		if (path_file_name) apiUrl.searchParams.append('path_file_name', path_file_name);
		if (formatImage.value && formatImage.value != 'png')
			apiUrl.searchParams.append('format', formatImage.value);
		if (quality && quality != 100) apiUrl.searchParams.append('quality', quality.toString());

		return apiUrl.toString();
	};

	onMount(async () => {
		access_key = localStorage.getItem('access_key') || '';

		const data = await pb.collection('custom_sets').getFullList();
		//map data and customList with name
		fullCustomData = data;
		customList = data.map((item) => {
			return {
				value: item.id,
				label: item.name
			};
		});
	});

	const checkCustomSetHasS3 = () => {
		if (selectedCustomSet.value === 'custom') return false;
		if (selectedCustomSet.value === 'none') return false;
		if (!selectedCustomSet.value) return false;
		const data = fullCustomData.find((item) => item.id === selectedCustomSet.value);
		if (data && data.bucket_endpoint && data.bucket_access_key && data.bucket_secret_key)
			return true;

		return false;
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
				{#if isFullScreen}
					<div class="grid grid-cols-2 gap-4 mt-4">
						<InputField
							icon="material-symbols:height"
							label="Scroll Delay (Seconds)"
							disabled={!isFullScreen}
							help="Delay in seconds before scrolling to the next section."
							type="number"
							bind:value={scrollDelay}
							placeholder="1"
						/>
					</div>
				{/if}
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
						bind:value={timeout}
						placeholder="30"
					/>
				</div>
			</div>
			<div class="bg-white p-5 rounded-md">
				<div class="grid grid-cols-2 gap-4">
					<div>
						<label for="format" class="text-muted-foreground text-sm">Format</label>
						<Select.Root bind:selected={formatImage}>
							<Select.Trigger class="mt-2">
								<Select.Value placeholder="Format Image" />
							</Select.Trigger>
							<Select.Content>
								<Select.Group>
									{#each formatImageList as fruit}
										<Select.Item value={fruit.value} label={fruit.label}>{fruit.label}</Select.Item>
									{/each}
								</Select.Group>
							</Select.Content>
							<Select.Input id="format" name="format" />
						</Select.Root>
					</div>
					<div>
						<label for="format" class="text-muted-foreground text-sm">Quality</label>
						<input
							class="w-full block form border rounded px-3 py-1.5 mt-2"
							placeholder="100"
							type="number"
							bind:value={quality}
						/>
					</div>
				</div>
			</div>
			<div class="bg-white p-5 rounded-md">
				<div class="grid grid-cols-2 gap-4">
					<div>
						<label for="customSet" class="text-gray-500 text-sm">Custom Set</label>
						<Select.Root bind:selected={selectedCustomSet}>
							<Select.Trigger class="mt-2">
								<Select.Value placeholder="Custom Set" />
							</Select.Trigger>
							<Select.Content>
								<Select.Group>
									<Select.Item value="none" label="none">None</Select.Item>
									<Select.Separator />
									{#each customList as fruit}
										<Select.Item value={fruit.value} label={fruit.label}>{fruit.label}</Select.Item>
									{/each}
									<Select.Separator />
									<Select.Item value="custom" label="Create New Custom SET">
										<Icon icon="mdi:plus" class="text-primary mr-2" width="20px" height="20px" />
										Create New</Select.Item
									>
								</Select.Group>
							</Select.Content>
							<Select.Input id="customSet" name="customSet" />
						</Select.Root>
					</div>
					<div>
						<label for="response" class="text-gray-500 text-sm"> Response </label>
						<Select.Root bind:selected={selectedResponse}>
							<Select.Trigger class="mt-2">
								<Select.Value placeholder="Response" />
							</Select.Trigger>
							<Select.Content>
								<Select.Group>
									<Select.Separator />
									<Select.Item value={'image'} label={'Image'}>Image</Select.Item>
									<Select.Item value={'json'} label={'JSON'}>JSON</Select.Item>
								</Select.Group>
							</Select.Content>
							<Select.Input name="response" id="response" />
						</Select.Root>
					</div>
				</div>
				<!-- check if when custom set have s3 data -->
				{#if selectedCustomSet.value && checkCustomSetHasS3()}
					<div class="flex flex-col mt-2 space-y-2">
						<div class="flex items-center space-x-3">
							<Switch bind:checked={saveToS3} id="save-to-s3" />
							<Label for="save-to-s3" class="text-gray-500">Save to S3</Label>
						</div>
					</div>
					{#if saveToS3}
						<div class="grid grid-cols-2 gap-4 mt-4">
							<InputField
								icon="ph:path"
								label="Path & File Name ({formatImage?.value?.toUpperCase()})"
								help="File name of the screenshot. Empty for random name."
								type="text"
								placeholder="screenshots/screenshot_github"
								bind:value={path_file_name}
							/>
						</div>
					{/if}
				{/if}
			</div>

			<div class="bg-white p-5 rounded-md flex-col flex space-y-2">
				<div class="flex items-center space-x-3">
					<Switch bind:checked={async} id="no-async" />
					<Label for="no-async" class="text-gray-500">Async Screenshot Request</Label>
				</div>
			</div>
			<div class="bg-white p-5 rounded-md flex-col flex space-y-2">
				<div class="flex items-center space-x-3">
					<Switch bind:checked={noAds} id="no-ads" />
					<Label for="no-ads" class="text-gray-500">Block ads</Label>
				</div>
				<div class="flex items-center space-x-3">
					<Switch bind:checked={noCookie} id="no-cookie" />
					<Label for="no-cookie" class="text-gray-500">Block Cookie Popups</Label>
				</div>
				<div class="flex items-center space-x-3">
					<Switch bind:checked={blockTracker} id="no-trakcer" />
					<Label for="no-trakcer" class="text-gray-500">Block Tracker</Label>
				</div>
			</div>
		</div>
		<div class="bg-white p-5 rounded-md">
			<textarea
				rows="5"
				disabled
				class="text-mute mt-2 w-full overflow-auto bg-[#E4E9EC] hover:bg-[#d3d4d4] p-2 text-sm cursor-text rounded"
				>{apiText}</textarea
			>
			{#if Object.keys(dataRespnse).length > 0}
				<h2 class="text-xl font-semibold">API Response</h2>
				<textarea
					rows="5"
					disabled
					class="text-mute mt-2 w-full overflow-auto bg-[#E4E9EC] hover:bg-[#d3d4d4] p-2 text-sm cursor-text rounded"
					>{JSON.stringify(dataRespnse, null, 2)}</textarea
				>
			{:else}
				<h2 class="text-xl font-semibold">Screenshot</h2>
				{#if screenshot}
					<div class="p-5 rounded-md bg-[#E4E9EC] mt-2">
						<img src={screenshot} alt="screenshot" class="w-full rounded-md" />
					</div>
				{/if}
			{/if}
			<!-- <h2 class="text-xl font-semibold">Screenshot</h2>
			{#if screenshot}
				<div class="p-5 rounded-md bg-[#E4E9EC] mt-2">
					<img src={screenshot} alt="screenshot" class="w-full rounded-md" />
				</div>
			{/if} -->
		</div>
	</div>
</div>
<Toaster />
