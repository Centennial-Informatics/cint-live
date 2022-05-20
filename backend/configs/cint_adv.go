package configs

import (
	"servermodule/app/models"
)

func CintAdv() (*models.Configuration, error) {
	return Config("cint_adv.yml", false)
}
