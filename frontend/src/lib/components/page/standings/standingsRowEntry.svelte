<script lang="ts">
	import { RankColors } from '$lib/data/constants/colors';

	import { problemNames } from '$lib/data/stores/contestData';
	import { fmtTimeHS } from '$lib/utils/time';
	import type { StandingsEntry } from 'src/types/contestData';
	import Trophy from 'svelte-material-icons/Trophy.svelte';
	import StandingsCell from './standingsCell.svelte';
	import StandingsRow from './standingsRow.svelte';
	export let rank = 0;
	export let standingsEntry: StandingsEntry;
	export let active = false;
	console.log(standingsEntry);
</script>

<StandingsRow {active}>
	<StandingsCell>
		{#if rank <= 3}
			<Trophy width={'32'} height={'32'} color={RankColors[rank]} />
		{:else}
			{rank}
		{/if}
	</StandingsCell>
	<StandingsCell className="flex-grow" left>
		{standingsEntry.Name}
	</StandingsCell>
	<div class="float-right">
		<div class="hidden lg:inline-block">
			{#each $problemNames as problem (problem.ID)}
				{#if problem.ID in standingsEntry.Submissions}
					<StandingsCell>
						<div class="relative">
							{#if standingsEntry.Points[problem.ID]}
								<span class={standingsEntry.Points[problem.ID] > 0 ? 'text-good-50 font-bold' : ''}
									>{standingsEntry.Points[problem.ID]}</span
								>
								<span
									class="absolute top-5 left-0 right-0 text-xs {active
										? 'text-gray-300'
										: 'text-gray-400'}"
									>{fmtTimeHS(standingsEntry.Submissions[problem.ID].time)}</span
								>
							{/if}
						</div>
					</StandingsCell>
				{/if}
			{/each}
		</div>
		<StandingsCell>
			{standingsEntry.TotalPoints}
		</StandingsCell>
	</div>
</StandingsRow>
