package routes

import (
	"servermodule/app/controllers"
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
}

/* PublicAPIRoutes - Public api routes. */
func PublicAPIRoutes(router fiber.Router, config *models.Configuration, client *scraper.Client,
	f *models.FirebaseService) {
	router.Post("/collect", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) < 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		return controllers.Collect(c, config, client)
	})

	router.Get("/schedule", func(c *fiber.Ctx) error {
		return controllers.Schedule(c, config)
	})
	router.Get("/about", func(c *fiber.Ctx) error {
		return controllers.About(c, client)
	})
}

/* PublicTimeAPIRoutes - Public time-restricted api routes. */
func PublicTimeAPIRoutes(router fiber.Router, config *models.Configuration, client *scraper.Client,
	f *models.FirebaseService) {
	router.Get("/standings", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) < 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		return controllers.Standings(c, f.Cache)
	})

	router.Get("/problems", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) < 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		return controllers.Problems(c, client)
	})
	router.Get("/problems/:problemId", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) < 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		return controllers.ProblemPage(c, client)
	})
}
