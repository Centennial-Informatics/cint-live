<script lang="ts">
	import Centered from '../centered.svelte';

	export let centered = false;
	export let margin = true;
</script>

<h2 class="{margin ? 'my-6' : ''} text-6xl font-semibold md:text-5xl dark:text-gray-100">
	{#if centered}
		<Centered>
			<slot />
		</Centered>
	{:else}
		<slot />
	{/if}
</h2>
