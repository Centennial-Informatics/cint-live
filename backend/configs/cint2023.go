package configs

import (
	"servermodule/app/models"
)

func Cint2023() (*models.Configuration, error) {
	return Config("cint2023.yml", false)
}
