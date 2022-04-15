package controllers

import (
	"log"
	"servermodule/app/database"
	"servermodule/app/models"

	"github.com/gofiber/websocket/v2"
)

func WsController(c *websocket.Conn, db *database.ContestDB, ts *models.TokenService) {
	auth := false
	userID := ""

	var (
		msg []byte
		err error
	)

	for {
		if _, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}

		if !auth {
			token := string(msg)
			log.Printf("recv: %s", msg)

			err = ts.SetConnection(token, c)
			if err != nil {
				log.Println("no message")
			} else {
				//log.Println(*ts.GetUser(token))
				auth = true
				userID = ts.GetUserFromToken(token).ID
				_, _, submissions := db.GetTeamByCode(db.GetUser(userID).TeamCode)
				err = c.WriteJSON(submissions)

				if err != nil {
					c.Close()
					break
				}
			}
		} else {
			_, _, submissions := db.GetTeamByCode(db.GetUser(userID).TeamCode)
			err = c.WriteJSON(submissions)
			if err != nil {
				c.Close()
				break
			}
		}
	}
}
