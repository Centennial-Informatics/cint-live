package configs

import (
	"servermodule/app/models"
)

func CintTestAdv() (*models.Configuration, error) {
	return Config("cint_test_adv.yml", false)
}
