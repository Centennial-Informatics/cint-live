<script lang="ts">
	import { brand } from '$lib/data/stores/userInfo';

	export let teamCode: string;
	let copied = false;
	let copyErr = '';

	async function copyCodeToClipboard() {
		try {
			await navigator.clipboard.writeText(teamCode);
			copied = true;
		} catch (err) {
			copyErr = 'Failed to copy: ' + err;
		}
	}
</script>

<div
	class="bg-{$brand} w-full p-6 rounded-xl flex flex-row text-md text-gray-50 items-center hover:bg-opacity-70 cursor-pointer"
	on:click={copyCodeToClipboard}
>
	Your Team Code
	{#if copied}
		<span class="ml-2 font-bold">Copied to clipboard!</span>
	{:else if copyErr}
		<span class="ml-2 font-bold">{copyErr}</span>
	{/if}
	<div class="flex-grow" />
	<div class="text-4xl font-bold">{teamCode.toUpperCase()}</div>
</div>
