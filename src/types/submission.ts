import type { BAD, ERROR, GOOD, INFO, PENDING, UNSUBMITTED } from '$lib/data/constants/status';

export interface Submission {
	file?: File;
	submission?: string;
	problem: string;
	language: string;
}

export type VerdictStatus =
	| typeof GOOD
	| typeof BAD
	| typeof PENDING
	| typeof INFO
	| typeof UNSUBMITTED
	| typeof ERROR;
