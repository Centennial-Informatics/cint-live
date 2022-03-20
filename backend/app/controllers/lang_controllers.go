package controllers

import (
	"servermodule/app/scraper"

	"github.com/gofiber/fiber/v2"
)

// Languages godoc
// @Summary Get valid languages for the contest.
// @Description Returns an array of JSON objects representing valid languages.
// @Tags Contest Info
// @Produce  application/json
// @Success 200 {array} models.Lang
// @Router /languages [get]
func Languages(c *fiber.Ctx, client *scraper.Client) error {
	return c.JSON(client.Languages())
}
