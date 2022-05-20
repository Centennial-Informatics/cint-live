package controllers

import (
	"servermodule/app/models"
	"servermodule/app/scraper"
	"servermodule/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

//  ContestPage - /
func ContestPage(c *fiber.Ctx) error {
	return c.SendFile("../frontend/build/index.html")
}

// Collect godoc
// @Summary Collect all page data.
// @Description Returns a JSON object containing all relevant contest-level information (nothing user-specific).
// @Tags Contest Info
// @Produce  application/json
// @Success 200 {object} models.Page
// @Router /collect [get]
func Collect(c *fiber.Ctx, config *models.Configuration, client *scraper.Client) error {
	page := models.Page{
		Info:         client.Cached.Info,
		Problems:     client.Cached.Problems,
		Languages:    client.Languages(),
		ProblemPages: client.Cached.ProblemPages,
		StartTime:    config.StartTime,
		StopTime:     config.StopTime,
		Points:       config.Points,
	}

	return c.JSON(page)
}

var challenges = []string{
	"https://docs.google.com/document/d/1xODEbqsxtqoFsREiW6t9sh9TOSuU_kAi1u9j_CBz9I4/edit?usp=sharing",
	"https://docs.google.com/document/d/14rD_Za8hU_NicYtu93FT5jRvD1ODDzwcrIAPngfjIpw/edit?usp=sharing",
	"https://docs.google.com/document/d/13XWS_DVRG5FAkLbW7xZGqlwaRhTvDf_2WZvb6TJtheo/edit?usp=sharing",
	"https://docs.google.com/document/d/1mCvXEc3yM1wJTMF1jnewYa4lXGTLqfOtEQx3PMixaFg/edit?usp=sharing",
	"https://docs.google.com/document/d/1-0YYnhhV7sxP5Ws69_a-QhRXi0eFff0MgAGdEWAL0Bs/edit?usp=sharing",
	"https://docs.google.com/document/d/1hoOjom_uKNUcL26rvR0fYIuGlRdx6JCkY6Ow63UN0RA/edit?usp=sharing",
	"https://docs.google.com/document/d/1PNo0vgGRo8T8otTE1pmAe9GrS-8oNwF9CKLBdd89xNs/edit?usp=sharing",
	"https://docs.google.com/document/d/1znk012J-If9umgm6i3mjaIN19JfLik1Mzmiq6dh1Vr4/edit?usp=sharing",
}

func Challenges(c *fiber.Ctx, config *models.Configuration) error {
	// times := []time.Time{
	// 	config.StartTime,
	// 	utils.XAfter(time.Duration(30*time.Minute), config.StartTime),
	// 	utils.XAfter(time.Duration(60*time.Minute), config.StartTime),
	// 	utils.XAfter(time.Duration(90*time.Minute), config.StartTime),
	// 	utils.XAfter(time.Duration(120*time.Minute), config.StartTime),
	// 	utils.XAfter(time.Duration(150*time.Minute), config.StartTime),
	// 	utils.XAfter(time.Duration(180*time.Minute), config.StartTime),
	// 	utils.XAfter(time.Duration(225*time.Minute), config.StartTime),
	// }
	times := []time.Time{
		config.StartTime,
		utils.XAfter(time.Duration(1*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(2*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(3*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(4*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(5*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(6*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(7*time.Minute), config.StartTime),
	}
	for i, time := range times {
		if utils.IsBefore(time) {
			return c.JSON(challenges[:i+1])
		}
	}

	return c.JSON(challenges)
}
