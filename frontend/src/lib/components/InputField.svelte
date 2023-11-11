<script lang="ts">
	import Icon from '@iconify/svelte';

	let isFocused = false;
	let inputRef: any;
	export let value:any = undefined;
	export let type: string = 'text';
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div>
	<div
		class:bg-gray-100={$$props.disabled}
		class={`relative block w-full ${isFocused ? 'bg-[#c5c8c9]' : 'bg-[#E4E9EC]'} duration-300 ${
			isFocused ? 'text-[#141414]' : 'text-[#666f75]'
		}   rounded py-1`}
		on:click={() => inputRef.focus()}
	>
		<label for={$$props.id} class=" w-full px-4 py-2 text-xs flex items-center">
			{#if $$props.icon}
				<Icon class="text-gray-500 mr-1 -mt-0.5" icon={$$props.icon} width="15px" height="15px" />
			{/if}
			{$$props.label}
			{#if $$props.required}
				<span class="text-red-500 ml-1">*</span>
			{/if}
		</label>
		<input
			on:focus={() => (isFocused = true)}
			on:blur={() => (isFocused = false)}
			id={$$props.id}
			name={$$props.name}
			bind:this={inputRef}
			bind:value
			{...$$props}
			class="w-full bg-transparent outline-none px-4 py-1 text-sm"
		/>
	</div>
	{#if $$props.help}
		<p class="text-xs text-gray-500 py-1">{$$props.help}</p>
	{/if}
</div>
