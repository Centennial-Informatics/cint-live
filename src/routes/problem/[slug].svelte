<script lang="ts">
	import { page } from '$app/stores';
	import { currentPage } from '$lib/data/stores/currentPage';
	import { onMount } from 'svelte';

	onMount(() => {
		if (!$currentPage) currentPage.set($page.params.slug);
	});
</script>

<div>{$page.params.slug}</div>
