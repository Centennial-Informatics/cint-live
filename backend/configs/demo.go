package configs

import (
	"servermodule/app/models"
)

func Demo() (*models.Configuration, error) {
	return Config("default.yml", false)
}
