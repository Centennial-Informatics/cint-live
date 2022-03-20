<script lang="ts" context="module">
	import axios from 'axios';
	import Standings from '$lib/utils/networking/standings';
	import { browser } from '$app/env';
	export async function load() {
		const res = browser
			? await Standings()
			: ((await axios.get('http://localhost:8000/api/v1/standings')).data as StandingsData);
		return {
			props: {
				standingsDataObj: res
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
	import { problemNames } from '$lib/data/stores/contestData';
	import Page from '$lib/components/templates/page/page.svelte';
	import StandingsRowEntry from '$lib/components/page/standings/standingsRowEntry.svelte';
	import { userInfo } from '$lib/data/stores/userInfo';
	import StandingsPagination from '$lib/components/page/standings/standingsPagination.svelte';

	export let standingsDataObj: StandingsData;

	onMount(() => {
		if (!$currentPage) currentPage.set('standings');
	});

	function cleanData(standingsData: StandingsData): StandingsEntry[] {
		let entries: StandingsEntry[] = [];
		/* Sum up points */
		Object.keys(standingsData).forEach((userID) => {
			const entry = standingsData[userID];
			entry.TotalPoints = 0;
			entry.ID = userID;
			Object.keys(standingsData[userID].Points).forEach((c: string) => {
				entry.TotalPoints += entry.Points[c];
			});
			entries.push(entry);
		});

		/* Sort by total points and most recent submission */
		entries.sort((a, b) => {
			if (a.TotalPoints == b.TotalPoints) {
				let aMostRecent = 0;
				let bMostRecent = 0;
				Object.keys(a.Submissions).forEach((probID) => {
					aMostRecent = Math.max(aMostRecent, a.Submissions[probID].Time);
					bMostRecent = Math.max(bMostRecent, b.Submissions[probID].Time);
				});
				return aMostRecent < bMostRecent ? -1 : 1;
			} else {
				return a.TotalPoints > b.TotalPoints ? -1 : 1;
			}
		});

		return entries;
	}

	let standingsEntries: StandingsEntry[];
	$: standingsEntries = cleanData(standingsDataObj);

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
						active={entry.ID === $userInfo.ID}
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
