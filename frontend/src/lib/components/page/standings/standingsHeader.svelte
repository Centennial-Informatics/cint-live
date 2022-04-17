<script>
	import Header from '$lib/components/templates/page/header.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import { brand, TeamInfoData } from '$lib/data/stores/userInfo';
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
	<Subtitle>Refresh page to update.</Subtitle>
</Header>
