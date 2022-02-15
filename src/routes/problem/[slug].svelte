<script lang="ts">
	import { page } from '$app/stores';
	import ProblemStatement from '$lib/components/page/problem/problemStatement.svelte';
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
		<div class={$page.params.slug === problem.ID ? 'block' : 'hidden'}>
			<ProblemStatement pageData={$problemPages[problem.ID]} />
		</div>
	{/each}
</Wrapper>
