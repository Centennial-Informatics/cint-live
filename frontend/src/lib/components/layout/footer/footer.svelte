<script lang="ts">
	import Sponsors from '$lib/data/sponsors';
	import { brand } from '$lib/data/stores/userInfo';
	import { SponsorTiers } from '../../../../../src/types/sponsors';
	import SponsorRecognition from './sponsorRecognition.svelte';

	export let bumpUp = false;

	const order = {
		[SponsorTiers.PLATINUM]: 0,
		[SponsorTiers.GOLD]: 1,
		[SponsorTiers.SILVER]: 2,
		[SponsorTiers.BRONZE]: 3
	};

	const labelStyle: Record<string, string> = {
		[SponsorTiers.PLATINUM]: 'text-sky-300',
		[SponsorTiers.GOLD]: 'text-amber-400',
		[SponsorTiers.SILVER]: 'text-slate-400',
		[SponsorTiers.BRONZE]: 'text-orange-500'
	};

	const dividerStyle: Record<string, string> = {
		[SponsorTiers.PLATINUM]: 'border-sky-300',
		[SponsorTiers.GOLD]: 'border-amber-400',
		[SponsorTiers.SILVER]: 'border-slate-400',
		[SponsorTiers.BRONZE]: 'border-orange-500'
	};

	const sponsorsByTier = Object.keys(order).map((tier) => Sponsors.filter((s) => s.tier === tier));
</script>

<div class="text-white w-full text-center py-10 px-8 {bumpUp ? '-mt-32' : ''} space-y-28">
	<div class="flex flex-col items-baseline space-y-10">
		<span class="font-bold">Thank you to all of our generous sponsors:</span>
		{#each Object.keys(order) as tier, i}
			<div class="font-bold text-xl flex flex-row items-center w-full space-x-4 {labelStyle[tier]}">
				<span>{tier}</span>
				<div class="w-full {dividerStyle[tier]} border-t h-0" />
			</div>
			<div class="flex flex-row flex-wrap items-baseline space-x-4">
				{#each sponsorsByTier[i] as sponsor (sponsor.name)}
					<SponsorRecognition
						name={sponsor.name}
						tier={sponsor.tier}
						logo={sponsor.logo}
						link={sponsor.link}
					/>
				{/each}
			</div>
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
