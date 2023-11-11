<script>
	import Button from '$lib/components/ButtonCustom.svelte';
	import InputField from '$lib/components/InputField.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { Switch } from '$lib/components/ui/switch';
	import Icon from '@iconify/svelte';

    let isFullScreen = false;

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
				<InputField icon="mdi:web" label="URL" required={true} placeholder="https://example.com" />
				<p class="text-xs text-gray-500 py-1">Enter a webpage URL here</p>
				<div class="flex justify-end">
					<Button class="flex items-center">
						<Icon class="text-white mr-1" icon="tabler:capture-filled" width="20px" height="20px" />
						Capture</Button
					>
				</div>
			</div>
			<div class="bg-white p-5 rounded-md">
				<div class="grid grid-cols-2 gap-4">
					<InputField
						icon="mdi:web"
						label="Width"
						help="The browser window width."
						type="number"
						required={true}
						placeholder="1280"
					/>
					<InputField
						icon="mdi:web"
						label="Height"
						disabled={isFullScreen}
						help="The browser window height."
						type="number"
						required={true}
						placeholder="1024"
					/>
				</div>
				<div class="flex space-x-2 items-center mt-1">
					<Switch
                        bind:checked={isFullScreen}
                    />
					<Label class="text-gray-500">Full Screen</Label>
				</div>
			</div>
		</div>
		<div class="bg-white p-5 rounded-md">
			<h1 class="font-bold text-2xl">API Playground</h1>
			<p class="text-mute mt-2">
				Experiment with our API and see how easily you can capture high-quality screenshots of web
				pages.
			</p>
		</div>
	</div>
</div>
