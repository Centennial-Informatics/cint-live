<script lang="ts">
	import Header from '$lib/components/templates/page/header.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import { STANDARD } from '$lib/data/constants/division';

	export let frozen = false;
	export let division: string;
</script>

<Header>
	<div class="text-lg">
		{#if division === STANDARD}
			<span class="text-brand-light">{division} Division</span>
		{:else}
			<span class="text-alt-light">{division} Division</span>
		{/if}
	</div>
	<Title>Scoreboard</Title>
	<Subtitle>
		{#if frozen}
			Scoreboard is frozen. Final results will be revealed at the closing ceremony in the
			Auditorium.
		{:else}
			Refresh page to update. Tiebreaker is the time of the most recent submission.
		{/if}
	</Subtitle>
</Header>
