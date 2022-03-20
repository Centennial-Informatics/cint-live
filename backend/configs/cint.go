package configs

import (
	"servermodule/app/models"
)

func Cint() (*models.Configuration, error) {
	return Config("cint.yml", false)
}
