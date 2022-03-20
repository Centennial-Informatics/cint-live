package middleware

import (
	"servermodule/app/models"
	"servermodule/utils/jobs"
	"time"
)

type Writer struct {
	F        *models.FirebaseService
	JobQueue []jobs.Job
}

func (w *Writer) AddJob(j jobs.Job) {
	if w.JobQueue == nil {
		w.JobQueue = make([]jobs.Job, 0)
	}

	w.JobQueue = append(w.JobQueue, j)
}

func (w *Writer) ProcessNextJob() {
	if len(w.JobQueue) > 0 {
		w.JobQueue[0].Run(w.F)

		w.JobQueue = w.JobQueue[1:]
	}
}

func WriteWorker(f *models.FirebaseService, interval time.Duration) (*Writer, chan struct{}) {
	w := Writer{
		F:        f,
		JobQueue: make([]jobs.Job, 0),
	}

	ticker := time.NewTicker(interval)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				w.ProcessNextJob()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return &w, quit
}
