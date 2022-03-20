package routes

import (
	"servermodule/app/controllers"
	"servermodule/app/models"
	"servermodule/app/scraper"
	"servermodule/utils/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WsRoutes(app *fiber.App, config *models.Configuration,
	ts *models.TokenService, client *scraper.Client,
	f *models.FirebaseService, w *middleware.Writer) {
	app.Use("/ws", websocket.New(func(c *websocket.Conn) {
		controllers.WsController(c, f, ts)
	}))
}
