<script lang="ts">
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import Page from '$lib/components/templates/page/page.svelte';
	import Header from '$lib/components/templates/page/header.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Standings from './standings.svelte';
	import { IDToken } from '$lib/data/stores/userInfo';
	import type { StandingsData } from 'src/types/contestData';
	import AdminStandings from '$lib/utils/networking/admin/standings';
	import Download from '$lib/components/page/admin/download.svelte';
	import { ADVANCED, STANDARD } from '$lib/data/constants/division';
	import Create from '$lib/components/page/admin/create.svelte';
	import Delete from '$lib/components/page/admin/delete.svelte';
	import Announce from '$lib/components/page/admin/announce.svelte';
	import UpdateCache from '$lib/utils/networking/admin/updateCache';
	import UpdateStart from '$lib/components/page/admin/updateStart.svelte';
	import UpdateEnd from '$lib/components/page/admin/updateEnd.svelte';

	let standingsData: StandingsData[];
	let valid = false;

	async function setAdminStandings() {
		standingsData = await AdminStandings($IDToken);
		valid = !!standingsData;
	}

	async function updateCache() {
		await UpdateCache($IDToken);
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
						<Announce />
						<Download />
						<Create />
						<Delete />
						<UpdateStart />
						<UpdateEnd />
					</div>
					<button
						class="bg-gray-500 text-white font-bold text-xl w-full py-4 rounded-xl hover:opacity-70"
						on:click={updateCache}>Update contest data</button
					>
					<button
						class="bg-gray-500 text-white font-bold text-xl w-full py-4 rounded-xl hover:opacity-70"
						on:click={setAdminStandings}>Refresh Standings</button
					>
				{/if}
			</Header>
		</Page>
	</Wrapper>
	{#if standingsData}
		<div class="h-screen">
			<Standings
				standingsDataObj={standingsData}
				frozenOverride
				divisionOverride={STANDARD}
				showFooter={false}
			/>
			<Standings
				standingsDataObj={standingsData}
				frozenOverride
				divisionOverride={ADVANCED}
				showFooter={false}
			/>
		</div>
	{/if}
</div>
