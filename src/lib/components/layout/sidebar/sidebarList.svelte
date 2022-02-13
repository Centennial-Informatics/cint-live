<script lang="ts">
	export let label: string = '';
</script>

<div class="flex flex-col overflow-y-auto">
	<div class="text-gray-400 text-xs mb-2 px-5">{label}</div>
	<div class="h-0 list-none flex-1 "><slot /></div>
</div>
