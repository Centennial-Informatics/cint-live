<script lang="ts">
	import Page from '$lib/components/templates/page/page.svelte';
	import type { VerdictStatus } from 'src/types/submission';
	import VerdictIcon from '$lib/components/templates/icon/verdictIcon.svelte';
	import { GOOD, UNSUBMITTED } from '$lib/data/constants/status';
	import { verdictStatusColor } from '$lib/utils/verdictStatus';

	export let pointsPossible = 0;
	export let points = 0;
	export let status: VerdictStatus = UNSUBMITTED;
	export let verdict = 'Unsubmitted';
</script>

<Page>
	<div
		class="flex font-medium text-md shadow-lg {status === GOOD
			? 'text-green-200'
			: verdictStatusColor(status)} {status === GOOD
			? 'bg-green-700'
			: 'bg-gray-600'} p-4 rounded-xl items-center mt-8 -mb-8"
	>
		<div class="flex items-center flex-grow text-xl font-bold py-1">
			<VerdictIcon size="36" defaultIcon {status} />
			<div>{verdict}</div>
		</div>
		{points}/{pointsPossible} Points
	</div>
</Page>
