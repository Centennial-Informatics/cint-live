package controllers

import (
	"servermodule/app/models"

	"github.com/gofiber/fiber/v2"
)

// Standings godoc
// @Summary Get the scores of all participants in the contest.
// @Description Returns an object of user objects.
// @Tags Contest Info
// @Produce  application/json
// @Success 200 {object} map[string]models.UserEntity
// @Router /standings [get]
func Standings(c *fiber.Ctx, fc *models.FirebaseCache) error {
	users := map[string]*models.UserEntity{}

	fc.Users.Range(func(key, value interface{}) bool {
		users[key.(string)] = value.(*models.UserEntity)

		return true
	})

	return c.JSON(users)
}
