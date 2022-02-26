<script lang="ts">
	export let transparent = false;
</script>

<div
	class="flex flex-col items-center pb-24 pt-2 min-h-full {transparent
		? 'dark:bg-transparent mt-20 overflow-y-auto relative'
		: 'dark:bg-gray-700 md:shadow-xl md:mx-6 mt-24 relative'} mb-10 rounded-t-3xl md:rounded-3xl"
>
	<slot />
</div>
