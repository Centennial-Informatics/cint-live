<!-- <script lang="ts" context="module">
	// import axios from 'axios';
	import Standings from '$lib/utils/networking/standings';
	import { browser, mode } from '$app/env';
	export async function load() {
		const res = browser
			? await Standings()
			: await fetchType<StandingsData[]>(
					BASE_URL[mode as 'production' | 'development'] + '/api/v1/standings'
			  );
		return {
			props: {
				standingsDataObj: res
			}
		};
	}
</script> -->
<script lang="ts">
	import Standings from '$lib/utils/networking/standings';
	import { currentPage } from '$lib/data/stores/currentPage';
	import { onMount } from 'svelte';
	import type { StandingsData, StandingsEntry } from 'src/types/contestData';
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import StandingsHeader from '$lib/components/page/standings/standingsHeader.svelte';
	import StandingsTable from '$lib/components/page/standings/standingsTable.svelte';
	import StandingsHead from '$lib/components/page/standings/standingsHead.svelte';
	import StandingsCell from '$lib/components/page/standings/standingsCell.svelte';
	import {
		problemNames,
		problemPages,
		standingsData,
		stopTime,
		submissionData
	} from '$lib/data/stores/contestData';
	import Page from '$lib/components/templates/page/page.svelte';
	import StandingsRowEntry from '$lib/components/page/standings/standingsRowEntry.svelte';
	import { TeamInfoData } from '$lib/data/stores/userInfo';
	import StandingsPagination from '$lib/components/page/standings/standingsPagination.svelte';
	import { fillEmptyVerdicts } from '$lib/utils/verdictStatus';
	import { ADVANCED, STANDARD } from '$lib/data/constants/division';
	import { contestEnded, currentTime } from '$lib/data/stores/currentTime';

	export let standingsDataObj: StandingsData[];
	export let frozenOverride = false;
	export let divisionOverride = '';

	$: if (standingsDataObj) standingsData.set(standingsDataObj);

	let frozen = false;

	async function fetchStandings() {
		if (!standingsDataObj) standingsData.set(await Standings());
	}

	function updateFrozen(currentTime: number) {
		if (!frozenOverride) {
			const sinceEnd = Math.abs($stopTime.valueOf() / 1000 - Math.floor(currentTime / 1000));
			if ($contestEnded && sinceEnd < 1800) {
				frozen = true;
			} else if (!$contestEnded && sinceEnd <= 1200) {
				frozen = true;
			} else {
				frozen = false;
			}

			if (!$contestEnded && sinceEnd == 1200) {
				alert('Scoreboard is now frozen. Results will not update during the final 20 minutes.');
			} else if ($contestEnded && sinceEnd == 1800) {
				alert('Final results are now available.');
			}
		}
	}

	$: updateFrozen($currentTime);

	onMount(async () => {
		if (!$currentPage) currentPage.set('standings');
		fetchStandings();
	});

	function cleanData(standingsData: StandingsData[], problems: string[]): StandingsEntry[] {
		const newData = standingsData.filter((entry) => {
			if (divisionOverride === STANDARD) return entry.division === divisionOverride;
			else if (divisionOverride === ADVANCED) return entry.division === divisionOverride;
			else if ($TeamInfoData.team) return entry.division === $TeamInfoData.team.division;
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
			// only update the user's submissions when standings are frozen
			if (
				Object.keys($submissionData).length === problems.length &&
				$TeamInfoData.team &&
				res.ID === $TeamInfoData.team.ID
			) {
				res.Submissions = fillEmptyVerdicts($submissionData, problems);
			}
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

	$: if ($standingsData) {
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

<Wrapper transparent fullHeight>
	<Page>
		<div class="relative">
			<StandingsHeader
				{frozen}
				division={divisionOverride
					? divisionOverride
					: $TeamInfoData.team
					? $TeamInfoData.team.division
					: STANDARD}
			/>
			<div class="absolute left-0 right-0">
				<div class="overflow-x-scroll">
					<StandingsTable>
						<StandingsHead
							division={divisionOverride
								? divisionOverride
								: $TeamInfoData.team
								? $TeamInfoData.team.division
								: STANDARD}
						>
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
								{frozen}
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
