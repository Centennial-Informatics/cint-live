package scraper

import (
	"encoding/json"
	"log"
	"os"
	"servermodule/app/models"
	"servermodule/configs"
	"testing"
)

func TestInfo(t *testing.T) {
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
	if cachedData.Info.Name != configs.TestData().Info.Name {
		t.Fail()
	}

	if cachedData.Info.Description != configs.TestData().Info.Description {
		t.Fail()
	}

	if cachedData.Info.Website != configs.TestData().Info.Website {
		t.Fail()
	}
}
