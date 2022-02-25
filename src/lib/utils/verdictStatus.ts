import { GOOD, BAD, PENDING, ERROR } from '$lib/data/constants/status';
import type { SubmissionVerdict } from 'src/types/contestData';

/* Convert something like "wrong answer on test 7" to BAD */
export default function verdictStatus(verdict: SubmissionVerdict): string {
	let status = '';
	if (verdict.Verdict === 'Accepted') status = GOOD;
	else if (verdict.Status === 'Pending') status = PENDING;
	else if (verdict.Verdict.length >= 5 && verdict.Verdict.substring(0, 5) === 'Wrong') status = BAD;
	else if (verdict.Verdict !== 'Unsubmitted') status = ERROR;
	return status;
}
