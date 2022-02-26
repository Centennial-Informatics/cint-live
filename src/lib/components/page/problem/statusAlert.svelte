<script lang="ts">
	import VerdictIcon from '$lib/components/templates/icon/verdictIcon.svelte';
	import { PENDING, UNSUBMITTED } from '$lib/data/constants/status';
	import { submissionData } from '$lib/data/stores/contestData';
	import verdictStatus, { verdictStatusBgColor } from '$lib/utils/verdictStatus';
	import type { VerdictStatus } from 'src/types/submission';

	export let problemName: string = '';
	export let problemID: string = '';
	let visible = false;
	let finished = true;
	let status: VerdictStatus = UNSUBMITTED;
	let verdict: string;
	$: $submissionData, showAlert(); // run showAlert() when submissionData updates
	function showAlert() {
		status = problemID in $submissionData ? verdictStatus($submissionData[problemID]) : UNSUBMITTED;
		verdict = problemID in $submissionData ? $submissionData[problemID].Verdict : '';

		if (status === PENDING) finished = false; // once submitted, show alert when done

		if (problemID in $submissionData && status !== PENDING && !finished) {
			visible = true;
			setTimeout(() => {
				visible = false;
			}, 5000);
			finished = true;
		}
	}
</script>

<div class="{visible ? 'opacity-100' : 'opacity-0'} transition-opacity duration-500">
	<div
		class="fixed bottom-6 right-8 z-50 text-lg {verdictStatusBgColor(
			status
		)} shadow-xl py-4 px-6 text-white font-medium rounded-xl flex flex-row items-center"
	>
		<VerdictIcon {status} />
		{problemName} —
		{verdict}
	</div>
</div>
