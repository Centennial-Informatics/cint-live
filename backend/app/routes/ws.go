package routes

import (
	"servermodule/app/controllers"
	"servermodule/app/database"
	"servermodule/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WsRoutes(app *fiber.App, config *models.Configuration,
	ts *models.TokenService,
	db *database.ContestDB) {
	app.Use("/ws", websocket.New(func(c *websocket.Conn) {
		controllers.WsController(c, db, ts)
	}))
}
