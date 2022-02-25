<script lang="ts">
	import { currentPage } from '$lib/data/stores/currentPage';
	import Home from 'svelte-material-icons/Home.svelte';
	import Poll from 'svelte-material-icons/Poll.svelte';

	export let to = '/';
	export let tabId = '';

	let selected = false;
	$: selected = tabId === $currentPage;
</script>

<a href={to} on:click={() => currentPage.set(tabId)}>
	<div
		class="flex items-center no-underline text-gray-300 bg-brand bg-opacity-0 p-2 max-w-full rounded-md text-base font-medium font-body transition-colors {!selected &&
			'hover:bg-opacity-40'} {selected && 'bg-opacity-80'}"
	>
		{#if tabId === 'home'}
			<span class="mr-2"><Home width="20" height="20" /></span>
		{:else if tabId === 'standings'}
			<span class="mr-2"><Poll width="20" height="20" /></span>
		{/if}
		<span class="max-w-10 overflow-hidden overflow-ellipsis whitespace-nowrap">
			<slot />
		</span>
	</div>
</a>
