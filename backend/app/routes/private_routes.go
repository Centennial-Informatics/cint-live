package routes

import (
	"log"
	"servermodule/app/controllers"
	"servermodule/app/database"
	"servermodule/app/models"
	"servermodule/app/scraper"
	"servermodule/configs/constants"
	"servermodule/utils"
	"servermodule/utils/workers"
	"time"

	"github.com/gofiber/fiber/v2"
)

/* PrivateAPIRoutes - Private login-restricted api endpoints. */
func PrivateAPIRoutes(router fiber.Router, config *models.Configuration,
	ts *models.TokenService, client *scraper.Client, db *database.ContestDB) {
	router.Get("/languages", func(c *fiber.Ctx) error {
		return controllers.Languages(c, client)
	})

	router.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, config, ts, client, db)
	})

	router.Post("/profile", func(c *fiber.Ctx) error {
		return controllers.GetProfile(c, ts, db)
	})

	router.Post("/team", func(c *fiber.Ctx) error {
		return controllers.GetTeam(c, ts, db)
	})
}

/* PrivateTimeAPIRoutes - Private login-restricted and time-restricted api endpoints. */
func PrivateTimeAPIRoutes(router fiber.Router, config *models.Configuration,
	ts *models.TokenService, standardClient *scraper.Client, advancedClient *scraper.Client, practiceClient *scraper.Client,
	s *workers.Submitter, db *database.ContestDB) {
	divisionClient := map[string]*scraper.Client{
		"Standard": standardClient,
		"Advanced": advancedClient,
	}
	// during contest
	router.Post("/submit", func(c *fiber.Ctx) error {
		userID, err := ts.AuthorizeUser(c.FormValue("token"))
		if err != nil || utils.IsAfter(config.StopTime) { // no new submissions after contest ends
			if err != nil {
				log.Println("Submission error:", err)
			}
			return c.SendStatus(constants.StatusUnauthorized)
		}

		if utils.IsBefore(config.StartTime) {
			return controllers.Submit(c, userID, ts, config, practiceClient, s, db)
		} else if c.FormValue("problem") == "S1" || c.FormValue("problem") == "S2" {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		team, _ := db.GetTeamByCode(db.GetUser(userID).TeamCode)

		return controllers.Submit(c, userID, ts, config, divisionClient[team.Division], s, db)
	})

	// before contest
	// TODO: change < 0 to >= 0
	router.Post("/join", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) >= 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}
		userID, err := ts.AuthorizeUser(c.FormValue("token"))
		if err != nil {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		return controllers.UpdateUserTeam(c, userID, db, config, ts)
	})

	router.Post("/update", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) >= 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}
		userID, err := ts.AuthorizeUser(c.FormValue("token"))
		if err != nil {
			return c.SendStatus(constants.StatusUnauthorized)
		}
		return controllers.UpdateTeam(c, userID, db)
	})

	router.Post("/division", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) >= 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}
		userID, err := ts.AuthorizeUser(c.FormValue("token"))
		if err != nil {
			return c.SendStatus(constants.StatusUnauthorized)
		}
		return controllers.UpdateTeamDivision(c, userID, db)
	})

	router.Post("/leave", func(c *fiber.Ctx) error {
		if time.Since(config.StartTime) >= 0 {
			return c.SendStatus(constants.StatusUnauthorized)
		}
		userID, err := ts.AuthorizeUser(c.FormValue("token"))
		if err != nil {
			return c.SendStatus(constants.StatusUnauthorized)
		}
		return controllers.LeaveTeam(c, userID, db, ts)
	})
}
