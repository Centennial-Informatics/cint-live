package configs

import (
	"servermodule/app/models"
)

func QuHacks() (*models.Configuration, error) {
	return Config("quhacks.yml", false)
}
