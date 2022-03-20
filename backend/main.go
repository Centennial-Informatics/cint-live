package main

import (
	"encoding/json"
	"log"
	"os"
	"servermodule/app/models"
	"servermodule/app/routes"
	"servermodule/app/scraper"
	"servermodule/configs"
	"servermodule/utils"
	"servermodule/utils/jobs"
	"servermodule/utils/middleware"

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
	middleware.ScrapeWorker(appClient, config.ContestURL, config.ContestID, config.ScrapeIntervalAll)

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
	// initialize the firebase connection if present
	f, err := models.NewFirebaseService(config.FirebaseURL, appClient.Cached)
	if err != nil {
		log.Println(err)
		log.Println("Failed to establish firebase connection. Attempting to run without database.")
	}
	// write worker writes to firebase
	writeWorker, _ := middleware.WriteWorker(f, config.WriteInterval)
	ts := models.NewTokenService()
	// submit worker submits to codeforces and checks for verdicts, callback runs when verdict status is updated
	submitWorker, _ := middleware.SubmitWorker(clients, config.ScrapeIntervalVerdicts,
		func(submissionID string, verdict *models.Verdict) {
			log.Println("Submission", submissionID, "by", *verdict.UserID, "is", verdict.Verdict)
			f.Cache.Users.Get(*verdict.UserID).Submissions[*verdict.ProblemID].Status = verdict.Status
			f.Cache.Users.Get(*verdict.UserID).Submissions[*verdict.ProblemID].Verdict = verdict.Verdict

			user := ts.GetUserFromID(*verdict.UserID)

			var points int

			if verdict.Status == "Final" {
				points = utils.GetPoints(*verdict.ProblemID, verdict.Verdict, &config.Points)
				f.Cache.Users.Get(*verdict.UserID).Points[*verdict.ProblemID] = points
				f.Cache.Users.Get(*verdict.UserID).Submissions[*verdict.ProblemID].Points = points
			}
			/* if ws is open, live update submission status */
			if user.Conn != nil {
				conns := user.Conn
				newConns := make([]*websocket.Conn, 0)
				for _, conn := range conns {
					if conn.Conn != nil {
						newConns = append(newConns, conn)

						err = conn.WriteJSON(f.Cache.Users.Get(user.ID).Submissions)
						if err != nil {
							log.Println(err)
						}
					}
				}

				user.Conn = newConns
			}

			if verdict.Status == "Final" {
				writeWorker.AddJob(jobs.PointsJob{
					UserID:    *verdict.UserID,
					Points:    points,
					ProblemID: *verdict.ProblemID,
				})

				writeWorker.AddJob(jobs.SubmissionJob{
					UserID:    *verdict.UserID,
					ProblemID: *verdict.ProblemID,
					Submission: models.Submission{
						ID:      submissionID,
						Status:  verdict.Status,
						Verdict: verdict.Verdict,
						Time:    f.Cache.Users.Get(*verdict.UserID).Submissions[*verdict.ProblemID].Time,
						Points:  points,
					},
				})

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
	routes.PublicAPIRoutes(v1, config, appClient, writeWorker.F)
	routes.PublicTimeAPIRoutes(v1, config, appClient, writeWorker.F)
	routes.PrivateAPIRoutes(v1, config, ts, appClient, writeWorker.F, writeWorker)
	routes.PrivateTimeAPIRoutes(v1, config, ts, appClient, writeWorker.F, submitWorker, writeWorker)
	routes.WsRoutes(app, config, ts, appClient, writeWorker.F, writeWorker)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	middleware.MemLogger(config.LogInterval)

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
