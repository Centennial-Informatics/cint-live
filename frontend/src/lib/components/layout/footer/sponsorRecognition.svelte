<script lang="ts">
	import { SponsorTiers } from '../../../../types/sponsors';

	export let name: string;
	export let logo: string;
	export let link: string;
	export let tier: SponsorTiers;

	const size = {
		[SponsorTiers.GOLD]: 'w-96 h-72',
		[SponsorTiers.SILVER]: 'w-72 h-48',
		[SponsorTiers.BRONZE]: 'w-32 h-24'
	};
</script>

<a href={link} target="_blank">
	<div class={size[tier] + ' relative overflow-hidden'}>
		<img class="object-contain w-full h-full" src={logo} alt={name} />
	</div>
</a>
