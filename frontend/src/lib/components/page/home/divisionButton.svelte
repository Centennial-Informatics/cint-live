<script lang="ts">
	import { STANDARD } from '$lib/data/constants/division';

	export let onClick: () => void;
	export let active = false;
	export let division = STANDARD;
	const color = division === STANDARD ? 'brand' : 'alt';
</script>

<div
	class="w-full space-y-2 {active
		? `bg-${color}`
		: ''} px-4 py-2 rounded-xl hover:bg-{color} cursor-pointer"
	on:click={onClick}
>
	<span class="{active ? `text-${color}-lightest` : 'text-gray-50'} font-bold text-2xl"
		>{division}</span
	>
</div>
