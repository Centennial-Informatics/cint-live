package scraper

import (
	"encoding/json"
	"log"
	"os"
	"servermodule/app/models"
	"testing"
)

func TestLanguages(t *testing.T) {
	var accounts []*models.CFAccount

	err := json.Unmarshal([]byte(os.Getenv("CF_ACCOUNTS")), &accounts)
	if err != nil {
		log.Fatal(err)
	}

	client, err := NewClient(accounts[0].User, accounts[0].User, accounts[0].Pass)
	if err != nil {
		t.Error(err)
	}

	languages := client.Languages()
	if len(languages) == 0 {
		t.Fail()
	}
}
