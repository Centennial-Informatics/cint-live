<script lang="ts">
	import { currentPage } from '$lib/data/stores/currentPage';
	import Home from 'svelte-material-icons/Home.svelte';
	import Poll from 'svelte-material-icons/Poll.svelte';
	import CheckCircle from 'svelte-material-icons/CheckCircle.svelte';
	import CloseCircle from 'svelte-material-icons/CloseCircle.svelte';
	import AlertCircle from 'svelte-material-icons/AlertCircle.svelte';
	import { BAD, ERROR, GOOD, PENDING } from '$lib/data/constants/status';
	import { Colors } from '$lib/data/constants/colors';
	import Loading from '$lib/components/template/loading.svelte';

	export let to = '/';
	export let tabId = '';
	export let status = '';

	let selected = false;
	$: selected = tabId === $currentPage;
</script>

<a href={to} on:click={() => currentPage.set(tabId)}>
	<div
		class="flex items-center no-underline text-gray-300 bg-brand bg-opacity-0 p-2 pl-5 max-w-full rounded-r-full text-base font-medium font-body transition-colors {!selected
			? 'hover:bg-opacity-40'
			: 'bg-opacity-70 text-brand-lightest'}"
	>
		{#if tabId === 'home'}
			<span class="mr-2"><Home width="20" height="20" /></span>
		{:else if tabId === 'standings'}
			<span class="mr-2"><Poll width="20" height="20" /></span>
		{:else if status === GOOD}
			<span class="mr-2"><CheckCircle width="20" height="20" color={Colors.GOOD[50]} /></span>
		{:else if status === BAD}
			<span class="mr-2"><CloseCircle width="20" height="20" color={Colors.BAD[50]} /></span>
		{:else if status === ERROR}
			<span class="mr-2"><AlertCircle width="20" height="20" color={Colors.BAD[0]} /></span>
		{:else if status === PENDING}
			<span class="mr-2"><Loading size="20" color={Colors.WARNING[50]} /></span>
		{/if}
		<span class="max-w-10 overflow-hidden overflow-ellipsis whitespace-nowrap">
			{#if !status}
				<span class="w-2" />
				{tabId + '. '}
			{/if}
			<slot />
		</span>
	</div>
</a>
