package jobs

import (
	"log"
	"servermodule/app/models"
)

/* SubmissionJob - Job to update a user's most recent submission to a problem in the contest sheet. */
type SubmissionJob struct {
	UserID     string
	ProblemID  string
	Submission models.Submission
}

func (j SubmissionJob) Run(f *models.FirebaseService) {
	log.Println("SubmissionJob on", j.UserID, "problem", j.ProblemID, "verdict", j.Submission.Verdict)
	f.UpdateSubmission(j.UserID, j.ProblemID, j.Submission)
}
