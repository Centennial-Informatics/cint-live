<script lang="ts">
	import { brand } from '$lib/data/stores/userInfo';

	export let onClick: () => void;
	export let selected = false;
</script>

<button
	on:click={onClick}
	type="button"
	class="font-bold text-xl px-10 py-2 {selected
		? `bg-${$brand}-light text-gray-100`
		: 'hover:text-gray-200 text-gray-400'} rounded-lg box-border w-1/3 transition-color duration-200"
>
	<slot />
</button>
