<script lang="ts">
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import Page from '$lib/components/templates/page/page.svelte';
	import Header from '$lib/components/templates/page/header.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Standings from './standings.svelte';
	import { IDToken, TeamInfoData } from '$lib/data/stores/userInfo';
	import type { StandingsData, StandingsEntry } from 'src/types/contestData';
	import AdminStandings from '$lib/utils/networking/admin/standings';
	import Download from '$lib/components/page/admin/download.svelte';
	import { ADVANCED, STANDARD } from '$lib/data/constants/division';
	import Create from '$lib/components/page/admin/create.svelte';
	import Delete from '$lib/components/page/admin/delete.svelte';
	import SectionTitle from '$lib/components/templates/typography/sectionTitle.svelte';

	let standingsData: StandingsData[];
	let valid = false;

	async function setAdminStandings() {
		standingsData = await AdminStandings($IDToken);
		valid = !!standingsData;
	}

	$: if ($IDToken) {
		setAdminStandings();
	}
</script>

<div>
	<Wrapper>
		<Page>
			<Header>
				<Title>Admin Panel</Title>
				<Subtitle>This page will only work if you're logged into an Admin account.</Subtitle>
				{#if valid}
					<div class="flex flex-col space-y-10 py-10">
						<Download />
						<Create />
						<Delete />
					</div>
				{/if}
			</Header>
      <button
			class="bg-gray-500 text-white font-bold text-xl w-full py-4 rounded-xl hover:opacity-70"
			on:click={setAdminStandings}>Refresh Standings</button
		>
		</Page>
	</Wrapper>
	{#if standingsData}
		<div class="h-screen">
			<Standings standingsDataObj={standingsData} frozenOverride divisionOverride={STANDARD} />
			<Standings standingsDataObj={standingsData} frozenOverride divisionOverride={ADVANCED} />
		</div>
	{/if}
</div>
