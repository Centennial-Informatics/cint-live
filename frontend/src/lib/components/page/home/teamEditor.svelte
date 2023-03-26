<script>
	import { contestStarted } from '$lib/data/stores/currentTime';
	import { TeamInfoData } from '$lib/data/stores/userInfo';
	import JoinTeam from './joinTeam.svelte';
	import LogOut from './logOut.svelte';
	import Register from './register.svelte';
	import TeamView from './teamView.svelte';
</script>

<div class="py-10">
	{#if $TeamInfoData.team}
		<TeamView
			teamInfo={$TeamInfoData.team}
			teamMembers={$TeamInfoData.members}
			teamName={$TeamInfoData.team.name}
			editable={!$contestStarted}
		/>
		{#if !$contestStarted}
			<JoinTeam />
		{/if}
		<LogOut />
	{:else}
		<Register />
	{/if}
</div>
