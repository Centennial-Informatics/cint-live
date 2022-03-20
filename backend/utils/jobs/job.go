package jobs

import "servermodule/app/models"

type Job interface {
	Run(f *models.FirebaseService)
}
