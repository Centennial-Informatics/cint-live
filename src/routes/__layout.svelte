<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/layout/sidebar/sidebar.svelte';
	import Navbar from '$lib/components/layout/navbar.svelte';
	import { onMount, setContext } from 'svelte';
	import axios from 'axios';
	import type { ContestData } from 'src/types';
	import { contestData } from '$lib/data/stores/contestData';
	import { DefaultConfig } from '$lib/data/configs/default';

	setContext('contestConfig', DefaultConfig);

	console.log($contestData);
	onMount(async () => {
		if (!$contestData.Info)
			contestData.set((await axios.post('/api/v1/collect')).data as ContestData);
	});
</script>

<svelte:head>
	<link
		href="https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;700&display=swap"
		rel="stylesheet"
	/>
	<link
		href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;700;900&display=swap"
		rel="stylesheet"
	/>
	<link href="https://fonts.googleapis.com/css?family=Fira+Mono" rel="stylesheet" />
	<link
		href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@300;400;600;700;800&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<div class="min-h-screen">
	<div class="flex flex-col md:flex-row dark">
		<Sidebar />
		<div class="flex flex-row w-full">
			<Navbar />
			<slot />
		</div>
	</div>
</div>
