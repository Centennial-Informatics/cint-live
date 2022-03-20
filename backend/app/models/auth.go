package models

import (
	"errors"
	"servermodule/utils"

	"github.com/gofiber/websocket/v2"
)

type AuthUser struct {
	ID   string
	Conn []*websocket.Conn
}

type TokenService struct {
	user  map[string]*AuthUser
	token map[string]string
}

func NewTokenService() *TokenService {
	return &TokenService{
		user:  map[string]*AuthUser{},
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

func (ts *TokenService) UpdateToken(userID string, length int) string {
	if _, ok := ts.token[userID]; ok {
		return ts.token[userID]
	}

	ts.token[userID] = utils.GenerateSecureToken(length)
	if _, ok := ts.user[ts.token[userID]]; !ok {
		ts.user[ts.token[userID]] = &AuthUser{
			ID:   userID,
			Conn: make([]*websocket.Conn, 0),
		}
	}

	return ts.token[userID]
}

func (ts *TokenService) SetConnection(token string, c interface{}) error {
	if _, ok := ts.user[token]; !ok {
		return errors.New("user does not exist")
	}

	ts.user[token].Conn = append(ts.user[token].Conn, c.(*websocket.Conn))

	return nil
}

func (ts *TokenService) GetUserFromToken(token string) *AuthUser {
	return ts.user[token]
}

func (ts *TokenService) GetUserFromID(userID string) *AuthUser {
	return ts.user[ts.token[userID]]
}
