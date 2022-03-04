<script lang="ts">
	import { RankColors } from '$lib/data/constants/colors';

	import { problemNames } from '$lib/data/stores/contestData';
	import type { StandingsEntry } from 'src/types/contestData';
	import Trophy from 'svelte-material-icons/Trophy.svelte';
	import StandingsCell from './standingsCell.svelte';
	import StandingsRow from './standingsRow.svelte';
	export let rank = 0;
	export let standingsEntry: StandingsEntry;
	export let active = false;
</script>

<StandingsRow {active}>
	<StandingsCell>
		{#if rank <= 3}
			<Trophy width={'32'} height={'32'} color={RankColors[rank]} />
		{:else}
			{rank}
		{/if}
	</StandingsCell>
	<StandingsCell className="pl-10 w-72 flex-grow" left>
		{standingsEntry.Username}
	</StandingsCell>
	<div class="float-right">
		<div class="hidden md:inline-block">
			{#each $problemNames as problem (problem.ID)}
				<StandingsCell>
					{standingsEntry.Points[problem.ID]}
				</StandingsCell>
			{/each}
		</div>
		<StandingsCell>
			{standingsEntry.TotalPoints}
		</StandingsCell>
	</div>
</StandingsRow>
