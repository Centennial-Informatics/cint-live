package routes

import (
	"os"
	"servermodule/app/controllers"
	"servermodule/app/database"
	"servermodule/app/models"
	"servermodule/configs/constants"

	"github.com/gofiber/fiber/v2"
)

/* AdminAPIRoutes - Public routes requiring a passcode. */
func AdminAPIRoutes(router fiber.Router, config *models.Configuration,
	db *database.ContestDB) {
	router.Post("/admin/create", func(c *fiber.Ctx) error {
		if c.FormValue("ADMIN_TOKEN") != os.Getenv("ADMIN_TOKEN") {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		return controllers.CreateUser(c, db)
	})
}
