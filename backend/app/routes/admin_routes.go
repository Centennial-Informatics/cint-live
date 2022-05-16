package routes

import (
	"servermodule/app/controllers"
	"servermodule/app/database"
	"servermodule/app/models"
	"servermodule/configs/constants"

	"github.com/gofiber/fiber/v2"
)

func isAdminAcc(ADMIN_ACCOUNTS []string, email string) bool {
	for _, adminEmail := range ADMIN_ACCOUNTS {
		if adminEmail == email {
			return true
		}
	}
	return false
}

func handleAdminAuth(c *fiber.Ctx, ADMIN_ACCOUNTS []string, ts *models.TokenService) bool {
	userEmail, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		return false
	}
	// if c.FormValue("ADMIN_TOKEN") != os.Getenv("ADMIN_TOKEN") {
	// 	return false
	// }
	if !isAdminAcc(ADMIN_ACCOUNTS, userEmail) {
		return false
	}
	return true
}

/* AdminAPIRoutes - Public routes requiring a passcode. */
func AdminAPIRoutes(router fiber.Router, ADMIN_ACCOUNTS []string, config *models.Configuration, ts *models.TokenService,
	db *database.ContestDB) {
	// router.Post("/admin/create", func(c *fiber.Ctx) error {
	// 	if c.FormValue("ADMIN_TOKEN") == os.Getenv("ADMIN_TOKEN") {
	// 		return controllers.CreateUser(c, db)
	// 	}
	// 	return c.SendStatus(constants.StatusUnauthorized)
	// })
	// router.Post("/admin/dangerous/clear", func(c *fiber.Ctx) error {
	// 	if handleAdminAuth(c, ADMIN_ACCOUNTS, ts) {
	// 		return controllers.ClearSubmissions(c, db)
	// 	}
	// 	return c.SendStatus(constants.StatusUnauthorized)

	// })
	// router.Post("/admin/dangerous/delete", func(c *fiber.Ctx) error {
	// 	if handleAdminAuth(c, ADMIN_ACCOUNTS, ts) {
	// 		return controllers.DeleteUser(c, db)
	// 	}
	// 	return c.SendStatus(constants.StatusUnauthorized)
	// })
	router.Post("/admin/standings", func(c *fiber.Ctx) error {
		if handleAdminAuth(c, ADMIN_ACCOUNTS, ts) {
			return controllers.SecretStandings(c, db)
		}
		return c.SendStatus(constants.StatusUnauthorized)
	})
	router.Post("/admin/download", func(c *fiber.Ctx) error {
		if handleAdminAuth(c, ADMIN_ACCOUNTS, ts) {
			return controllers.DownloadStandings(c, db, config)
		}
		return c.SendStatus(constants.StatusUnauthorized)
	})
	router.Get("/admin/connections", func(c *fiber.Ctx) error {
		if handleAdminAuth(c, ADMIN_ACCOUNTS, ts) {
			return controllers.ConnectionMetrics(c, ts)
		}
		return c.SendStatus(constants.StatusUnauthorized)
	})
}
