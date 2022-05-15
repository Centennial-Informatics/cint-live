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
	<div class="float-right relative">
		<div class="hidden lg:inline-block">
			{#each $problemNames as problem (problem.ID)}
				{#if problem.ID in standingsEntry.Submissions}
					<StandingsCell>
						<div class="flex flex-col items-center">
							{#if standingsEntry.Points[problem.ID] > 0}
								<span class={standingsEntry.Points[problem.ID] > 0 ? 'text-good-50 font-bold' : ''}
									>{standingsEntry.Points[problem.ID]}</span
								>
								<span class="text-xs {active ? 'text-gray-300' : 'text-gray-400'}"
									>{fmtTimeHS(standingsEntry.Submissions[problem.ID].time)}</span
								>
							{/if}
						</div>
					</StandingsCell>
				{/if}
			{/each}
		</div>
		<div class="inline-block w-12 ml-2 h-full">
			<div class="absolute right-0 {standingsEntry.TotalPoints > 0 ? 'top-2' : 'top-0'}">
				<StandingsCell className="">
					{standingsEntry.TotalPoints}
				</StandingsCell>
			</div>
		</div>
	</div>
</StandingsRow>
