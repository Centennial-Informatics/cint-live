package controllers

import (
	"log"
	"servermodule/app/database"
	"servermodule/app/models"
	"servermodule/app/scraper"
	"servermodule/configs/constants"
	"servermodule/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/oauth2/v2"
)

// Login godoc
// @Summary Login to an account.
// @Description Returns an authorization token to be used in place of Google creds in future api requests.
// @Tags Authorization
// @Accept  application/x-www-form-urlencoded
// @Produce  text/plain
// @Param token body string false "Authorization Token"
// @Param id_token body string false "Google ID Token"
// @Success 200 {string} string abcdefg123456789
// @Failure 401
// @Router /login [post]
func Login(c *fiber.Ctx, config *models.Configuration, ts *models.TokenService, client *scraper.Client, db *database.ContestDB) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	var tokenInfo *oauth2.Tokeninfo
	if err != nil {
		tokenInfo, err = utils.VerifyToken(c.FormValue("id_token"))

		/* This will (most likely) only occur maliciously, unless the whole google sign-in flow breaks. */
		if err != nil {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		contestStarted := !utils.IsBefore(config.StartTime)
		hasUser := db.HasUser(tokenInfo.Email)

		/* Only new logins before contest begins. Mid-contest bans are permanent. */
		if (config.AllowLateRegistration || !contestStarted) && !hasUser {
			db.CreateUser(c.FormValue("name"), tokenInfo.Email, "Standard")
		} else if !hasUser {
			return c.JSON(map[string]string{
				"error": "No new registrations allowed. Contest has already started. Make sure you're using the correct Google email. Ask an organizer if you need help.",
			})
		}

		userID = tokenInfo.Email
	}

	team, _ := db.GetTeamByCode(db.GetUser(userID).TeamCode)

	log.Println("Login stuff", userID, team.Code)

	token := ts.UpdateToken(userID, team.Code, config.AuthTokenLength)

	log.Println("logging in", token)

	return c.JSON(map[string]string{
		"token":   token,
		"team_id": team.Code,
	})
}

// GetProfile godoc
// @Summary Get account profile info.
// @Description Returns username, email, submissions, and points.
// @Tags User Actions
// @Accept  application/x-www-form-urlencoded
// @Produce  text/plain
// @Param token body string false "Authorization Token"
// @Success 200 {object} models.UserEntity
// @Failure 401
// @Router /profile [post]
func GetProfile(c *fiber.Ctx, ts *models.TokenService, db *database.ContestDB) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	return c.JSON(db.GetUser(userID))
}

func GetTeam(c *fiber.Ctx, ts *models.TokenService, db *database.ContestDB) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	team, members := db.GetTeamByCode(db.GetUser(userID).TeamCode)

	return c.JSON(map[string]interface{}{
		"team":    team,
		"members": members,
	})
}

// UpdateProfile godoc
// @Summary Update account profile info.
// @Description Update username or email.
// @Tags User Actions
// @Accept  application/x-www-form-urlencoded
// @Produce  text/plain
// @Param token body string false "Authorization Token"
// @Param username body string false "Username"
// @Param email body string false "Email"
// @Success 200
// @Failure 401
// @Router /update [post]
func UpdateUserTeam(c *fiber.Ctx, userID string, db *database.ContestDB, config *models.Configuration, ts *models.TokenService) error {
	var team database.Team

	userTeamCode := db.GetUser(userID).TeamCode

	if teamID := c.FormValue("team_code"); teamID != "" {
		teamID := strings.ToLower(teamID)
		if !db.HasTeam(teamID) {
			return c.JSON(map[string]interface{}{
				"error": "Team not found.",
			})
		}
		if userTeamCode == teamID {
			return c.JSON(map[string]interface{}{
				"error": "You are already on that team.",
			})
		}
		members := db.GetTeamMembers(teamID)
		if len(members) == config.MaxTeamSize {
			return c.JSON(map[string]interface{}{
				"error": "Team is full.",
			})
		}
		ts.ChangeConnections(userID, teamID)
		team = *db.UpdateUserTeam(userID, teamID)
	} else {
		team, _ = db.GetTeamByCode(userTeamCode)
	}

	return c.JSON(map[string]interface{}{
		"team":    team,
		"members": db.GetTeamMembers(team.Code),
	})
}

func UpdateTeam(c *fiber.Ctx, userID string, db *database.ContestDB) error {
	name := c.FormValue("team_name")

	if name == "" {
		return c.SendString("Name is too short.")
	} else if len(name) > 30 {
		return c.SendString("Name can be at most 30 characters.")
	} else {
		user := db.GetUser(userID)
		db.UpdateTeamName(user.TeamCode, name)
	}

	return c.SendString("")
}

func LeaveTeam(c *fiber.Ctx, userID string, db *database.ContestDB, ts *models.TokenService) error {
	db.LeaveTeam(userID)

	team, members := db.GetTeamByCode(db.GetUser(userID).TeamCode)
	ts.ChangeConnections(userID, team.Code)

	return c.JSON(map[string]interface{}{
		"team":    team,
		"members": members,
	})
}

func UpdateTeamDivision(c *fiber.Ctx, userID string, db *database.ContestDB) error {
	code := db.GetUser(userID).TeamCode
	// team, _ := db.GetTeamByCode(code)

	if division := c.FormValue("division"); division != "" {
		if division == "Standard" || division == "Advanced" {
			db.UpdateTeamDivision(code, division)
		}
	}

	team, members := db.GetTeamByCode(db.GetUser(userID).TeamCode)

	return c.JSON(map[string]interface{}{
		"team":    team,
		"members": members,
	})
}
