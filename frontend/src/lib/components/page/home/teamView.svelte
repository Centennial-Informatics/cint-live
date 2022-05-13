<script lang="ts">
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { brand, IDToken, TeamInfoData } from '$lib/data/stores/userInfo';
	import LeaveTeam from '$lib/utils/networking/leaveTeam';
	import UpdateTeamName from '$lib/utils/networking/updateTeam';
	import type { TeamInfo, TeammateInfo } from 'src/types/teamInfo';
	import Button from '../problem/submit/button.svelte';
	import TeamCode from './teamCode.svelte';
	import Check from 'svelte-material-icons/Check.svelte';
	import UpdateTeamDivision from '$lib/utils/networking/updateDivision';
	import { contestData } from '$lib/data/stores/contestData';
	import CollectAdvanced from '$lib/utils/networking/collectAdvanced';
	import { ADVANCED, STANDARD } from '$lib/data/constants/division';
	import CollectStandard from '$lib/utils/networking/collectStandard';
	import DivisionSelect from './divisionSelect.svelte';

	export let teamInfo: TeamInfo;
	export let teamMembers: TeammateInfo[];
	export let teamName = $TeamInfoData.team.name;
	export let editing = false;
	export let editable = true;
	let error = '';

	// $: if ($TeamInfoData.team) {
	// 	teamName = $TeamInfoData.team.name;
	// }

	async function leaveTeam() {
		TeamInfoData.set(await LeaveTeam($IDToken));
	}

	async function updateTeamName(teamName: string) {
		error = await UpdateTeamName($IDToken, teamName);
		if (error === '') {
			const data = $TeamInfoData;
			data.team.name = teamName;
			TeamInfoData.set(data);
			editing = false;
		}
	}
	async function updateTeamDivision(division: string) {
		await UpdateTeamDivision($IDToken, division);
		const data = $TeamInfoData;
		data.team.division = division;
		TeamInfoData.set(data);
		if (division === ADVANCED) {
			(async () => {
				contestData.set(await CollectAdvanced());
			})();
		} else if (division === STANDARD) {
			(async () => {
				contestData.set(await CollectStandard());
			})();
		}
	}
</script>

<Wrapper margin={false} padding={false}>
	<div class="py-8 flex flex-col w-full px-8 items-center space-y-6">
		<TeamCode teamCode={teamInfo.code} />
		<div class="w-full space-y-2">
			<div>
				<Subtitle>Division Selection</Subtitle>
				<DivisionSelect {updateTeamDivision} />
			</div>
			<Subtitle>
				{#if error === ''}
					Team Name
				{:else}
					<span class="text-bad-50">{error}</span>
				{/if}
			</Subtitle>

			<div class="flex flex-row space-x-3 transition-all">
				{#if editable}
					<input
						class="outline-none bg-gray-800 p-3 rounded-lg w-full text-gray-50 font-bold"
						bind:value={teamName}
						on:keydown={(e) => {
							if (e.key === 'Enter') {
								updateTeamName(teamName);
							}
						}}
						on:focus={() => {
							editing = true;
						}}
					/>
					{#if editing}
						<div
							class="flex justify-center items-center bg-{$brand} m-auto p-1 rounded-lg hover:bg-opacity-70 cursor-pointer"
							on:click={() => {
								updateTeamName(teamName);
							}}
						>
							<Check width="40" height="40" color="#ffffff" />
						</div>
					{/if}
				{:else}
					<div class="w-full text-gray-50 font-bold text-2xl">
						{teamName}
					</div>
				{/if}
			</div>
		</div>
		<div class="w-full space-y-2">
			<Subtitle>Members</Subtitle>
			{#each teamMembers as member}
				<div class="text-gray-50 rounded-lg p-4 font-bold">
					{member.username}
				</div>
			{/each}
		</div>
		{#if teamMembers.length > 1}
			<Button onClick={leaveTeam} small className="px-16">Leave Team</Button>
		{/if}
	</div>
</Wrapper>
