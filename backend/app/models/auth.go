package models

import (
	"errors"
	"servermodule/utils"

	"github.com/gofiber/websocket/v2"
)

type AuthUser struct {
	ID     string
	TeamID string
	Conn   []*websocket.Conn
}

type AuthTeam struct {
	Conn []*websocket.Conn
}

type TokenService struct {
	user  map[string]*AuthUser
	team  map[string]*AuthTeam
	token map[string]string
}

func NewTokenService() *TokenService {
	return &TokenService{
		user:  map[string]*AuthUser{},
		team:  map[string]*AuthTeam{},
		token: map[string]string{},
	}
}

func (ts *TokenService) AuthorizeUser(token string) (string, error) {
	user, ok := ts.user[token]

	if !ok {
		return "", errors.New("Unauthorized")
	}

	return user.ID, nil
}

func (ts *TokenService) newTeamIfNecessary(teamID string) {
	if _, ok := ts.team[teamID]; !ok {
		ts.team[teamID] = &AuthTeam{
			Conn: make([]*websocket.Conn, 0),
		}
	}
}

func (ts *TokenService) UpdateToken(userID string, teamID string, length int) string {
	if _, ok := ts.token[userID]; ok {
		return ts.token[userID]
	}

	ts.token[userID] = utils.GenerateSecureToken(length)
	if _, ok := ts.user[ts.token[userID]]; !ok {
		ts.user[ts.token[userID]] = &AuthUser{
			ID:     userID,
			TeamID: teamID,
			Conn:   make([]*websocket.Conn, 0),
		}
	}

	ts.newTeamIfNecessary(teamID)

	return ts.token[userID]
}

func (ts *TokenService) SetConnection(token string, c interface{}) error {
	if _, ok := ts.user[token]; !ok {
		return errors.New("user does not exist")
	}
	ts.user[token].Conn = append(ts.user[token].Conn, c.(*websocket.Conn))
	ts.team[ts.user[token].TeamID].Conn = append(ts.team[ts.user[token].TeamID].Conn, c.(*websocket.Conn))

	return nil
}

func (ts *TokenService) ChangeConnections(userID string, teamID string) error {
	if _, ok := ts.token[userID]; !ok {
		return errors.New("user does not exist")
	}
	token := ts.token[userID]
	team := ts.team[ts.user[token].TeamID]

	// for efficiency
	userConns := map[(*websocket.Conn)]bool{}
	for _, conn := range ts.user[ts.token[userID]].Conn {
		userConns[conn] = true
	}

	// log.Println("old team len before:", len(team.Conn))
	// remove user conns from current team conn
	if team.Conn != nil {
		conns := team.Conn
		newConns := make([]*websocket.Conn, 0)
		for _, conn := range conns {
			if conn.Conn != nil {
				// if not in user conns
				if _, g := userConns[conn]; !g {
					newConns = append(newConns, conn)
				}
			}
		}

		team.Conn = newConns
	}
	// log.Println("old team len now:", len(team.Conn))

	ts.newTeamIfNecessary(teamID)
	ts.user[token].TeamID = teamID

	// log.Println("new team len before:", len(ts.team[teamID].Conn))

	// add user conns to new team conn
	ts.team[teamID].Conn = append(ts.team[teamID].Conn, ts.user[ts.token[userID]].Conn...)

	// log.Println("new team len after:", len(ts.team[teamID].Conn))

	return nil
}

func (ts *TokenService) GetUserFromToken(token string) *AuthUser {
	return ts.user[token]
}

func (ts *TokenService) GetUserFromID(userID string) *AuthUser {
	return ts.user[ts.token[userID]]
}

func (ts *TokenService) GetTeamFromUser(userID string) *AuthTeam {
	return ts.team[ts.GetUserFromID(userID).TeamID]
}

func (ts *TokenService) GetConnMetrics() map[string]int {
	metrics := map[string]int{}

	for teamID, team := range ts.team {
		metrics[teamID] = len(team.Conn)
	}

	return metrics
}
