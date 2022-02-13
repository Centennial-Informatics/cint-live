<script lang="ts">
	import { currentPage } from '$lib/data/stores/currentPage';
	import { onMount } from 'svelte';

	onMount(() => {
    if (!$currentPage)
      currentPage.set('home')
  });
</script>
