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

/* Returns whether a submission was processed or not. */
func (s *Submitter) ProcessNextSubmission(client *scraper.Client) bool {
	if len(s.SubmissionQueue) > 0 {
		sub := s.SubmissionQueue[0]

		// check if submission is on the correct contest scraper
		if sub.ContestURL == client.Cached.ContestURL {
			log.Println("Processing submission", client.Name, sub.SubmissionID, client.Cached.ContestURL)
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

			return true
		} else {
			log.Println("Wrong client, deferring submission ", sub.SubmissionID)
		}
	}

	return false
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
		front := 0
		for {
			select {
			case <-ticker.C:
				submitted := false
				// update all verdicts and submit on the first valid and available worker
				for range clients {
					client := submitter.Clients[front]
					if !submitted && client.Available {
						submitted = submitter.ProcessNextSubmission(client)
					}

					go submitter.UpdateVerdicts(client, cbs...)

					front = (front + 1) % len(submitter.Clients)
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return &submitter, quit
}
