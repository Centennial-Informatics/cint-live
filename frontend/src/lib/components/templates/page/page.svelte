<script lang="ts">
	export let visible = true;
</script>

<div class="py-15 px-6 sm:px-12 md:px-14 w-full {!visible ? 'hidden' : ''}"><slot /></div>
