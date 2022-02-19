<script lang="ts">
	import Code from '$lib/components/template/typography/code.svelte';

	export let input = '';
	export let output = '';
</script>

<div class="flex flex-row w-full space-x-4 pt-1">
	<div class="w-1/2 text-gray-400 bg-gray-600 p-3 rounded-lg shadow-lg">
		Input
		<Code code={input} />
	</div>
	<div class="w-1/2 text-gray-400 bg-gray-600 p-3 rounded-lg shadow-lg">
		Output
		<Code code={output} />
	</div>
</div>
