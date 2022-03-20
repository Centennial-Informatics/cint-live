package controllers

import (
	"log"
	"servermodule/app/models"
	"servermodule/app/scraper"
	"servermodule/configs/constants"
	"servermodule/utils"
	"servermodule/utils/jobs"
	"servermodule/utils/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// Submit godoc
// @Summary Submit code to a problem.
// @Description Submits code and returns a pending verdict.
// @Tags User Actions
// @Accept  application/x-www-form-urlencoded
// @Produce  application/json
// @Param token body string false "Authorization Token"
// @Param id_token body string false "Google ID Token"
// @Param problem body string true "Problem ID"
// @Param language body string true "Language ID"
// @Param submission body string true "Submission source code"
// @Success 200 {object} models.Submission
// @Failure 401
// @Router /submit [post]
func Submit(c *fiber.Ctx, ts *models.TokenService, config *models.Configuration, client *scraper.Client,
	fc *models.FirebaseCache, s *middleware.Submitter, w *middleware.Writer) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		tokenInfo, err := utils.VerifyToken(c.FormValue("id_token"))

		if err != nil {
			log.Println("Unauthorized submission by userID", userID)
			return c.SendStatus(constants.StatusUnauthorized)
		}

		userID = tokenInfo.UserId
	}

	if !fc.Users.Has(userID) {
		log.Println("userID ", userID, "does not exist")
		return c.SendStatus(constants.StatusUnauthorized)
	}

	file, err := c.FormFile("file")

	submission := models.QueuedSubmission{
		UserID:     userID,
		Source:     c.FormValue("submission"),
		ProblemID:  c.FormValue("problem"),
		Language:   c.FormValue("language"),
		File:       file,
		ContestURL: client.Cached.ContestURL,
	}

	if err != nil {
		submission.File = nil
	}

	sub := models.Submission{
		Status:  "Pending",
		Verdict: "Pending",
		Time:    int64(time.Since(config.StartTime) / time.Second),
	}

	fc.Users.Get(userID).Submissions[c.FormValue("problem")] = &sub

	w.AddJob(jobs.SubmissionJob{
		UserID:     userID,
		Submission: sub,
		ProblemID:  c.FormValue("problem"),
	})

	newConns := make([]*websocket.Conn, 0)
	user := ts.GetUserFromID(userID)

	conns := user.Conn
	for _, conn := range conns {
		if conn.Conn != nil {
			newConns = append(newConns, conn)

			err = conn.WriteJSON(fc.Users.Get(userID).Submissions)
			if err != nil {
				return c.SendString("error")
			}
		}
	}

	user.Conn = newConns

	s.Submit(&submission)

	return c.SendStatus(constants.StatusOk)
}

// Submissions godoc
// @Summary Get submission verdicts for a user.
// @Description Returns a JSON object with updated verdicts for user's most recent submission on each problem.
// @Tags User Actions
// @Produce  application/json
// @Param token body string false "Authorization Token"
// @Param id_token body string false "Google ID Token"
// @Success 200 {array} models.Submission
// @Failure 401
// @Router /submissions [post]
func Submissions(c *fiber.Ctx, ts *models.TokenService, fc *models.FirebaseCache) error {
	userID, err := ts.AuthorizeUser(c.FormValue("token"))
	if err != nil {
		tokenInfo, err := utils.VerifyToken(c.FormValue("id_token"))

		if err != nil {
			return c.SendStatus(constants.StatusUnauthorized)
		}

		userID = tokenInfo.UserId
	}

	if !fc.Users.Has(userID) {
		return c.SendStatus(constants.StatusUnauthorized)
	}

	return c.JSON(fc.Users.Get(userID).Submissions)
}
