import { BAD, ERROR, GOOD, PENDING, UNSUBMITTED } from '$lib/data/constants/status';
import type { SubmissionVerdict } from 'src/types/contestData';
import type { VerdictStatus } from 'src/types/submission';

/* Convert something like "wrong answer on test 7" to BAD */
export default function verdictStatus(verdict: SubmissionVerdict): VerdictStatus {
	let status: VerdictStatus = UNSUBMITTED;
	if (verdict.Verdict === 'Accepted') status = GOOD;
	else if (verdict.Status === 'Pending') status = PENDING;
	else if (verdict.Verdict.length >= 5 && verdict.Verdict.substring(0, 5) === 'Wrong') status = BAD;
	else if (verdict.Verdict !== 'Unsubmitted') status = ERROR;
	return status;
}

export function verdictStatusBgColor(status: VerdictStatus): string {
	switch (status) {
		case GOOD:
			return 'bg-green-700';
		case BAD:
			return 'bg-gray-500';
		case ERROR:
			return 'bg-gray-500';
		case PENDING:
			return 'bg-amber-700';
		default:
			return 'bg-gray-500';
	}
}

export function verdictStatusColor(status: VerdictStatus): string {
	switch (status) {
		case GOOD:
			return 'text-green-400';
		case BAD:
			return 'text-red-400';
		case ERROR:
			return 'text-red-400';
		case PENDING:
			return 'text-amber-500';
		default:
			return 'text-gray-300';
	}
}
