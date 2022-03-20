package configs

import (
	"io/ioutil"
	"servermodule/app/models"
	"testing"

	"gopkg.in/yaml.v2"
)

/* TestDefault — test that the config is consistent with the yaml file */
func TestDefault(t *testing.T) {
	defaultFile, err := ioutil.ReadFile("./data/default.yml")
	if err != nil {
		t.Error()
	}

	shouldBe := models.Configuration{}

	err = yaml.Unmarshal(defaultFile, &shouldBe)
	if err != nil {
		t.Error()
	}

	foundConfig, err := Config("default.yml", true)
	if err != nil {
		t.Error(err)
	}

	if foundConfig.ContestID != shouldBe.ContestID {
		t.Fail()
	}

	if foundConfig.ContestURL != shouldBe.ContestURL {
		t.Fail()
	}
}

/* TestConfig — test that unspecified config fields receive the proper default values */
func TestConfig(t *testing.T) {
	defaults, err := Config("default.yml", true)
	if err != nil {
		t.Error(err)
	}

	testConfig, err := Config("testdata.yml", true)
	if err != nil {
		t.Error(err)
	}

	if testConfig.AuthTokenLength != defaults.AuthTokenLength {
		t.Fail()
	}

	if testConfig.LogInterval.Milliseconds() == defaults.LogInterval.Milliseconds() {
		t.Fail()
	}
}
