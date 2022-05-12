<!-- SSR Code -->
<script lang="ts" context="module">
	import CollectStandard from '$lib/utils/networking/collectStandard';
	import axios from 'axios';
	import type { ContestData } from 'src/types';
	import { browser } from '$app/env';
	export async function load() {
		const res = browser
			? await CollectStandard()
			: ((await axios.get(process.env.BASE_URL + '/api/v1/collect/standard')).data as ContestData);
		return {
			props: {
				contestDataObj: res
			}
		};
	}
</script>

<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/layout/sidebar/sidebar.svelte';
	import Navbar from '$lib/components/layout/navbar/navbar.svelte';
	import { onMount, setContext } from 'svelte';
	import {
		contestData,
		problemNames,
		problemPages,
		submissionData,
		submissionWS
	} from '$lib/data/stores/contestData';
	import { DefaultConfig } from '$lib/data/configs/default';
	import { IDToken, TeamInfoData } from '$lib/data/stores/userInfo';
	import NewWebSocket from '$lib/utils/networking/ws';
	import type { SubmissionVerdict } from 'src/types/contestData';
	import { convertVerdictsToMap } from '$lib/utils/verdictStatus';
	import GetTeam from '$lib/utils/networking/team';
	import CollectAdvanced from '$lib/utils/networking/collectAdvanced';
	import { ADVANCED } from '$lib/data/constants/division';
	import Loading from '$lib/components/templates/loading.svelte';

	export let contestDataObj: ContestData;
	contestData.set(contestDataObj);

	setContext('contestConfig', DefaultConfig);

	// onMount(async () => {
	// 	contestData.set(await CollectStandard());
	// });

	$: if ($IDToken && !$submissionWS) {
		submissionWS.set(
			NewWebSocket($IDToken, (data) => {
				submissionData.set(
					convertVerdictsToMap(JSON.parse(data) as SubmissionVerdict[], Object.keys($problemPages))
				);
			})
		);
	}

	$: if ($IDToken && !$TeamInfoData.members) {
		(async () => {
			TeamInfoData.set(await GetTeam($IDToken));
		})();
	}

	$: if ($TeamInfoData.team) {
		if ($TeamInfoData.team.division === ADVANCED) {
			(async () => {
				contestData.set(await CollectAdvanced());
			})();
		}
	}
</script>

<div class="min-h-screen">
	{#if $contestData.Info}
		<div class="flex flex-col md:flex-row dark">
			<div class="md:fixed">
				<Sidebar />
			</div>
			<div class="flex flex-row w-full bg-gray-800">
				<Navbar />
				<div class="md:min-h-screen md:ml-60 w-full dark:bg-gray-800">
					<slot />
				</div>
			</div>
		</div>
	{:else}
		<div class="flex justify-center items-center w-screen h-screen">
			<Loading />
		</div>
	{/if}
</div>
