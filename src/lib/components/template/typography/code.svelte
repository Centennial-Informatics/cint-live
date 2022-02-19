<script lang="ts">
	export let code: string = '';
</script>

<pre class="text-gray-200">{@html code.replace(/\n/g, '<br>')}</pre>
