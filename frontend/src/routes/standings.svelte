<script lang="ts" context="module">
	import axios from 'axios';
	import Standings from '$lib/utils/networking/standings';
	import { browser } from '$app/env';
	export async function load() {
		const res = browser
			? await Standings()
			: ((await axios.get(process.env.HOST + '/api/v1/standings')).data as StandingsData[]);
		return {
			props: {
				standingsData: res
			}
		};
	}
</script>

<script lang="ts">
	import { currentPage } from '$lib/data/stores/currentPage';
	import { onMount } from 'svelte';
	import type { StandingsData, StandingsEntry } from 'src/types/contestData';
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import StandingsHeader from '$lib/components/page/standings/standingsHeader.svelte';
	import StandingsTable from '$lib/components/page/standings/standingsTable.svelte';
	import StandingsHead from '$lib/components/page/standings/standingsHead.svelte';
	import StandingsCell from '$lib/components/page/standings/standingsCell.svelte';
	import { problemNames, problemPages } from '$lib/data/stores/contestData';
	import Page from '$lib/components/templates/page/page.svelte';
	import StandingsRowEntry from '$lib/components/page/standings/standingsRowEntry.svelte';
	import { TeamInfoData, userInfo } from '$lib/data/stores/userInfo';
	import StandingsPagination from '$lib/components/page/standings/standingsPagination.svelte';
	import { fillEmptyVerdicts } from '$lib/utils/verdictStatus';
	import { STANDARD } from '$lib/data/constants/division';

	export let standingsData: StandingsData[];

	onMount(() => {
		if (!$currentPage) currentPage.set('standings');
	});

	function cleanData(standingsData: StandingsData[]): StandingsEntry[] {
		standingsData = standingsData.filter((entry) => {
			if ($TeamInfoData.team) return entry.division === $TeamInfoData.team.division;
			else return entry.division === STANDARD;
		});
		const entries = standingsData.map((entry) => {
			let res: StandingsEntry = {
				Name: entry.team_name,
				Submissions: fillEmptyVerdicts(entry.verdicts, Object.keys($problemPages)),
				TotalPoints: 0,
				Points: {},
				ID: entry.team_id
			};
			Object.keys(res.Submissions).forEach((problemID) => {
				res.Points[problemID] = res.Submissions[problemID].points;
				res.TotalPoints += res.Points[problemID];
			});
			return res;
		});

		/* Sort by total points and most recent submission */
		entries.sort((a, b) => {
			if (a.TotalPoints == b.TotalPoints) {
				let aMostRecent = 0;
				let bMostRecent = 0;
				Object.keys(a.Submissions).forEach((probID) => {
					aMostRecent = Math.max(aMostRecent, a.Submissions[probID].time);
					bMostRecent = Math.max(bMostRecent, b.Submissions[probID].time);
				});
				return aMostRecent < bMostRecent ? -1 : 1;
			} else {
				return a.TotalPoints > b.TotalPoints ? -1 : 1;
			}
		});

		return entries;
	}

	let standingsEntries: StandingsEntry[];

	// run it once per page but re-run after login/logout
	$: if ($TeamInfoData.team) {
		standingsEntries = cleanData(standingsData);
	}
	$: if (!$TeamInfoData.team) {
		standingsEntries = cleanData(standingsData);
	}

	/* Table pagination */

	const entriesPerPage = 10;
	let page = 0;

	let displayedEntries: StandingsEntry[] = [];

	function displayPage(pageStart: number, pageEnd: number) {
		page = pageStart;
		displayedEntries = standingsEntries.slice(
			Math.max(0, pageStart),
			Math.min(standingsEntries.length, pageEnd)
		);
	}
</script>

<Wrapper transparent>
	<Page>
		<StandingsHeader />
		<div class="overflow-x-auto">
			<StandingsTable>
				<StandingsHead>
					<StandingsCell>Rank</StandingsCell>
					<StandingsCell className="pl-10 w-72 flex-grow" left>Name</StandingsCell>
					<div class="float-right whitespace-nowrap">
						<div class="hidden lg:inline-block">
							{#each $problemNames as problem (problem.ID)}
								<StandingsCell>{problem.ID}</StandingsCell>
							{/each}
						</div>
						<StandingsCell>Total</StandingsCell>
					</div>
				</StandingsHead>
				{#each displayedEntries as entry, i (entry.ID)}
					<StandingsRowEntry
						standingsEntry={entry}
						active={entry.ID === ($TeamInfoData.team ? $TeamInfoData.team.ID : 0)}
						rank={i + 1 + page}
					/>
				{/each}
			</StandingsTable>
		</div>
		<StandingsPagination
			itemsPerPage={entriesPerPage}
			totalItems={standingsEntries.length}
			{displayPage}
		/>
	</Page>
</Wrapper>
