package controllers

import (
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

		if err != nil {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		if !db.HasUser(tokenInfo.Email) {
			return c.JSON(map[string]string{
				"error": "You are not registered for this competition.",
			})
		}

		userID = tokenInfo.Email
	}

	team, _ := db.GetTeamByCode(db.GetUser(userID).TeamCode)

	return c.JSON(map[string]string{
		"token":   ts.UpdateToken(userID, team.Code, config.AuthTokenLength),
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
func UpdateUserTeam(c *fiber.Ctx, ts *models.TokenService, db *database.ContestDB, config *models.Configuration) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		return c.SendStatus(constants.StatusUnauthorized)
	}

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
		team = *db.UpdateUserTeam(userID, teamID)
	} else {
		team, _ = db.GetTeamByCode(userTeamCode)
	}

	return c.JSON(map[string]interface{}{
		"team":    team,
		"members": db.GetTeamMembers(team.Code),
	})
}

func UpdateTeam(c *fiber.Ctx, ts *models.TokenService, db *database.ContestDB) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	if name := c.FormValue("team_name"); name != "" {
		user := db.GetUser(userID)
		db.UpdateTeam(user.TeamCode, name)
	}

	return c.SendStatus(constants.StatusOk)
}

func LeaveTeam(c *fiber.Ctx, ts *models.TokenService, db *database.ContestDB) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	db.LeaveTeam(userID)

	team, members := db.GetTeamByCode(db.GetUser(userID).TeamCode)

	return c.JSON(map[string]interface{}{
		"team":    team,
		"members": members,
	})
}
