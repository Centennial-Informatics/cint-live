<script lang="ts">
	import Centered from '$lib/components/templates/centered.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { startTime, stopTime } from '$lib/data/stores/contestData';
	import { currentTime } from '$lib/data/stores/currentTime';
	import fmtTime, { fmtLabel } from '$lib/utils/time';

	let label = '';
	let time = '';

	$: label = fmtLabel($startTime.valueOf(), $stopTime.valueOf(), $currentTime);
	$: time = fmtTime($startTime.valueOf(), $stopTime.valueOf(), $currentTime);
</script>

<div class="mt-1 mb-5">
	<Centered>
		<div class="p-2 flex items-center justify-center flex-col text-center">
			<Subtitle>
				<div class="text-xs text-gray-400 mb-1">{label + ' '}</div>
				<div
					class="flex flex-row m-0 text-base w-full text-center justify-center items-center text-gray-300"
				>
					{#if label !== 'COMPETITION HAS CONCLUDED'}
						{time}
					{/if}
				</div>
			</Subtitle>
		</div>
	</Centered>
</div>
