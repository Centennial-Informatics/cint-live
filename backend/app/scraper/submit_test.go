package scraper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"servermodule/app/models"
	"servermodule/configs"
	"testing"
	"time"
)

func TestGoodSubmit(t *testing.T) {
	submission, err := ioutil.ReadFile(configs.TestData().Submission)
	if err != nil {
		t.Error(err)
	}

	var accounts []*models.CFAccount

	err = json.Unmarshal([]byte(os.Getenv("CF_ACCOUNTS")), &accounts)
	if err != nil {
		log.Fatal(err)
	}

	client, err := NewClient(accounts[0].User, accounts[0].User, accounts[0].Pass)
	if err != nil {
		t.Error(err)
	}

	client.InitCache(configs.TestData().ContestURL, configs.TestData().ContestID)

	submissionID, err := client.Submit(configs.TestData().ContestURL, configs.TestData().ProblemID,
		string(submission), nil, configs.TestData().Language)
	if err != nil {
		t.Error(err)
	}

	ticker := time.NewTicker(1 * time.Second)

	var submissionStatus *models.Verdict
	for range ticker.C {
		submissionStatus = client.Status(configs.TestData().ContestURL, submissionID)
		if submissionStatus.Status == "Final" {
			ticker.Stop()

			if submissionStatus.Verdict != "Accepted" {
				t.Error()
			}

			return
		}
	}
}
