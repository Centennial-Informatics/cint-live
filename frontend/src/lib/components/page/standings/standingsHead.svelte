<script lang="ts">
	import { STANDARD } from '$lib/data/constants/division';

	export let division = STANDARD;
</script>

<div
	class="z-20 flex flex-row items-center overflow-x-visible bg-{division === STANDARD
		? 'brand'
		: 'alt'} text-white text-md rounded-xl px-5 h-16 font-medium space-x-3 whitespace-nowrap mb-1"
>
	<slot />
</div>
