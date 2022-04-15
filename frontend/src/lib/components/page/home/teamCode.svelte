<script lang="ts">
	import { brand } from '$lib/data/stores/userInfo';

	export let teamCode: string;
</script>

<div
	class="bg-{$brand} w-full p-6 rounded-xl flex flex-row text-md text-gray-50 items-center hover:bg-{$brand}-light cursor-pointer"
>
	Your Team Code
	<div class="flex-grow" />
	<div class="text-4xl font-bold">{teamCode.toUpperCase()}</div>
</div>
