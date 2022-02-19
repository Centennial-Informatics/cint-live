<script lang="ts">
	import { page } from '$app/stores';
	import ProblemStatement from '$lib/components/page/problem/problemStatement.svelte';
	import SubmitForm from '$lib/components/page/problem/submit/submitForm.svelte';
	import Wrapper from '$lib/components/template/page/wrapper.svelte';
	import { problemNames, problemPages } from '$lib/data/stores/contestData';
	import { currentPage } from '$lib/data/stores/currentPage';
	import { onMount } from 'svelte';
	onMount(() => {
		if (!$currentPage) currentPage.set($page.params.slug);
	});
</script>

<Wrapper>
	{#each $problemNames as problem (problem.ID)}
		<ProblemStatement
			pageData={$problemPages[problem.ID]}
			visible={$page.params.slug === problem.ID}
		/>
	{/each}
	{#if $problemPages[$page.params.slug]}
		<SubmitForm title={$problemPages[$page.params.slug].Header.Title} ID={$page.params.slug} />
	{/if}
</Wrapper>
