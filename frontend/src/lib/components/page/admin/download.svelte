<script lang="ts">
	import SectionTitle from '$lib/components/templates/typography/sectionTitle.svelte';
import { ADVANCED, STANDARD } from '$lib/data/constants/division';

	import { IDToken } from '$lib/data/stores/userInfo';

	import DownloadStandings from '$lib/utils/networking/admin/download';

	async function onClick(division: string) {
		const res = await DownloadStandings($IDToken, division);
		const url = window.URL.createObjectURL(new Blob([res]));
		window.open(url, '_blank');
	}

	function onClickDivision(division: string) {
		return () => onClick(division);
	}
</script>

<div class="flex flex-col space-y-4 w-full">
  <SectionTitle>Download Standings as CSV</SectionTitle>
  <div class="flex flex-row space-x-4 w-full">
    <button class="bg-brand hover:opacity-70 py-2 px-4 font-bold text-brand-lightest rounded-xl text-lg" on:click={onClickDivision(STANDARD)}>Download Standard</button>
    <button class="bg-alt hover:opacity-70 py-2 px-4 font-bold text-alt-lightest rounded-xl text-lg" on:click={onClickDivision(ADVANCED)}>Download Advanced</button>
  </div>

</div>
