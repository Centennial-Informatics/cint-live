<script lang="ts" context="module">
	// import axios from 'axios';
	import Standings from '$lib/utils/networking/standings';
	import { browser, mode } from '$app/env';
	export async function load() {
		const res = browser
			? await Standings()
			: await fetchType(BASE_URL[mode as 'production' | 'development'] + '/api/v1/standings');
		return {
			props: {
				standingsDataObj: res
			}
		};
	}
</script>

<script lang="ts">
	// import Standings from '$lib/utils/networking/standings';
	import { currentPage } from '$lib/data/stores/currentPage';
	import { onMount } from 'svelte';
	import type { StandingsData, StandingsEntry } from 'src/types/contestData';
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import StandingsHeader from '$lib/components/page/standings/standingsHeader.svelte';
	import StandingsTable from '$lib/components/page/standings/standingsTable.svelte';
	import StandingsHead from '$lib/components/page/standings/standingsHead.svelte';
	import StandingsCell from '$lib/components/page/standings/standingsCell.svelte';
	import { problemNames, problemPages, standingsData } from '$lib/data/stores/contestData';
	import Page from '$lib/components/templates/page/page.svelte';
	import StandingsRowEntry from '$lib/components/page/standings/standingsRowEntry.svelte';
	import { TeamInfoData } from '$lib/data/stores/userInfo';
	import StandingsPagination from '$lib/components/page/standings/standingsPagination.svelte';
	import { fillEmptyVerdicts } from '$lib/utils/verdictStatus';
	import { STANDARD } from '$lib/data/constants/division';
	import { BASE_URL } from '$lib/data/constants/url';
	import fetchType from '$lib/utils/networking/serverFetch';

	export let standingsDataObj: StandingsData[];
	standingsData.set(standingsDataObj);

	// async function fetchStandings() {
	// 	standingsData.set(await Standings());
	// }

	onMount(async () => {
		if (!$currentPage) currentPage.set('standings');
		// fetchStandings();
	});

	function cleanData(standingsData: StandingsData[], problems: string[]): StandingsEntry[] {
		const newData = standingsData.filter((entry) => {
			if ($TeamInfoData.team) return entry.division === $TeamInfoData.team.division;
			else return entry.division === STANDARD;
		});
		const entries = newData.map((entry) => {
			let res: StandingsEntry = {
				Name: entry.team_name,
				Submissions: fillEmptyVerdicts(entry.verdicts, problems),
				TotalPoints: 0,
				Points: {},
				ID: entry.team_id
			};
			Object.keys(res.Submissions).forEach((problemID) => {
				// filter out sample problems
				if (problemID in $problemPages) {
					res.Points[problemID] = res.Submissions[problemID].points;
					res.TotalPoints += res.Points[problemID];
				}
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

	let standingsEntries: StandingsEntry[] = cleanData($standingsData, Object.keys($problemPages));

	// run it once per page but re-run after login/logout
	$: if ($TeamInfoData.team && $standingsData) {
		standingsEntries = cleanData($standingsData, Object.keys($problemPages));
	}
	$: if ($TeamInfoData.team && !$standingsData) {
		standingsEntries = cleanData($standingsData, Object.keys($problemPages));
	}

	/* Table pagination */

	const entriesPerPage = 10;
	let page = 0;

	let displayedEntries: StandingsEntry[] = [];

	$: if (standingsEntries) {
		displayPage(page, page + entriesPerPage);
	}

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
		<div class="relative">
			<StandingsHeader />
			<div class="absolute left-0 right-0">
				<div class="overflow-x-scroll">
					<StandingsTable>
						<StandingsHead>
							<StandingsCell>Rank</StandingsCell>
							<StandingsCell left className="flex-grow">Name</StandingsCell>
							<div class="hidden lg:inline-block float-right">
								{#each $problemNames as problem (problem.ID)}
									<StandingsCell>{problem.ID}</StandingsCell>
								{/each}
							</div>
							<StandingsCell>Total</StandingsCell>
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
			</div>
		</div>
	</Page>
</Wrapper>
