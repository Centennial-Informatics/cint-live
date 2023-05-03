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
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

func LiveServer(standardConfig *models.Configuration, advancedConfig *models.Configuration, practiceConfig *models.Configuration) {
	var accounts []*models.CFAccount
	// initialize the service accounts
	err := json.Unmarshal([]byte(os.Getenv("CF_ACCOUNTS")), &accounts)
	if err != nil {
		log.Fatal(err)
	}

	appClient, err := scraper.NewClient(accounts[0].User+" standard", accounts[0].User, accounts[0].Pass)
	if err != nil {
		log.Fatal(err)
	}

	appClient2, err := scraper.NewClient(accounts[0].User+" advanced", accounts[0].User, accounts[0].Pass)
	if err != nil {
		log.Fatal(err)
	}

	appClient3, err := scraper.NewClient(accounts[0].User+" practice", accounts[0].User, accounts[0].Pass)
	if err != nil {
		log.Fatal(err)
	}
	// scrape worker scrapes problem statements from codeforces
	workers.ScrapeWorker(appClient, standardConfig.ContestURL, standardConfig.ContestID, standardConfig.ScrapeIntervalAll)
	workers.ScrapeWorker(appClient2, advancedConfig.ContestURL, advancedConfig.ContestID, advancedConfig.ScrapeIntervalAll)
	workers.ScrapeWorker(appClient3, practiceConfig.ContestURL, practiceConfig.ContestID, practiceConfig.ScrapeIntervalAll)

	clients := make([]*scraper.Client, 0)

	clients = append(clients, appClient)
	// clients = append(clients, appClient2)
	// clients = append(clients, appClient3)

	for i, acc := range accounts[1:] {
		c, err := scraper.NewClient(acc.User, acc.User, acc.Pass)
		if err != nil {
			log.Fatal(err)
		}

		if i%2 == 0 {
			c.CopyCache(appClient.Cached)
		} else {
			c.CopyCache(appClient2.Cached)
		}
		clients = append(clients, c)
	}

	db, err := database.NewPostgresDB(os.Getenv("DATABASE_URL"), appClient.Cached, appClient2.Cached)
	if err != nil {
		log.Fatal(err)
	}

	workers.StandingsWorker(db, standardConfig.ScrapeIntervalVerdicts, standardConfig)

	ts := models.NewTokenService()
	// submit worker submits to codeforces and checks for verdicts, callback runs when verdict status is updated
	submitWorker, _ := workers.SubmitWorker(clients, standardConfig.ScrapeIntervalVerdicts,
		func(submissionID string, verdict *models.Verdict) {
			log.Println("Submission", submissionID, "by", *verdict.UserID, "is", verdict.Verdict)
			teamCode := db.GetUser(*verdict.UserID).TeamCode

			team := ts.GetTeamFromUser(*verdict.UserID)

			var points int

			if verdict.Status == "Final" {
				points = utils.GetPoints(*verdict.ProblemID, verdict.Verdict, &standardConfig.Points)
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

	// app.Static("/", "../frontend/build")
	app.Use(logger.New())
	app.Use(cors.New(
		cors.Config{},
	))

	var ADMIN_ACCOUNTS []string

	err = json.Unmarshal([]byte(os.Getenv("ADMIN_ACCOUNTS")), &ADMIN_ACCOUNTS)
	if err != nil {
		log.Fatal(err)
	}

	routes.PublicRoutes(app)
	routes.PublicAPIRoutes(v1, standardConfig, appClient)
	routes.PublicTimeAPIRoutes(v1, standardConfig, appClient, appClient2, appClient3, db)
	routes.PrivateAPIRoutes(v1, standardConfig, ts, appClient, db)
	routes.PrivateTimeAPIRoutes(v1, standardConfig, ts, appClient, appClient2, appClient3, submitWorker, db)
	routes.WsRoutes(app, standardConfig, ts, db)
	routes.AdminAPIRoutes(v1, ADMIN_ACCOUNTS, standardConfig, ts, db, []*scraper.Client{appClient, appClient2})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	workers.MemLogger(standardConfig.LogInterval)

	log.Fatal(app.Listen(":" + port))
}

// @title CCC Live Server API
// @version 1.0
// @description A live competition server for coding competitions built and maintained by CCC.

// @contact.name Developers
// @contact.email chscompcode@gmail.com

// @BasePath /api/v1

func main() {
	standardConfig, err := configs.Cint2023()
	if err != nil {
		log.Fatal(err)
	}

	advancedConfig, err := configs.CintAdv()
	if err != nil {
		log.Fatal(err)
	}

	practiceConfig, err := configs.Practice()
	if err != nil {
		log.Fatal(err)
	}

	LiveServer(standardConfig, advancedConfig, practiceConfig)
}
