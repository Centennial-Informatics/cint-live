<script lang="ts">
	export let transparent = false;
	export let margin = true;
	export let padding = true;
</script>

<div
	class="flex flex-col items-center min-h-full {transparent
		? 'dark:bg-transparent mt-20 overflow-y-auto relative'
		: 'dark:bg-gray-700 md:shadow-xl relative'} {margin ? ' md:mx-6 mt-24 ' : ''} 
		{padding ? ' pt-2 pb-24 ' : ''} mb-10 rounded-t-3xl md:rounded-3xl"
>
	<slot />
</div>
