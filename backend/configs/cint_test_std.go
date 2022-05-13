package configs

import (
	"servermodule/app/models"
)

func CintTestStd() (*models.Configuration, error) {
	return Config("cint_test_std.yml", false)
}
