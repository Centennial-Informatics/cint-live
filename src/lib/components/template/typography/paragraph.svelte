<script lang="ts">
	import Centered from '../centered.svelte';

	export let centered: boolean = false;
</script>

<p class="font-body text-lg my-4 md:text-base leading-relaxed md:leading-normal dark:text-gray-300">
	{#if centered}
		<Centered>
			<slot />
		</Centered>
	{:else}
		<slot />
	{/if}
</p>
