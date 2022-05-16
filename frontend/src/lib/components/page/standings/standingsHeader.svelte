<script>
	import Header from '$lib/components/templates/page/header.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import { brand, TeamInfoData } from '$lib/data/stores/userInfo';

	export let frozen = false;
</script>

<Header>
	{#if $TeamInfoData.team}
		<div class="text-lg">
			<span class="text-{$brand}-light">{$TeamInfoData.team.division} Division</span>
		</div>
	{:else}
		<div class="text-lg"><span class="text-brand-light">Standard Division</span></div>
	{/if}
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
