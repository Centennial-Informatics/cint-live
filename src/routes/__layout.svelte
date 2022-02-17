<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/layout/sidebar/sidebar.svelte';
	import Navbar from '$lib/components/layout/navbar/navbar.svelte';
	import { onMount, setContext } from 'svelte';
	import { contestData } from '$lib/data/stores/contestData';
	import { DefaultConfig } from '$lib/data/configs/default';
	import Collect from '$lib/utils/networking/collect';

	setContext('contestConfig', DefaultConfig);

	$: console.log($contestData);
	onMount(async () => {
		if (!$contestData.Info) contestData.set(await Collect());
	});
</script>

<div class="min-h-screen">
	<div class="flex flex-col md:flex-row dark">
		<Sidebar />
		<div class="flex flex-row w-full bg-gray-800">
			<Navbar />
			<div class="md:overflow-y-scroll md:max-h-screen w-full dark:bg-gray-800">
				<slot />
			</div>
		</div>
	</div>
</div>
