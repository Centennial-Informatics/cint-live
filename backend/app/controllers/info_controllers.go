package controllers

import (
	"servermodule/app/models"
	"servermodule/app/scraper"

	"github.com/gofiber/fiber/v2"
)

// Schedule godoc
// @Summary Get the contest schedule.
// @Description Returns a JSON object with the start and stop times for the contest.
// @Tags Contest Info
// @Produce  text/plain
// @Success 200 {object} models.Schedule
// @Router /schedule [get]
func Schedule(c *fiber.Ctx, config *models.Configuration) error {
	return c.JSON(models.Schedule{
		Start: config.StartTime.Unix(),
		Stop:  config.StopTime.Unix(),
	})
}

func About(c *fiber.Ctx, client *scraper.Client) error {
	return c.JSON(client.Cached.Info)
}
