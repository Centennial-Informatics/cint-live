<script lang="ts">
	import ChallengeLink from './challengeLink.svelte';

	export let links: string[];
</script>

<div class="flex flex-row flex-wrap py-4 items-center">
	{#each links as link, i}
		<ChallengeLink active={i === links.length - 1} {link} name={'Challenge ' + (i + 1)} />
	{/each}
</div>
