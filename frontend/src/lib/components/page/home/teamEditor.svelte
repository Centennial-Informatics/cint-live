<script>
	import { TeamInfoData } from '$lib/data/stores/userInfo';
	import JoinTeam from './joinTeam.svelte';
	import Register from './register.svelte';
	import TeamView from './teamView.svelte';
</script>

<div class="py-10">
	{#if $TeamInfoData.team}
		<TeamView
			teamInfo={$TeamInfoData.team}
			teamMembers={$TeamInfoData.members}
			teamName={$TeamInfoData.team.name}
		/>
		<JoinTeam />
	{:else}
		<Register />
	{/if}
</div>
