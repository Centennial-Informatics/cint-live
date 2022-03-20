export interface ContestData {
	Info: ContestInfo;
	Languages: ContestLanguage[];
	ProblemPages: { [key: string]: ProblemPage };
	Problems: Problem[];
	StartTime: string;
	StopTime: string;
	Points: { [key: string]: number };
}

export interface ContestInfo {
	Description: string;
	Logo: string;
	Name: string;
	Website: string;
}

export interface ContestLanguage {
	Ext: string;
	ID: string;
	Name: string;
}

export interface ProblemPage {
	Header: {
		Input: string;
		MemLimit: string;
		Output: string;
		TimeLimit: string;
		Title: string;
	};
	InputSpecification: ProblemText[];
	Note: ProblemText[];
	OutputSpecification: ProblemText[];
	SampleTests: ProblemTest[];
	Statement: ProblemText[];
}

export interface ProblemText {
	Code: string;
	Image: string;
	Text: string;
}

export interface ProblemTest {
	Input: string;
	Output: string;
}

export interface Problem {
	ID: string;
	Name: string;
}

export interface SubmissionVerdict {
	ID: string;
	Verdict: string;
	Status: 'Pending' | 'Final';
	Time: number;
	Points: number;
}

export interface SubmissionVerdictUpdate {
	[key: string]: SubmissionVerdict;
}

export interface StandingsEntry {
	Username: string;
	Email: string; //TODO: hide this
	Submissions: SubmissionVerdictUpdate;
	Points: { [key: string]: number };
	TotalPoints: number;
	ID: string;
}

export interface StandingsData {
	[key: string]: StandingsEntry;
}
