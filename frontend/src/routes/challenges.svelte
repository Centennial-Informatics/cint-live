<script lang="ts">
	import Footer from '$lib/components/layout/footer/footer.svelte';
	import ChallengeLinks from '$lib/components/page/challenges/challengeLinks.svelte';
	import Header from '$lib/components/templates/page/header.svelte';
	import Page from '$lib/components/templates/page/page.svelte';
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import Paragraph from '$lib/components/templates/typography/paragraph.svelte';
	import SectionTitle from '$lib/components/templates/typography/sectionTitle.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import { contestData, contestInfo } from '$lib/data/stores/contestData';
	import { currentPage } from '$lib/data/stores/currentPage';
	import CollectChallenges from '$lib/utils/networking/challenges';
	import { onMount } from 'svelte';

	let challengeLinks: string[] = [];

	async function fetchLinks() {
		challengeLinks = await CollectChallenges();
	}

	onMount(() => {
		if (!$currentPage) currentPage.set('challenges');
		fetchLinks();
	});
</script>

<Wrapper transparent>
	<Page>
		{#if $contestData.Info}
			<Header>
				<Title>Challenges</Title>
			</Header>
			<Paragraph>
				Challenges are a set of short open-ended programming challenges requiring both creativity
				and technical skill.
			</Paragraph>
			<Paragraph
				>There will be 8 total challenges released periodically throughout the afternoon. An alert
				will be sent out when a new challenge is released, so be sure to return to this page to
				access the challenge links!</Paragraph
			>
			<Paragraph>
				<strong
					>Once a new challenge is released, no more submissions are allowed for previous
					challenges.</strong
				>
			</Paragraph>
			<Paragraph>
				<strong>
					Challenges are separate from the actual competition and award raffle tickets. They will
					not affect the competition standings. Challenges are the same for both divisions.</strong
				>
			</Paragraph>
			<SectionTitle>Challenge Links</SectionTitle>
			<Subtitle>Only 1 challenge link is active at a time.</Subtitle>
			<ChallengeLinks links={challengeLinks} />
			<SectionTitle>Challenge Release Schedule</SectionTitle>
			<Subtitle
				>Challenges are due before the release of the next challenge or the end of the competition.</Subtitle
			>
			<Paragraph>
				<div class="flex flex-col">
					<span> 4:15pm - Challenge 1</span>
					<span>5:00pm - Challenge 2</span>
					<span>5:30pm - Challenge 3</span>
					<span>6:00pm - Challenge 4</span>
					<span>6:30pm - Challenge 5</span>
					<span>7:00pm - Challenge 6</span>
					<span>7:30pm - Challenge 7</span>
					<span>8:00pm - Challenge 8</span>
					<span>8:45pm - Competition end</span>
				</div>
			</Paragraph>
		{/if}
	</Page>
	<Footer />
</Wrapper>
