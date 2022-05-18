export interface Sponsor {
	name: string;
	link: string;
	logo: string;
	tier: SponsorTiers;
}

export enum SponsorTiers {
	PLATINUM = 'Platinum',
	GOLD = 'Gold',
	SILVER = 'Silver',
	BRONZE = 'Bronze'
}
