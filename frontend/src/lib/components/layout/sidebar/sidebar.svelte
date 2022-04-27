<script lang="ts">
	import { problemNames, submissionData } from '$lib/data/stores/contestData';
	import type { ContestConfig } from 'src/types';
	import { getContext } from 'svelte';
	import Logo from './logo/logo.svelte';
	import SidebarList from './sidebarList.svelte';
	import SidebarLink from './sidebarLink.svelte';
	import Time from './time/time.svelte';
	import verdictStatus from '$lib/utils/verdictStatus';

	let config: ContestConfig = getContext('contestConfig');
</script>

<div class="md:h-screen bg-gray-800 flex flex-col md:flex-shrink-0 md:w-60 min-h-screen">
	<Logo imgUrl={config.infoUrl || '/'} imgSrc={config.logoUrl} />
	<div class="flex flex-col h-0 flex-grow pt-5 md:space-y-10">
		<div class="px-5">
			<Time />
		</div>
		<div class="pb-10 pt-2 md:hidden">
			<SidebarList label="INFO">
				<SidebarLink to="/" tabId="home" status="info">Home</SidebarLink>
				<SidebarLink to="/standings" tabId="standings" status="info">Scoreboard</SidebarLink>
			</SidebarList>
		</div>
		<SidebarList label="PROBLEMS — {$problemNames.length}">
			{#each $problemNames as problem}
				<SidebarLink
					to="/problem/{problem.ID}"
					tabId={problem.ID}
					status={problem.ID in $submissionData ? verdictStatus($submissionData[problem.ID]) : ''}
					>{problem.Name}</SidebarLink
				>
			{/each}
		</SidebarList>
	</div>
</div>
