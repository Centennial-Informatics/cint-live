package jobs

import (
	"log"
	"servermodule/app/models"
)

/* UserJob - Job to add a new user to the contest sheet. */
type UserJob struct {
	UserID   string
	Username string
	Email    string
}

func (j UserJob) Run(f *models.FirebaseService) {
	log.Println("UserJob on", j.UserID, "username", j.Username, "email", j.Email)
	f.AddUser(j.UserID, j.Username, j.Email)
}
