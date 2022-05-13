<script lang="ts">
	import Code from '$lib/components/templates/typography/code.svelte';
	import { brand } from '$lib/data/stores/userInfo';

	export let input = '';
	export let output = '';

	let inputCopied = false;
	let outputCopied = false;

	async function copyInputToClipboard() {
		try {
			await navigator.clipboard.writeText(input);
			inputCopied = true;
		} catch (err) {
			console.log('Failed to copy input: ' + err);
		}
	}

	async function copyOutputToClipboard() {
		try {
			await navigator.clipboard.writeText(output);
			outputCopied = true;
		} catch (err) {
			console.log('Failed to copy output: ' + err);
		}
	}
</script>

<div class="flex flex-row w-full space-x-4 pt-1">
	<div
		class="w-1/2 text-gray-400 bg-gray-600 p-3 rounded-lg shadow-lg hover:bg-opacity-70 cursor-pointer"
		on:click={copyInputToClipboard}
	>
		<div class="flex flex-row justify-between">
			Input
			{#if inputCopied}
				<span class="font-bold text-{$brand}-light">Copied!</span>
			{/if}
		</div>
		<Code code={input} />
	</div>
	<div
		class="w-1/2 text-gray-400 bg-gray-600 p-3 rounded-lg shadow-lg hover:bg-opacity-70 cursor-pointer"
		on:click={copyOutputToClipboard}
	>
		<div class="flex flex-row justify-between">
			Output
			{#if outputCopied}
				<span class="font-bold text-{$brand}-light">Copied!</span>
			{/if}
		</div>
		<Code code={output} />
	</div>
</div>
