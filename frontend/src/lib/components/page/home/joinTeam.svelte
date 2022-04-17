<script lang="ts">
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { brand, IDToken, TeamInfoData } from '$lib/data/stores/userInfo';
	import JoinTeam from '$lib/utils/networking/joinTeam';
	import ChevronRight from 'svelte-material-icons/ChevronRight.svelte';

	let error = '';
	let teamCode = '';

	async function joinTeam(teamCode: string) {
		const data = await JoinTeam($IDToken, teamCode);
		if (data.error) {
			error = data.error;
		} else {
			TeamInfoData.set(data);
			error = '';
		}
	}
</script>

<Wrapper margin={false} padding={false}>
	<div class="p-6 flex flex-col w-full space-y-2">
		<Subtitle>
			{#if error}
				<span class="text-bad-50">{error}</span>
			{:else}
				Join a different team
			{/if}
		</Subtitle>
		<div class="flex flex-row space-x-3">
			<input
				class="outline-none bg-gray-800 p-3 rounded-lg w-full text-gray-50 font-bold"
				placeholder="Enter code"
				bind:value={teamCode}
				on:keydown={(e) => {
					if (e.key === 'Enter') {
						joinTeam(teamCode);
						teamCode = '';
					}
				}}
			/>
			<div
				class="flex justify-center items-center bg-{$brand} h-full rounded-lg hover:bg-opacity-70 cursor-pointer"
				on:click={() => {
					joinTeam(teamCode);
					teamCode = '';
				}}
			>
				<ChevronRight width="50" height="50" color="#ffffff" />
			</div>
		</div>
	</div>
</Wrapper>
