package main

import (
	"encoding/json"
	"log"
	"os"
	"servermodule/app/database"
	"servermodule/app/models"
	"servermodule/app/routes"
	"servermodule/app/scraper"
	"servermodule/configs"
	"servermodule/utils"
	"servermodule/utils/workers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

func LiveServer(config *models.Configuration) {
	var accounts []*models.CFAccount
	// initialize the service accounts
	err := json.Unmarshal([]byte(os.Getenv("CF_ACCOUNTS")), &accounts)
	if err != nil {
		log.Fatal(err)
	}

	appClient, err := scraper.NewClient(accounts[0].User, accounts[0].User, accounts[0].Pass)
	if err != nil {
		log.Fatal(err)
	}
	// scrape worker scrapes problem statements from codeforces
	workers.ScrapeWorker(appClient, config.ContestURL, config.ContestID, config.ScrapeIntervalAll)

	clients := make([]*scraper.Client, 0)

	clients = append(clients, appClient)

	for _, acc := range accounts[1:] {
		c, err := scraper.NewClient(acc.User, acc.User, acc.Pass)
		if err != nil {
			log.Fatal(err)
		}

		c.CopyCache(appClient.Cached)
		clients = append(clients, c)
	}

	db, err := database.NewDB("data/contest_data.db")
	if err != nil {
		log.Fatal(err)
	}
	ts := models.NewTokenService()
	// submit worker submits to codeforces and checks for verdicts, callback runs when verdict status is updated
	submitWorker, _ := workers.SubmitWorker(clients, config.ScrapeIntervalVerdicts,
		func(submissionID string, verdict *models.Verdict) {
			log.Println("Submission", submissionID, "by", *verdict.UserID, "is", verdict.Verdict)
			teamCode := db.GetUser(*verdict.UserID).TeamCode

			team := ts.GetTeamFromUser(*verdict.UserID)

			var points int

			if verdict.Status == "Final" {
				points = utils.GetPoints(*verdict.ProblemID, verdict.Verdict, &config.Points)
			}
			db.UpdateSubmissionByID(verdict.SubmissionID, submissionID, points, verdict.Verdict, verdict.Status)
			/* if ws is open, live update submission status */
			if team.Conn != nil {
				conns := team.Conn
				newConns := make([]*websocket.Conn, 0)
				for _, conn := range conns {
					if conn.Conn != nil {
						newConns = append(newConns, conn)
						submissions := db.GetTeamSubmissions(teamCode)
						err = conn.WriteJSON(submissions)
						if err != nil {
							log.Println(err)
						}
					}
				}

				team.Conn = newConns
			}

			if verdict.Status == "Final" {
				for _, client := range clients {
					delete(client.Verdict, submissionID)
				}
			}
		})

	app := fiber.New()
	api := app.Group("/api") // baseurl/api
	v1 := api.Group("/v1")   // baseurl/api/v1

	app.Static("/", "./docs")
	app.Use(logger.New())

	routes.PublicRoutes(app)
	routes.PublicAPIRoutes(v1, config, appClient)
	routes.PublicTimeAPIRoutes(v1, config, appClient, db)
	routes.PrivateAPIRoutes(v1, config, ts, appClient, db)
	routes.PrivateTimeAPIRoutes(v1, config, ts, appClient, submitWorker, db)
	routes.WsRoutes(app, config, ts, db)
	routes.AdminAPIRoutes(v1, config, db)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	workers.MemLogger(config.LogInterval)

	log.Fatal(app.Listen(":" + port))
}

// @title CCC Live Server API
// @version 1.0
// @description A live competition server for coding competitions built and maintained by CCC.

// @contact.name Developers
// @contact.email chscompcode@gmail.com

// @BasePath /api/v1

func main() {
	config, err := configs.Cint()
	if err != nil {
		log.Fatal(err)
	}

	LiveServer(config)
}
