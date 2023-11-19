<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import * as Select from '$lib/components/ui/select';
	import Icon from '@iconify/svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import InputField from '$lib/components/InputField.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { Toaster, toast } from 'svelte-french-toast';
	import { pb, currentUser } from '$lib/pocketbase';
	import { onMount } from 'svelte';
	import CodeMirror from 'svelte-codemirror-editor';
	import { javascript } from '@codemirror/lang-javascript';
	import { css } from '@codemirror/lang-css';

	let customList = [];
	let selectedCustomSet: any = {};
	$: if (selectedCustomSet.value === 'custom') {
		selectedCustomSet = {};
		openDialog = true;
	}
	let selectedData = {};

	// check when selectedCustomSet change
	$: if (selectedCustomSet.value) {
		console.log(selectedCustomSet.value);
	}

	let openDialog = false;
	let createCustomSetName = '';

	const saveNewCustomSet = async () => {
		if (createCustomSetName.length === 0) {
			toast.error('Please enter a name', {
				duration: 1500,
				position: 'top-right'
			});
			return;
		}
		//create new custom set to database
		const data = await pb.collection('custom_sets').create({
			name: createCustomSetName,
			user_id: $currentUser.id
		});
		//add to list
		selectedCustomSet = {
			value: data.id,
			label: data.name
		};
		customList.push(selectedCustomSet);
		openDialog = false;
		toast.success('Custom SET created successfully!', {
			duration: 1500,
			position: 'top-right'
		});
	};

	onMount(async () => {
		//get custom set list
		const data = await pb.collection('custom_sets').getFullList();
		//map data and customList with name
		customList = data.map((item) => {
			return {
				value: item.id,
				label: item.name
			};
		});

		if (customList.length > 0) {
			selectedCustomSet = customList[0];
			selectedData = data[0];
		}
	});

	let lang_javascript = '';
    let lang_css = '';
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Custom SET Screenshot API</h1>
		<p class="text-mute mt-2">
			Customize every snapshot with our Custom SET Screenshot API. Tailor CSS, JS, and more for
			screenshots that meet your exact needs.
		</p>
	</div>
	<div class="bg-white p-5 rounded-md">
		<Label>Select a custom SET screenshot API.</Label>
		<Select.Root bind:selected={selectedCustomSet}>
			<Select.Trigger class="mt-2">
				<Select.Value placeholder="Select a Custom SET" />
			</Select.Trigger>
			<Select.Content>
				<Select.Group>
					{#each customList as fruit}
						<Select.Item value={fruit.value} label={fruit.label}>{fruit.label}</Select.Item>
					{/each}
					<Select.Separator />
					<Select.Item value="custom" label="Create New Custom SET">
						<Icon icon="mdi:plus" class="text-primary mr-2" width="20px" height="20px" />
						Create New Custom SET</Select.Item
					>
				</Select.Group>
			</Select.Content>
			<Select.Input name="customSet" />
		</Select.Root>

		<!-- {JSON.stringify(selectedCustomSet)} -->
	</div>
	<!-- main -->
	{#if selectedCustomSet.value}
		<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
			<div class="bg-white p-5 rounded-md">
				<h2>Javascript</h2>
				<CodeMirror bind:value={lang_javascript} lang={javascript()} />
			</div>
			<div class="bg-white p-5 rounded-md">
				<h2>CSS</h2>
				<CodeMirror bind:value={lang_javascript} lang={css()} />
			</div>
		</div>
	{/if}
</div>

<!-- outer -->
<Dialog.Root
	open={openDialog}
	onOpenChange={(e) => {
		openDialog = e;
	}}
>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>Create custom</Dialog.Title>
			<Dialog.Description>Create a new custom SET</Dialog.Description>
		</Dialog.Header>
		<div class="py-4">
			<InputField
				bind:value={createCustomSetName}
				label="Name"
				placeholder="Custom SET Name"
				required
			/>
		</div>
		<Dialog.Footer>
			<Button on:click={saveNewCustomSet}>Create</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
<Toaster />
