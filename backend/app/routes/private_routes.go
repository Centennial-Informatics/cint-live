package routes

import (
	"servermodule/app/controllers"
	"servermodule/app/models"
	"servermodule/app/scraper"
	"servermodule/configs/constants"
	"servermodule/utils/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

/* PrivateAPIRoutes - Private login-restricted api endpoints. */
func PrivateAPIRoutes(router fiber.Router, config *models.Configuration,
	ts *models.TokenService, client *scraper.Client,
	f *models.FirebaseService, w *middleware.Writer) {
	router.Get("/languages", func(c *fiber.Ctx) error {
		return controllers.Languages(c, client)
	})

	router.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, config, ts, client, f, w)
	})

	router.Post("/profile", func(c *fiber.Ctx) error {
		return controllers.GetProfile(c, ts, f)
	})

	router.Post("/update", func(c *fiber.Ctx) error {
		return controllers.UpdateProfile(c, ts, f, w)
	})
}

/* PrivateTimeAPIRoutes - Private login-restricted and time-restricted api endpoints. */
func PrivateTimeAPIRoutes(router fiber.Router, config *models.Configuration,
	ts *models.TokenService, client *scraper.Client,
	f *models.FirebaseService, s *middleware.Submitter, w *middleware.Writer) {
	router.Post("/submit", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) < 0 || time.Until(config.StopTime) < 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		return controllers.Submit(c, ts, config, client, f.Cache, s, w)
	})

	router.Post("/submissions", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) < 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		return controllers.Submissions(c, ts, f.Cache)
	})
}
