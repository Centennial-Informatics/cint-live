<script lang="ts">
	export let input = '';
	export let output = '';
</script>

<div class="flex flex-row w-full space-x-4 pt-1">
	<div class="w-1/2 text-gray-400 bg-gray-600 p-3 rounded-lg shadow-lg">
		Input
		<pre class="text-gray-200 text-lg">{input}</pre>
	</div>
	<div class="w-1/2 text-gray-400 bg-gray-600 p-3 rounded-lg shadow-lg">
		Output
		<pre class="text-gray-200 text-lg">{output}</pre>
	</div>
</div>
