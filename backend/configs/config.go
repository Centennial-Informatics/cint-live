package configs

import (
	"io/ioutil"
	"path/filepath"
	"servermodule/app/models"
	"servermodule/configs/constants"
	"strings"

	"gopkg.in/yaml.v2"
)

func Config(file string, test bool) (*models.Configuration, error) {
	/* get the defaults */
	var path string

	if test {
		path = "./data/default.yml"
	} else {
		path = "configs/data/default.yml"
	}

	defaultFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := models.Configuration{}

	err = yaml.Unmarshal(defaultFile, &config)
	if err != nil {
		return nil, err
	}

	var rel string

	if test {
		rel = "./data/"
	} else {
		rel = "configs/data/"
	}

	path, err = filepath.Abs(rel + file)
	if err != nil {
		return nil, err
	}

	/* get the specified config */
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	/* go-yaml parses time durations into ns by default, but we want ms */
	config.LogInterval *= constants.NsInMs
	config.ScrapeIntervalAll *= constants.NsInMs
	config.ScrapeIntervalVerdicts *= constants.NsInMs
	config.WriteInterval *= constants.NsInMs

	config.ScoreboardFreeze *= constants.NsInS
	config.ScoreboardUnfreeze *= constants.NsInS
	config.SubmissionGrace *= constants.NsInS

	/* yaml is lowercase, but our problem IDs should be uppercase */
	for k, v := range config.Points {
		delete(config.Points, k)
		config.Points[strings.ToUpper(k)] = v
	}

	return &config, err
}
