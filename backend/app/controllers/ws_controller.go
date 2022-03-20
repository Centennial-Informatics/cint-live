package controllers

import (
	"log"
	"servermodule/app/models"
	"time"

	"github.com/gofiber/websocket/v2"
)

func WsController(c *websocket.Conn, f *models.FirebaseService, ts *models.TokenService) {
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
				err = c.WriteJSON(f.Cache.Users.Get(userID).Submissions)

				if err != nil {
					c.Close()
					break
				}
			}
		} else {
			temp := f.Cache.Users.Get(userID).Submissions
			temp["A"].Time = time.Now().Unix()
			err = c.WriteJSON(temp)
			if err != nil {
				c.Close()
				break
			}
		}
	}
}
