<script lang="ts">
	import { problemNames } from '$lib/data/stores/contestData';
	import type { ContestConfig } from 'src/types';
	import { getContext } from 'svelte';
	import Logo from './logo/logo.svelte';
import Time from './time/time.svelte';

	let config: ContestConfig = getContext('contestConfig');
</script>

<div class="h-screen bg-gray-800 flex flex-col md:flex-shrink-0 md:w-60">
	<Logo imgUrl={config.infoUrl || '/'} imgSrc={config.logoUrl} />
	<div class="flex flex-col h-0 flex-grow pt-5 md:space-y-10">
		<div class="px-5">
			<Time />
		</div>
	</div>
	{#each $problemNames as problem}
		<a class="text-white" href="/problem/{problem.ID}">{problem.Name}</a>
	{/each}
</div>
