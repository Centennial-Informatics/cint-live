package workers

import (
	"log"
	"servermodule/app/models"
	"servermodule/app/scraper"
	"time"
)

type Submitter struct {
	Clients         []*scraper.Client
	SubmissionQueue []*models.QueuedSubmission
}

func (s *Submitter) Submit(submission *models.QueuedSubmission) {
	if s.SubmissionQueue == nil {
		s.SubmissionQueue = make([]*models.QueuedSubmission, 0)
	}

	s.SubmissionQueue = append(s.SubmissionQueue, submission)
}

func (s *Submitter) ProcessNextSubmission(client *scraper.Client) {
	if len(s.SubmissionQueue) > 0 {
		log.Println("Processing submission")

		sub := s.SubmissionQueue[0]
		s.SubmissionQueue = s.SubmissionQueue[1:]

		go func() {
			submissionID, err := client.Submit(sub.ContestURL, sub.ProblemID,
				sub.Source, sub.File, sub.Language)

			client.Verdict[submissionID].UserID = &sub.UserID
			client.Verdict[submissionID].SubmissionID = sub.SubmissionID

			if err != nil {
				log.Println(err)
			}
		}()
	}
}

/* UpdateVerdicts - Updates cached verdicts only. */
func (s *Submitter) UpdateVerdicts(client *scraper.Client, cbs ...models.SubmissionCallback) {
	for submissionID := range client.Verdict {
		v := client.Status(client.Cached.ContestURL, submissionID)
		for _, cb := range cbs {
			cb(submissionID, v)
		}
	}
}

func SubmitWorker(clients []*scraper.Client, interval time.Duration,
	cbs ...models.SubmissionCallback) (*Submitter, chan struct{}) {
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	submitter := Submitter{
		Clients: clients,
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				for _, client := range submitter.Clients {
					if client.Available {
						submitter.ProcessNextSubmission(client)
					}

					go submitter.UpdateVerdicts(client, cbs...)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return &submitter, quit
}
