package controllers

import (
	"servermodule/app/database"

	"github.com/gofiber/fiber/v2"
)

// Standings godoc
// @Summary Get the scores of all participants in the contest.
// @Description Returns an object of user objects.
// @Tags Contest Info
// @Produce  application/json
// @Success 200 {object} map[string]models.UserEntity
// @Router /standings [get]
func Standings(c *fiber.Ctx, db *database.ContestDB) error {
	standings := database.AggStandings(db)

	return c.JSON(standings)
}
