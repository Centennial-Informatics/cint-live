export interface Sponsor {
	name: string;
	link: string;
	logo: string;
	tier: SponsorTiers;
}

export enum SponsorTiers {
	GOLD = 'Gold',
	SILVER = 'Silver',
	BRONZE = 'Bronze'
}
