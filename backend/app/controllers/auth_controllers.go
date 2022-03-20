package controllers

import (
	"servermodule/app/models"
	"servermodule/app/scraper"
	"servermodule/configs/constants"
	"servermodule/utils"
	"servermodule/utils/jobs"
	"servermodule/utils/middleware"

	"github.com/gofiber/fiber/v2"
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
func Login(c *fiber.Ctx, config *models.Configuration, ts *models.TokenService, client *scraper.Client,
	f *models.FirebaseService, w *middleware.Writer) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		tokenInfo, err := utils.VerifyToken(c.FormValue("id_token"))

		if err != nil {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		if !f.Cache.Users.Has(tokenInfo.UserId) {
			f.Cache.Users.Store(tokenInfo.UserId, f.NewUserEntity(c.FormValue("username"), tokenInfo.Email))

			w.AddJob(jobs.UserJob{
				UserID:   tokenInfo.UserId,
				Username: tokenInfo.Email,
			})
		}

		userID = tokenInfo.UserId
	}

	return c.SendString(ts.UpdateToken(userID, config.AuthTokenLength))
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
func GetProfile(c *fiber.Ctx, ts *models.TokenService, f *models.FirebaseService) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	return c.JSON(f.Cache.Users.Get(userID))
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
func UpdateProfile(c *fiber.Ctx, ts *models.TokenService, f *models.FirebaseService, w *middleware.Writer) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	user := f.Cache.Users.Get(userID)

	if username := c.FormValue("username"); username != "" {
		user.Username = username
	}

	if email := c.FormValue("email"); email != "" {
		user.Email = email
	}

	if f.Cache.Users.Has(userID) {
		f.Cache.Users.Store(userID, user)

		w.AddJob(jobs.UpdateUserJob{
			UserID:   userID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return c.SendStatus(constants.StatusOk)
}
