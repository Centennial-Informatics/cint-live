package routes

import (
	"servermodule/app/controllers"
	"servermodule/app/database"
	"servermodule/app/models"
	"servermodule/app/scraper"
	"servermodule/configs/constants"
	"time"

	"github.com/gofiber/fiber/v2"
)

/* PublicRoutes - Page urls. */
func PublicRoutes(app *fiber.App) {
	// app.Get("/", hello)
	app.Get("/", controllers.ContestPage)
	app.Get("/standings", controllers.ContestPage)
	app.Get("/problem/*", controllers.ContestPage)
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendStatus(constants.StatusOk)
	})
}

/* PublicAPIRoutes - Public api routes. */
func PublicAPIRoutes(router fiber.Router, config *models.Configuration, standardClient *scraper.Client) {
	router.Get("/schedule", func(c *fiber.Ctx) error {
		return controllers.Schedule(c, config)
	})
	router.Get("/about", func(c *fiber.Ctx) error {
		return controllers.About(c, standardClient)
	})
}

/* PublicTimeAPIRoutes - Public time-restricted api routes. */
func PublicTimeAPIRoutes(router fiber.Router, config *models.Configuration, standardClient *scraper.Client, advancedClient *scraper.Client, clientPre *scraper.Client,
	db *database.ContestDB) {
	router.Get("/collect/standard", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) < 0 {
			return controllers.Collect(c, config, clientPre)
		}

		return controllers.Collect(c, config, standardClient)
	})

	router.Get("/collect/advanced", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) < 0 {
			return controllers.Collect(c, config, clientPre)
		}

		return controllers.Collect(c, config, advancedClient)
	})
	router.Get("/standings", func(c *fiber.Ctx) error {
		// if time.Since(config.StartTime) < 0 {
		// 	return c.SendStatus(constants.StatusUnauthorized)
		// }

		return controllers.Standings(c, db)
	})

	// router.Get("/problems", func(c *fiber.Ctx) error {
	// 	if time.Since(config.StartTime) < 0 {
	// 		return c.SendStatus(constants.StatusUnauthorized)
	// 	}

	// 	return controllers.Problems(c, client)
	// })
	// router.Get("/problems/:problemId", func(c *fiber.Ctx) error {
	// 	if time.Since(config.StartTime) < 0 {
	// 		return c.SendStatus(constants.StatusUnauthorized)
	// 	}

	// 	return controllers.ProblemPage(c, client)
	// })
}
