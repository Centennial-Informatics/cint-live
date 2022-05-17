<script lang="ts">
	import SectionTitle from '$lib/components/templates/typography/sectionTitle.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { ADVANCED, STANDARD } from '$lib/data/constants/division';
	import { IDToken } from '$lib/data/stores/userInfo';
	import CreateUser from '$lib/utils/networking/admin/create';
	import type { Division } from 'src/types/teamInfo';
	import DivisionButton from '../home/divisionButton.svelte';

	let email = '';
	let name = '';
	let division = STANDARD;
	function setDivision(divisionChoice: Division) {
		division = divisionChoice;
	}
	function setDivisionChoice(divisionChoice: Division) {
		return () => setDivision(divisionChoice);
	}
	async function onSubmit() {
		if (email && name)
			// make sure not empty
			await CreateUser($IDToken, email, name, division);
		email = '';
		name = '';
		division = STANDARD;
	}
</script>

<div>
	<SectionTitle>Manually Register a User</SectionTitle>
	<Subtitle
		>If email is already registered, nothing will happen. After the contest begins, this is only way
		to add a new user.</Subtitle
	>
	<form on:submit|preventDefault={onSubmit} class="space-y-2 py-2">
		<input
			class="outline-none bg-gray-800 p-3 rounded-lg w-full text-gray-50 font-bold"
			placeholder="Email"
			bind:value={email}
		/>
		<input
			class="outline-none bg-gray-800 p-3 rounded-lg w-full text-gray-50 font-bold"
			placeholder="Name"
			bind:value={name}
		/>
		<div class="flex flex-row space-x-2">
			<DivisionButton
				division={STANDARD}
				onClick={setDivisionChoice(STANDARD)}
				active={division === STANDARD}
			/>
			<DivisionButton
				division={ADVANCED}
				onClick={setDivisionChoice(ADVANCED)}
				active={division === ADVANCED}
			/>
		</div>
		<button
			class="bg-gray-500 text-white font-bold text-xl w-full py-4 rounded-xl hover:opacity-70"
			type="submit">Submit</button
		>
	</form>
</div>
