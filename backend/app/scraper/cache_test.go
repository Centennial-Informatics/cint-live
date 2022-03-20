package scraper

import (
	"encoding/json"
	"log"
	"os"
	"servermodule/app/models"
	"servermodule/configs"
	"testing"
)

func TestCache(t *testing.T) {
	var accounts []*models.CFAccount

	err := json.Unmarshal([]byte(os.Getenv("CF_ACCOUNTS")), &accounts)
	if err != nil {
		log.Fatal(err)
	}

	client, err := NewClient(accounts[0].User, accounts[0].User, accounts[0].Pass)
	if err != nil {
		t.Error(err)
	}

	client.InitCache(configs.TestData().ContestURL, configs.TestData().ContestID)

	cachedData := client.Cached
	if len(cachedData.Problems) != len(configs.TestData().Problems) {
		t.Fail()
	}

	if len(cachedData.ProblemPages) != len(configs.TestData().Problems) {
		t.Fail()
	}

	for page := range cachedData.ProblemPages {
		if len(page) == 0 {
			t.Fail()
		}
	}
}
