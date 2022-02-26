<script lang="ts">
	import { page } from '$app/stores';
	import ProblemStatement from '$lib/components/page/problem/problemStatement.svelte';
	import StatusAlert from '$lib/components/page/problem/statusAlert.svelte';
	import SubmissionStatus from '$lib/components/page/problem/submissionStatus.svelte';
	import SubmitForm from '$lib/components/page/problem/submit/submitForm.svelte';
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import { UNSUBMITTED } from '$lib/data/constants/status';
	import {
		contestPoints,
		problemNames,
		problemPages,
		submissionData
	} from '$lib/data/stores/contestData';
	import { currentPage } from '$lib/data/stores/currentPage';
	import verdictStatus from '$lib/utils/verdictStatus';
	import type { VerdictStatus } from 'src/types/submission';
	import { onMount } from 'svelte';
	onMount(() => {
		if (!$currentPage) currentPage.set($page.params.slug);
	});

	let status: VerdictStatus = UNSUBMITTED,
		verdict: string,
		points: number,
		pointsPossible: number;
	let st: Element;

	$: status =
		$page.params.slug in $submissionData
			? verdictStatus($submissionData[$page.params.slug])
			: UNSUBMITTED;
	$: verdict =
		$page.params.slug in $submissionData
			? $submissionData[$page.params.slug].Verdict
			: 'Unsubmitted';
	points = 0;
	$: pointsPossible = $page.params.slug in $contestPoints ? $contestPoints[$page.params.slug] : 0;
</script>

<div bind:this={st}>
	<Wrapper>
		{#each $problemNames as problem (problem.ID)}
			<StatusAlert
				problemName={$problemPages[problem.ID] ? $problemPages[problem.ID].Header.Title : ''}
				problemID={problem.ID}
			/>
			{#if $page.params.slug === problem.ID}
				<SubmissionStatus {pointsPossible} {points} {status} {verdict} />
				<ProblemStatement pageData={$problemPages[problem.ID]} visible={true} />
			{/if}
		{/each}
		{#if $problemPages[$page.params.slug]}
			<SubmitForm
				title={$problemPages[$page.params.slug].Header.Title}
				ID={$page.params.slug}
				submissionStatus={st}
			/>
		{/if}
	</Wrapper>
</div>
