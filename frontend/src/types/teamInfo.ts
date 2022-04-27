export interface TeamInfo {
	name: string;
	code: string;
	division: string;
	ID: string;
}

export interface TeamData {
	error?: string;
	team: TeamInfo;
	members: TeammateInfo[];
}

export interface TeammateInfo {
	username: string;
	division: string;
}

import type { STANDARD, ADVANCED } from '$lib/data/constants/division';

export type Division = typeof STANDARD | typeof ADVANCED;
