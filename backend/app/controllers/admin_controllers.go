package controllers

import (
	"servermodule/app/database"
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
		email := c.FormValue("Email")
		if email == "" {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		division := c.FormValue("Division")
		if division == "" {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		name := c.FormValue("Name")
		if name == "" {
			return c.SendStatus(constants.StatusUnauthorized)
		}
		db.CreateUser(name, email, division)
	return c.SendStatus(constants.StatusOk)
}