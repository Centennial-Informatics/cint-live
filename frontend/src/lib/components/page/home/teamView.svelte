<script lang="ts">
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { IDToken, TeamInfoData } from '$lib/data/stores/userInfo';
	import LeaveTeam from '$lib/utils/networking/leaveTeam';
	import type { TeamInfo, TeammateInfo } from 'src/types/teamInfo';
	import Button from '../problem/submit/button.svelte';
	import TeamCode from './teamCode.svelte';

	export let teamInfo: TeamInfo;
	export let teamMembers: TeammateInfo[];

	async function leaveTeam() {
		TeamInfoData.set(await LeaveTeam($IDToken));
	}
</script>

<Wrapper margin={false} padding={false}>
	<div class="py-8 flex flex-col w-full px-8 items-center space-y-6">
		<TeamCode teamCode={teamInfo.code} />
		<div class="w-full space-y-2">
			<Subtitle>Team Name</Subtitle>
			<div class="bg-gray-800 text-gray-50 rounded-lg p-4">{teamInfo.name}</div>
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
