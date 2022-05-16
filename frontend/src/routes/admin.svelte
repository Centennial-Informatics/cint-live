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

	let standingsData: StandingsData[];

	async function setAdminStandings() {
		standingsData = await AdminStandings($IDToken);
	}

	$: if ($IDToken) {
		setAdminStandings();
	}
</script>

<div class="h-full">
	<Wrapper>
		<Page>
			<Header>
				<Title>Admin Panel</Title>
				<Subtitle>This page will only work if you're logged into an Admin account.</Subtitle>
				<div class="flex flex-col space-y-10 py-10">
					<Download />
				</div>
			</Header>
		</Page>
	</Wrapper>
	{#if standingsData}
		<Standings standingsDataObj={standingsData} frozenOverride divisionOverride={STANDARD} />
		<Standings standingsDataObj={standingsData} frozenOverride divisionOverride={ADVANCED} />
	{/if}
</div>
