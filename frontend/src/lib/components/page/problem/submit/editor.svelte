<script lang="ts">
	export let submissionText: string;

	function onChange(e: Event & { currentTarget: EventTarget & HTMLTextAreaElement }) {
		submissionText = e.currentTarget?.value;
	}
</script>

<div class="min-h-64 text-sm bg-gray-600 text-gray-200 font-mono">
	<textarea
		class="bg-transparent resize-none w-full min-h-64 h-full overflow-visible outline-none p-5 flex"
		placeholder="Paste your code here…"
		on:input={onChange}>{submissionText}</textarea
	>
</div>
