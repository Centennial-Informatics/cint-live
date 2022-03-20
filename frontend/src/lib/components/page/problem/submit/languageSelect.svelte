<script lang="ts">
	import type { ContestLanguage } from 'src/types/contestData';
	import LanguageButton from './languageButton.svelte';

	export let langs: ContestLanguage[] = [];
	export let selectedLanguage: string;

	function onChange(ID: string) {
		selectedLanguage = ID;
	}
</script>

<div class="flex-row p-2 w-full bg-gray-600">
	{#each langs as lang (lang.ID)}
		<LanguageButton selected={selectedLanguage === lang.ID} onClick={() => onChange(lang.ID)}>
			{lang.Name}
		</LanguageButton>
	{/each}
</div>
