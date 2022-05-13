import { BAD, ERROR, GOOD, PENDING, UNSUBMITTED } from '$lib/data/constants/status';
import type { SubmissionVerdict, SubmissionVerdictUpdate } from 'src/types/contestData';
import type { VerdictStatus } from 'src/types/submission';

/* Convert something like "wrong answer on test 7" to BAD */
export default function verdictStatus(verdict: SubmissionVerdict): VerdictStatus {
	let status: VerdictStatus = UNSUBMITTED;
	if (verdict.verdict === 'Accepted') status = GOOD;
	else if (verdict.status === 'Pending') status = PENDING;
	else if (verdict.verdict.length >= 5 && verdict.verdict.substring(0, 5) === 'Wrong') status = BAD;
	else if (verdict.verdict !== 'Unsubmitted') status = ERROR;
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

export function convertVerdictsToMap(data: SubmissionVerdict[], problemIDs: string[]) {
	const res: SubmissionVerdictUpdate = {};
	data.forEach((s) => {
		if (!(s.problem_id in res) || s.time > res[s.problem_id].time) {
			res[s.problem_id] = s;
		}
	});
	problemIDs.forEach((problemID) => {
		if (!(problemID in res)) {
			res[problemID] = {
				ID: '',
				problem_id: problemID,
				verdict: 'Unsubmitted',
				status: 'Final',
				time: 0,
				points: 0
			};
		}
	});

	return res;
}

/* Replace missing verdicts and delete excess verdicts (deleted problems). */
export function fillEmptyVerdicts(data: SubmissionVerdictUpdate, problemIDs: string[]) {
	const newData: SubmissionVerdictUpdate = {};
	problemIDs.forEach((problemID) => {
		if (!(problemID in data)) {
			newData[problemID] = {
				ID: '',
				problem_id: problemID,
				verdict: 'Unsubmitted',
				status: 'Final',
				time: 0,
				points: 0
			};
		} else {
			newData[problemID] = data[problemID];
		}
	});

	return newData;
}
