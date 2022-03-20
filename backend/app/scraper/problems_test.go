package scraper

import (
	"encoding/json"
	"log"
	"os"
	"servermodule/app/models"
	"servermodule/configs"
	"testing"
)

func TestProblems(t *testing.T) {
	var accounts []*models.CFAccount

	err := json.Unmarshal([]byte(os.Getenv("CF_ACCOUNTS")), &accounts)
	if err != nil {
		log.Fatal(err)
	}

	client, err := NewClient(accounts[0].User, accounts[0].User, accounts[0].Pass)
	if err != nil {
		t.Error(err)
	}

	problems, err := client.Problems(configs.TestData().ContestURL)
	if err != nil {
		t.Error(err)
	}

	if len(problems) != len(configs.TestData().Problems) {
		t.Fail()
	}
}

func TestProblemPage(t *testing.T) {
	client, err := NewClient("Tester", os.Getenv("CF_USER"), os.Getenv("CF_PASS"))
	if err != nil {
		t.Error(err)
	}

	page, err := client.ProblemPage(configs.TestData().ContestURL, "D")
	if err != nil {
		t.Error(err)
	}

	log.Println(page)
}
