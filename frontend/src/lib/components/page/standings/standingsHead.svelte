<script>
	import { brand } from '$lib/data/stores/userInfo';
</script>

<div
	class="z-20 flex flex-row items-center overflow-x-visible bg-{$brand} text-white text-md rounded-xl px-5 h-14 font-medium space-x-3 whitespace-nowrap mb-1"
>
	<slot />
</div>
