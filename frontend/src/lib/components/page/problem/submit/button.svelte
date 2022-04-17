<script lang="ts">
	import { brand } from '$lib/data/stores/userInfo';

	export let small = false;
	export let disabled = false;
	export let onClick: () => void = () => {};
	export let className = '';
	export let outlined = false;
</script>

<button
	class="{className} {!small ? 'w-full h-20 text-xl' : 'h-10 text-sm'} {!outlined
		? `bg-${$brand} text-white hover:bg-opacity-70`
		: `text-brand-darkest bg-${brand}-lightest hover:bg-opacity-70`} font-medium tracking-tight rounded-xl transition-colors disabled:bg-gray-400 disabled:hover:bg-gray-400 shadow-lg
      "
	on:click={onClick}
	{disabled}
>
	<slot />
</button>
