<script lang="ts">
	import { ADVANCED, STANDARD } from '$lib/data/constants/division';
	import { contestStarted } from '$lib/data/stores/currentTime';
	import { TeamInfoData } from '$lib/data/stores/userInfo';
	import DivisionButton from './divisionButton.svelte';
	export let updateTeamDivision: (division: string) => void;
	function updateStandard() {
		if ($TeamInfoData.team.division !== STANDARD) {
			updateTeamDivision(STANDARD);
		}
	}
	function updateAdvanced() {
		if ($TeamInfoData.team.division !== ADVANCED) {
			updateTeamDivision(ADVANCED);
		}
	}
</script>

<div class="flex flex-row justify-between space-x-4 mt-2 mb-6">
	{#if !$contestStarted || $TeamInfoData.team.division === STANDARD}
		<DivisionButton
			active={$TeamInfoData.team.division === STANDARD}
			onClick={updateStandard}
			division={$TeamInfoData.team.division}
		/>
	{/if}
	{#if !$contestStarted || $TeamInfoData.team.division === ADVANCED}
		<DivisionButton
			active={$TeamInfoData.team.division === ADVANCED}
			onClick={updateAdvanced}
			division={$TeamInfoData.team.division}
		/>
	{/if}
</div>
