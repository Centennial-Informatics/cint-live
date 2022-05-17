<script lang="ts">
	import SectionTitle from '$lib/components/templates/typography/sectionTitle.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { IDToken } from '$lib/data/stores/userInfo';
	import MakeAnnouncement from '$lib/utils/networking/admin/announce';

	let title = '';
	let details = '';
	async function onSubmit() {
		if (title) await MakeAnnouncement($IDToken, title, details);
		title = '';
		details = '';
	}
</script>

<div>
	<SectionTitle>Make an Announcement</SectionTitle>
	<Subtitle
		>Announcement is sent as a one-time alert to all participants with an active connection. It is
		not saved anywhere.</Subtitle
	>
	<form on:submit|preventDefault={onSubmit} class="space-y-2 py-2">
		<input
			class="outline-none bg-gray-800 p-3 rounded-lg w-full text-gray-50 font-bold"
			placeholder="Title"
			bind:value={title}
		/>
		<input
			class="outline-none bg-gray-800 p-3 rounded-lg w-full text-gray-50 font-bold"
			placeholder="Details"
			bind:value={details}
		/>
		<button
			class="bg-gray-500 text-white font-bold text-xl w-full py-4 rounded-xl hover:opacity-70"
			type="submit">Submit</button
		>
	</form>
</div>
