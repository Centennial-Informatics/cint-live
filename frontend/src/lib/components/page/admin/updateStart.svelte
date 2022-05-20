<script lang="ts">
	import SectionTitle from '$lib/components/templates/typography/sectionTitle.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { IDToken } from '$lib/data/stores/userInfo';
	import UpdateStart from '$lib/utils/networking/admin/updatestart';

	let time = '';
	async function onSubmit() {
		if (time)
			// make sure not empty
			await UpdateStart($IDToken, time);
		time = '';
	}
</script>

<div>
	<SectionTitle>Update Start Time</SectionTitle>
	<Subtitle>2006-01-02T15:04:05Z07:00 Format</Subtitle>
	<form on:submit|preventDefault={onSubmit} class="space-y-2 py-2">
		<input
			class="outline-none bg-gray-800 p-3 rounded-lg w-full text-gray-50 font-bold"
			placeholder="Email"
			bind:value={time}
		/>
		<button
			class="bg-gray-500 text-white font-bold text-xl w-full py-4 rounded-xl hover:opacity-70"
			type="submit">Update</button
		>
	</form>
</div>
