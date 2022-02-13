<script lang="ts">
	import Centered from '../centered.svelte';

	export let centered: boolean = false;
</script>

<div class={centered ? 'w-full' : ''}>
	<h6 class="font-semibold text-2xl md:text-xl mt-8 dark:text-gray-100">
		{#if centered}
			<Centered>
				<slot />
			</Centered>
		{:else}
			<slot />
		{/if}
	</h6>
</div>
