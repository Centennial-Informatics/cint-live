<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/layout/sidebar/sidebar.svelte';
	import Navbar from '$lib/components/layout/navbar/navbar.svelte';
	import { onMount, setContext } from 'svelte';
	import axios from 'axios';
	import type { ContestData } from 'src/types';
	import { contestData } from '$lib/data/stores/contestData';
	import { DefaultConfig } from '$lib/data/configs/default';

	setContext('contestConfig', DefaultConfig);

	$: console.log($contestData);
	onMount(async () => {
		if (!$contestData.Info)
			contestData.set((await axios.post('/api/v1/collect')).data as ContestData);
	});
</script>


<div class="min-h-screen">
	<div class="flex flex-col md:flex-row dark">
		<Sidebar />
		<div class="flex flex-row w-full bg-gray-800">
			<Navbar />
			<slot />
		</div>
	</div>
</div>
