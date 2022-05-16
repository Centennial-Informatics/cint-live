<script lang="ts">
	import SectionTitle from '$lib/components/templates/typography/sectionTitle.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { IDToken } from '$lib/data/stores/userInfo';
	import DeleteUser from '$lib/utils/networking/admin/delete';

	let email = '';
	async function onSubmit() {
		if (email)
			// make sure not empty
			await DeleteUser($IDToken, email);
		email = '';
	}
</script>

<div>
	<SectionTitle>Delete/Ban a User</SectionTitle>
	<Subtitle>This is irreversible and not fully tested. Only use if absolutely necessary.</Subtitle>
	<form on:submit|preventDefault={onSubmit} class="space-y-2 py-2">
		<input
			class="outline-none bg-gray-800 p-3 rounded-lg w-full text-gray-50 font-bold"
			placeholder="Email"
			bind:value={email}
		/>
		<button
			class="bg-gray-500 text-white font-bold text-xl w-full py-4 rounded-xl hover:opacity-70"
			type="submit">Submit</button
		>
	</form>
</div>
