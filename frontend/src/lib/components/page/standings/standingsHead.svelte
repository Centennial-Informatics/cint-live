<script>
	import { brand } from '$lib/data/stores/userInfo';
</script>

<div
	class="z-20 block overflow-x-auto bg-{$brand} text-white text-md rounded-xl p-5 font-medium space-x-3 whitespace-nowrap mb-1"
>
	<slot />
</div>
