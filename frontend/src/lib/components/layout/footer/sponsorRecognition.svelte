<script lang="ts">
	import { SponsorTiers } from '../../../../types/sponsors';

	export let name: string;
	export let logo: string;
	export let link: string;
	export let tier: SponsorTiers;

	const size = {
		[SponsorTiers.PLATINUM]: 'h-72',
		[SponsorTiers.GOLD]: 'w-72 h-52',
		[SponsorTiers.SILVER]: 'w-48 h-28',
		[SponsorTiers.BRONZE]: 'w-28 h-16'
	};
</script>

<a href={link} target="_blank">
	<div class={size[tier] + ' relative overflow-hidden'}>
		<img class="object-contain w-full w- h-full" src={logo} alt={name} />
	</div>
</a>
