<script>
	import { IDToken, TeamInfoData } from '$lib/data/stores/userInfo';
	import JoinTeam from './joinTeam.svelte';
	import Register from './register.svelte';
	import TeamView from './teamView.svelte';
	import GetTeam from '$lib/utils/networking/team';

	$: if ($IDToken && !$TeamInfoData.members) {
		(async () => {
			TeamInfoData.set(await GetTeam($IDToken));
		})();
	}
</script>

<div class="py-10">
	{#if $TeamInfoData.team}
		<TeamView teamInfo={$TeamInfoData.team} teamMembers={$TeamInfoData.members} />
		<JoinTeam />
	{:else}
		<Register />
	{/if}
</div>
