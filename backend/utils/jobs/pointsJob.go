package jobs

import (
	"log"
	"servermodule/app/models"
)

/* PointsJob - Job to update a user's points to a problem in the database. */
type PointsJob struct {
	Points    int
	ProblemID string
	UserID    string
}

func (j PointsJob) Run(f *models.FirebaseService) {
	log.Println("PointsJob on", j.UserID, "problem", j.ProblemID, "points", j.Points)
	f.UpdatePoints(j.UserID, j.ProblemID, j.Points)
}
