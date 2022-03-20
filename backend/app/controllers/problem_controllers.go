package controllers

import (
	"servermodule/app/scraper"

	"github.com/gofiber/fiber/v2"
)

// Problems godoc
// @Summary Get a list of problems from the contest.
// @Description Returns an array of problem objects. Only available after contest has started.
// @Tags Contest Info
// @Produce  application/json
// @Success 200 {array} models.Problem
// @Failure 401
// @Router /problems [get]
func Problems(c *fiber.Ctx, client *scraper.Client) error {
	return c.JSON(client.Cached.Problems)
}

// ProblemPage godoc
// @Summary Get a problem statement from the contest.
// @Description Returns an html DOM subtree with the problem statement info. Only available after contest has started.
// @Tags Contest Info
// @Produce  application/json
// @Param problemId path string true "Problem ID"
// @Success 200 {string} string
// @Failure 401
// @Router /problem/{problemId} [get]
func ProblemPage(c *fiber.Ctx, client *scraper.Client) error {
	return c.JSON(client.Cached.ProblemPages[c.Params("problemId")])
}
