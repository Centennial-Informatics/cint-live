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
	"/",
	"/",
}

func Challenges(c *fiber.Ctx, config *models.Configuration) error {
	times := []time.Time{
		config.StartTime,
		utils.XAfter(time.Duration(30*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(60*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(90*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(120*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(150*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(180*time.Minute), config.StartTime),
		utils.XAfter(time.Duration(225*time.Minute), config.StartTime),
	}
	for i, time := range times {
		if utils.IsBefore(time) {
			return c.JSON(challenges[:i+1])
		}
	}

	return c.JSON(challenges)
}
