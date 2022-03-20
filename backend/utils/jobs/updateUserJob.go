package jobs

import (
	"log"
	"servermodule/app/models"
)

/* UpdateUserJob - Job to add a new user to the contest sheet. */
type UpdateUserJob struct {
	UserID   string
	Username string
	Email    string
}

func (j UpdateUserJob) Run(f *models.FirebaseService) {
	log.Println("UpdateUserJob on", j.UserID, "username", j.Username, "email", j.Email)

	err := f.UpdateUser(j.UserID, j.Username, j.Email)
	if err != nil {
		log.Println(err)
	}
}
