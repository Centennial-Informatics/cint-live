<script lang="ts">
	import { currentPage } from '$lib/data/stores/currentPage';
	import Home from 'svelte-material-icons/Home.svelte';
	import Poll from 'svelte-material-icons/Poll.svelte';
	import VerdictIcon from '$lib/components/templates/icon/verdictIcon.svelte';
	import type { VerdictStatus } from 'src/types/submission';
	import { BAD, ERROR, GOOD, PENDING, UNSUBMITTED } from '$lib/data/constants/status';

	export let to = '/';
	export let tabId = '';
	export let status: VerdictStatus = UNSUBMITTED;

	let selected = false;
	$: selected = tabId === $currentPage;
</script>

<a href={to} on:click={() => currentPage.set(tabId)}>
	<div
		class="flex items-center no-underline bg-opacity-0 md:p-2 md:pl-5 p-4 max-w-full md:rounded-r-full text-base font-medium font-body transition-colors {!selected
			? 'hover:bg-opacity-60'
			: `bg-opacity-100`}
			{status == GOOD ? 'bg-good-100 text-good-50' : ''}
			{status == BAD || status == ERROR ? 'bg-gray-600 text-bad-50' : ''}
			{status == PENDING ? 'bg-gray-600 text-warning-50' : ''}
			{status == UNSUBMITTED ? `bg-gray-600 text-gray-300` : ''}
			{tabId === 'home' || tabId == 'standings' ? 'bg-gray-600 text-gray-300' : ''}
			"
	>
		{#if tabId === 'home'}
			<span class="mr-2"><Home width="20" height="20" /></span>
		{:else if tabId === 'standings'}
			<span class="mr-2"><Poll width="20" height="20" /></span>
		{:else}
			<VerdictIcon size={'20'} defaultIcon={false} {status} />
		{/if}
		<span class="max-w-10 overflow-hidden overflow-ellipsis whitespace-nowrap">
			<!-- {#if !status} -->
			<span class="w-2" />
			{#if tabId !== 'home' && tabId !== 'standings'}
				{tabId + '. '}
			{/if}
			<!-- {/if} -->
			<slot />
		</span>
	</div>
</a>
