package configs

import (
	"servermodule/app/models"
)

func Practice() (*models.Configuration, error) {
	return Config("practice.yml", false)
}
