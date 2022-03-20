package jobs

import (
	"log"
	"servermodule/app/models"
)

/* DeleteUserJob - Job to delete a user from the db. Mostly used for cleanup in tests. */
type DeleteUserJob struct {
	UserID string
}

func (j DeleteUserJob) Run(f *models.FirebaseService) {
	log.Println("DeleteUserJob on", j.UserID)
	f.DeleteUser(j.UserID)
}
