<script lang="ts">
	export let active = false;
</script>

<div
	class="z-20 dark:text-gray-100 whitespace-nowrap h-12 px-5 text-sm space-x-3 flex flex-row items-center rounded-lg shadow-md my-1 {active
		? 'bg-gray-100 dark:bg-gray-500 font-semibold'
		: 'bg-white dark:bg-gray-700 font-medium'}"
>
	<slot />
</div>
