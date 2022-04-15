<script lang="ts">
	export let small = false;
	export let disabled = false;
	export let onClick: () => void = () => {};
	export let className = '';
	export let outlined = false;
</script>

<button
	class="{className} {!small ? 'w-full h-20 text-xl' : 'h-10 text-sm'} {!outlined
		? 'bg-brand text-white hover:bg-brand-light'
		: 'text-brand-darkest bg-brand-lightest hover:bg-brand-light'} font-medium tracking-tight rounded-xl transition-colors disabled:bg-gray-200 disabled:hover:bg-gray-200 shadow-lg
      "
	on:click={onClick}
	{disabled}
>
	<slot />
</button>
