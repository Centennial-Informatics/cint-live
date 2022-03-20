<script lang="ts">
	import Header from '$lib/components/templates/page/header.svelte';
	import Page from '$lib/components/templates/page/page.svelte';
	import Code from '$lib/components/templates/typography/code.svelte';
	import Image from '$lib/components/templates/typography/image.svelte';
	import Paragraph from '$lib/components/templates/typography/paragraph.svelte';
	import SectionTitle from '$lib/components/templates/typography/sectionTitle.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import type { ProblemPage } from 'src/types/contestData';
	import { onMount } from 'svelte';
	import SampleTest from './sampleTest.svelte';

	export let pageData: ProblemPage;
	export let visible = false;
	onMount(() => {
		if (visible) {
			// @ts-expect-error Cannot find name 'MathJax'.
			if (typeof MathJax !== 'undefined' && MathJax.typeset) {
				// @ts-expect-error Cannot find name 'MathJax'.
				MathJax.typeset();
			}
		}
	});
	$: if (visible) {
		// @ts-expect-error Cannot find name 'MathJax'.
		if (typeof MathJax !== 'undefined' && MathJax.typeset) {
			// @ts-expect-error Cannot find name 'MathJax'.
			MathJax.typeset();
		}
	}
</script>

<Page {visible}>
	{#if pageData}
		<Header>
			<Title>{pageData.Header.Title}</Title>
			<Subtitle>Time limit - {pageData.Header.TimeLimit}</Subtitle>
			<Subtitle>Memory limit - {pageData.Header.MemLimit}</Subtitle>
			<Subtitle>Input - {pageData.Header.Input}</Subtitle>
			<Subtitle>Output - {pageData.Header.Output}</Subtitle>
		</Header>
		{#each pageData.Statement as paragraph (paragraph.Image + paragraph.Code + paragraph.Text)}
			{#if paragraph.Image !== ''}
				<Image src={paragraph.Image} alt="Example" />
			{:else if paragraph.Code !== ''}
				<Code code={paragraph.Code} />
			{:else}
				<Paragraph>{paragraph.Text}</Paragraph>
			{/if}
		{/each}
		<SectionTitle>Input</SectionTitle>
		{#each pageData.InputSpecification as paragraph (paragraph.Image + paragraph.Code + paragraph.Text)}
			{#if paragraph.Image !== ''}
				<Image src={paragraph.Image} alt="Example" />
			{:else if paragraph.Code !== ''}
				<Code code={paragraph.Code} />
			{:else}
				<Paragraph>{paragraph.Text}</Paragraph>
			{/if}
		{/each}
		<SectionTitle>Output</SectionTitle>
		{#each pageData.OutputSpecification as paragraph (paragraph.Image + paragraph.Code + paragraph.Text)}
			{#if paragraph.Image !== ''}
				<Image src={paragraph.Image} alt="Example" />
			{:else if paragraph.Code !== ''}
				<Code code={paragraph.Code} />
			{:else}
				<Paragraph>{paragraph.Text}</Paragraph>
			{/if}
		{/each}
		<SectionTitle>Examples</SectionTitle>
		{#each pageData.SampleTests as sample, i (sample.Input)}
			<div class="text-gray-200 mt-4">
				Sample Test {i + 1}
				<SampleTest input={sample.Input} output={sample.Output} />
			</div>
		{/each}
		{#if pageData.Note.length > 0}
			<SectionTitle>Note</SectionTitle>
			{#each pageData.Note as paragraph (paragraph.Image + paragraph.Code + paragraph.Text)}
				{#if paragraph.Image !== ''}
					<Image src={paragraph.Image} alt="Example" />
				{:else if paragraph.Code !== ''}
					<Code code={paragraph.Code} />
				{:else}
					<Paragraph>{paragraph.Text}</Paragraph>
				{/if}
			{/each}
		{/if}
	{/if}
</Page>
