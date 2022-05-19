<script lang="ts">
	export let link: string = '/challenges';
	export let name: string = '';
	export let active: boolean = false;
</script>

{#if !active}
	<span class="text-white bg-gray-500 py-6 px-10 text-xl font-bold rounded-xl m-2">
		{name}
	</span>
{:else}
	<a
		class="text-white bg-brand cursor-pointer hover:opacity-70 py-6 px-10 text-xl font-bold rounded-xl m-2"
		href={link}
		target="_blank"
	>
		{name}
	</a>
{/if}
