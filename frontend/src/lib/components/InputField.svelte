<script lang="ts">
	let isFocused = false;
	let inputRef: any;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	class={`relative block w-full ${isFocused ? 'bg-[#c5c8c9]' : 'bg-[#E4E9EC]'} duration-300 ${
		isFocused ? 'text-[#141414]' : 'text-[#666f75]'
	}   rounded py-1`}
	on:click={() => inputRef.focus()}
>
	<label for={$$props.id} class=" w-full px-4 text-xs"
		>{$$props.label}
		{#if $$props.required}
			<span class="text-red-500">*</span>
		{/if}
	</label>
	<input
		on:focus={() => (isFocused = true)}
		on:blur={() => (isFocused = false)}
		type={$$props.type || 'text'}
		id={$$props.id}
		name={$$props.name}
		bind:this={inputRef}
		{...$$props}
		class="w-full bg-transparent outline-none px-4 py-1 text-sm"
	/>
</div>
