package workers

import (
	"log"
	"servermodule/app/models"
	"servermodule/app/scraper"
	"strconv"
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

				_, good := client.Verdict[submissionID]

				if err != nil || !good { // happens when submitting a non-public java class
					log.Println(err)
					// report compilation error instead of stuck on pending
					client.Verdict[strconv.Itoa(sub.SubmissionID)] = &models.Verdict{
						Status:       "Final",
						Verdict:      "Compilation error", // should be a constant but im too tired to fix it
						UserID:       &sub.UserID,
						SubmissionID: sub.SubmissionID,
						ProblemID:    &sub.ProblemID,
					}
				} else {
					client.Verdict[submissionID].UserID = &sub.UserID
					client.Verdict[submissionID].SubmissionID = sub.SubmissionID
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
		var v *models.Verdict
		if client.Verdict[submissionID].Status != "Final" {
			v = client.Status(client.Cached.ContestURL, submissionID)
		} else { // if verdict was error and manually overriden
			v = client.Verdict[submissionID]
		}
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
				submitted := len(submitter.SubmissionQueue) == 0
				// update all verdicts and submit on the first valid and available worker
				for i := range clients {
					client := submitter.Clients[front]
					if !submitted && client.Available {
						submitted = submitter.ProcessNextSubmission(client)
						front = (front + 1) % len(submitter.Clients)
					}

					go submitter.UpdateVerdicts(submitter.Clients[i], cbs...)
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return &submitter, quit
}
