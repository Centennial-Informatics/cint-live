<script lang="ts">
	import CheckCircle from 'svelte-material-icons/CheckCircle.svelte';
	import CloseCircle from 'svelte-material-icons/CloseCircle.svelte';
	import AlertCircle from 'svelte-material-icons/AlertCircle.svelte';
	import HelpCircle from 'svelte-material-icons/HelpCircle.svelte';
	import { BAD, ERROR, GOOD, PENDING } from '$lib/data/constants/status';
	import { Colors } from '$lib/data/constants/colors';
	import Loading from '$lib/components/template/loading.svelte';
	import Page from '$lib/components/template/page/page.svelte';

	export let pointsPossible = 0;
	export let points = 0;
	export let status: string = '';
	export let verdict: string = 'Unsubmitted';

	const color: { [key: string]: string } = {};
	color[GOOD] = Colors.GOOD[50];
	color[BAD] = Colors.BAD[50];
	color[ERROR] = Colors.BAD[50];
	color[PENDING] = Colors.WARNING[50];
</script>

<Page>
	<div
		class="flex font-medium text-md shadow-lg text-gray-300 bg-gray-600 p-4 rounded-xl items-center mt-8 -mb-8"
	>
		<div
			class="flex items-center flex-grow text-xl font-bold py-1"
			style="color: {status ? color[status] : 'rgb(203 213 225)'};"
		>
			{#if status === GOOD}
				<span class="mr-2"><CheckCircle width="36" height="36" color={Colors.GOOD[50]} /></span>
			{:else if status === BAD}
				<span class="mr-2"><CloseCircle width="36" height="36" color={Colors.BAD[50]} /></span>
			{:else if status === ERROR}
				<span class="mr-2"><AlertCircle width="36" height="36" color={Colors.BAD[50]} /></span>
			{:else if status === PENDING}
				<span class="mr-2"><Loading size="36" color={Colors.WARNING[50]} /></span>
			{:else}
				<span class="mr-2"><HelpCircle width="36" height="36" color={Colors.NEUTRAL[0]} /></span>
			{/if}
			<div>{verdict}</div>
		</div>
		{points}/{pointsPossible} Points
	</div>
</Page>
