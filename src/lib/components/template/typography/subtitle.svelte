<script lang="ts">
	import Centered from '../centered.svelte';

	export let centered = false;
</script>

<div class="text-gray-500 dark:text-gray-400 text-sm">
	{#if centered}
		<Centered>
			<slot />
		</Centered>
	{:else}
		<slot />
	{/if}
</div>
