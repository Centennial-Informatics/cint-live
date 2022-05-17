package controllers

import (
	"servermodule/app/database"
	"servermodule/app/models"
	"servermodule/configs/constants"

	"github.com/gofiber/fiber/v2"
)

// func CreateUser(c *fiber.Ctx, ts *models.TokenService,
// 	f *models.FirebaseService, w *workers.Writer) error {
// 		email := c.FormValue("Email")
// 		if email == "" {
// 			return c.SendStatus(constants.StatusUnauthorized)
// 		}
// 		newTeam := f.NewTeamEntity("Team #"+strconv.Itoa(f.TeamCount()+1), 1)
// 		if f.Cache.Teams.Has("12345") {
// 			newTeam = f.Cache.Teams.Get("12345")
// 		}
// 		newTeamID := "12345" //utils.GenerateSecureToken(3)
// 		f.Cache.Teams.Store(newTeamID, newTeam)
// 		f.Cache.Users.Store(email, f.NewUserEntity(email, email, newTeamID))

// 		w.AddJob(jobs.UserJob{
// 			UserID:   email,
// 			Username: email,
// 			TeamID:   newTeamID,
// 			Email:    email,
// 		})

// 		w.AddJob(jobs.TeamJob{
// 			TeamID:      newTeamID,
// 			Name:        (*newTeam).Name,
// 			MemberCount: 1,
// 		})
// 	return c.SendStatus(constants.StatusOk)
// }

func CreateUser(c *fiber.Ctx, db *database.ContestDB) error {
	email := c.FormValue("email")
	if email == "" {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	division := c.FormValue("division")
	if division == "" {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	name := c.FormValue("name")
	if name == "" {
		return c.SendStatus(constants.StatusUnauthorized)
	}
	db.CreateUser(name, email, division)
	return c.SendStatus(constants.StatusOk)
}

/* Clear all submissions from the submissions table. For testing purposes. */
func ClearSubmissions(c *fiber.Ctx, db *database.ContestDB) error {
	db.ClearSubmissions()

	return c.SendStatus(constants.StatusOk)
}

func DeleteUser(c *fiber.Ctx, db *database.ContestDB, ts *models.TokenService) error {
	db.DeleteUser(c.FormValue("email"))
	ts.DeleteUser(c.FormValue("email"))

	return c.SendStatus(constants.StatusOk)
}

func DownloadStandings(c *fiber.Ctx, db *database.ContestDB, config *models.Configuration) error {
	if c.FormValue("division") == "Standard" {
		database.WriteStandingsToCSV(db, db.StandardCache, "Standard", config)
		return c.SendFile("StandingsStandard.csv", true)
	}

	database.WriteStandingsToCSV(db, db.AdvancedCache, "Advanced", config)
	return c.SendFile("StandingsAdvanced.csv", true)
}

func ConnectionMetrics(c *fiber.Ctx, ts *models.TokenService) error {
	return c.JSON(ts.GetConnMetrics())
}

func MakeAnnouncement(c *fiber.Ctx, ts *models.TokenService) error {
	if c.FormValue("title") != "" {
		err := ts.Announce(c.FormValue("title"), c.FormValue("details"))
		if err != nil {
			return c.SendStatus(constants.StatusError)
		}
	}

	return c.SendStatus(constants.StatusOk)
}
