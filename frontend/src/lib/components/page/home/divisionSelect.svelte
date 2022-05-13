<script lang="ts">
	import { ADVANCED, STANDARD } from '$lib/data/constants/division';
	import { contestStarted } from '$lib/data/stores/currentTime';
	import { TeamInfoData } from '$lib/data/stores/userInfo';
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
	{#if !contestStarted || $TeamInfoData.team.division === STANDARD}
		<div
			class="w-full space-y-2 {$TeamInfoData.team.division === STANDARD
				? 'bg-brand'
				: ''} px-4 py-2 rounded-xl hover:bg-brand hover:text-brand-lightest cursor-pointer"
			on:click={updateStandard}
		>
			<span
				class="{$TeamInfoData.team.division === STANDARD
					? 'text-brand-lightest'
					: 'text-gray-50'} font-bold text-2xl">{STANDARD}</span
			>
		</div>
	{/if}
	{#if !contestStarted || $TeamInfoData.team.division === ADVANCED}
		<div
			class="w-full space-y-2 {$TeamInfoData.team.division === ADVANCED
				? 'bg-alt-dark'
				: ''} px-4 py-2 rounded-xl hover:bg-alt-dark hover:text-alt-lightest cursor-pointer"
			on:click={updateAdvanced}
		>
			<span
				class="{$TeamInfoData.team.division === ADVANCED
					? 'text-alt-lightest'
					: 'text-gray-50'} font-bold text-2xl">{ADVANCED}</span
			>
		</div>
	{/if}
</div>
