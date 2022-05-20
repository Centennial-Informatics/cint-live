package configs

import (
	"servermodule/app/models"
)

func CintStd() (*models.Configuration, error) {
	return Config("cint_std.yml", false)
}
