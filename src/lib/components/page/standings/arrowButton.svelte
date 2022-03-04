<script lang="ts">
	import ChevronLeft from 'svelte-material-icons/ChevronLeft.svelte';
	import ChevronRight from 'svelte-material-icons/ChevronRight.svelte';
	export let onClick: (e: Event) => void = () => {};
	export let disabled = false;
	export let direction = 'L';
	export let size = '20';
	export let mode = 'DARK'; // will display white in dark mode

	let color = '#64748b';

	$: if (!disabled) {
		if (mode === 'DARK') color = '#e2e8f0';
		else color = '#1e293b';
	} else {
		color = '#64748b';
	}
</script>

<button type="button" {disabled} on:click={onClick}>
	{#if direction === 'L'}
		<ChevronLeft width={size} height={size} {color} />
	{:else}
		<ChevronRight width={size} height={size} {color} />
	{/if}
</button>
