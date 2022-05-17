<script lang="ts">
	import Sponsors from '$lib/data/sponsors';
	import { brand } from '$lib/data/stores/userInfo';
	import { SponsorTiers } from '../../../../../src/types/sponsors';
	import SponsorRecognition from './sponsorRecognition.svelte';

	export let bumpUp = false;

	const order = {
		[SponsorTiers.GOLD]: 1,
		[SponsorTiers.SILVER]: 2,
		[SponsorTiers.BRONZE]: 3
	};

	const sortedSponsors = Sponsors.sort((a, b) => {
		return order[a.tier] - order[b.tier];
	});
</script>

<div class="text-white w-full text-center py-10 {bumpUp ? '-mt-32' : ''} px-20 space-y-28">
	<div class="flex flex-wrap items-baseline space-y-10">
		{#each sortedSponsors as sponsor (sponsor.name)}
			<SponsorRecognition
				name={sponsor.name}
				tier={sponsor.tier}
				logo={sponsor.logo}
				link={sponsor.link}
			/>
		{/each}
	</div>

	<div>
		Organized by students at
		<a
			class="text-{$brand}-light font-bold hover:underline"
			href="https://chs.hcpss.org"
			target="_blank">Centennial High School</a
		>
	</div>
</div>
