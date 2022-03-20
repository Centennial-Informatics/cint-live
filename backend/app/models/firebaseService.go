package models

import (
	"context"
	"errors"
	"log"
	"servermodule/utils"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

/* FirebaseService - Represents a Firebase worker. */
type FirebaseService struct {
	App          *firebase.App
	DB           *db.Client
	Refs         map[string]*db.Ref
	Cache        *FirebaseCache
	ContestCache *ContestData
}

type UserMap struct {
	sync.Map
}

func (m *UserMap) Get(key string) *UserEntity {
	val, _ := m.Load(key)

	return val.(*UserEntity)
}

func (m *UserMap) Has(key string) bool {
	_, has := m.Load(key)

	return has
}

/* FirebaseCache - Cached Firebase data model, identical to actual data model. */
type FirebaseCache struct {
	Users *UserMap
}

/* NewFirebaseService - Initialize and return Firebase worker. */
func NewFirebaseService(url string, contestCache ...*ContestData) (*FirebaseService, error) {
	log.Println("Initializing Firebase")

	f := FirebaseService{}

	f.Refs = map[string]*db.Ref{}
	f.Cache = &FirebaseCache{}
	f.Cache.Users = &UserMap{}

	if len(contestCache) > 0 {
		f.ContestCache = contestCache[0]
	}

	app, err := utils.FirebaseApp()
	if err != nil {
		return &f, err
	}

	f.App = app

	f.DB, err = app.DatabaseWithURL(context.Background(), url)
	if err != nil {
		return &f, err
	}

	err = f.UpdateCache()

	return &f, err
}

/* NewUserEntity - create a new user entity and return it. */
func (f *FirebaseService) NewUserEntity(username string, email string) *UserEntity {
	if username == "" {
		username = email
	}

	u := UserEntity{
		Username: username,
		Email:    email,
	}
	s := map[string]*Submission{}
	p := map[string]int{}

	for _, prob := range f.ContestCache.Problems {
		s[prob.ID] = &Submission{
			Status:  "Final",
			Verdict: "Unsubmitted",
		}
		p[prob.ID] = 0
	}

	u.Submissions = s
	u.Points = p

	return &u
}

/* AddUser - Add new UserEntity to Firebase. */
func (f *FirebaseService) AddUser(userID string, userName string, email string) {
	if f.DB != nil { // still works w/o database
		f.addUser(userID)

		err := f.Refs[userID].Set(context.Background(), f.NewUserEntity(userName, email))
		if err != nil {
			log.Println(err)
		}
	}
}

/* DeleteUser - Delete user from Firebase. */
func (f *FirebaseService) DeleteUser(userID string) {
	if f.DB != nil {
		err := f.Refs[userID].Delete(context.Background())
		if err != nil {
			log.Println(err)
		}

		delete(f.Refs, userID)
	}
}

/* GetUser - Get UserEntity from Firebase. */
func (f *FirebaseService) GetUser(userID string) (*UserEntity, error) {
	if _, ok := f.Refs[userID]; !ok {
		f.addUser(userID)
	}

	ret := UserEntity{}

	err := f.Refs[userID].Get(context.Background(), &ret)
	if err != nil {
		return &ret, err
	}

	err = f.Refs[userID].Child("Submissions").Get(context.Background(), &ret.Submissions)
	if err != nil {
		return &ret, err
	}

	err = f.Refs[userID].Child("Points").Get(context.Background(), &ret.Points)
	if err != nil {
		return &ret, err
	}

	return &ret, err
}

func (f *FirebaseService) UpdateUser(userID string, username string, email string) error {
	if f.DB != nil {
		if _, ok := f.Refs[userID]; !ok {
			return errors.New("user does not exist")
		}

		if username == "" {
			username = email
		}

		user := UserEntity{}

		err := f.Refs[userID].Get(context.Background(), &user)
		if err != nil {
			return err
		}

		user.Username = username
		user.Email = email

		err = f.Refs[userID].Set(context.Background(), user)

		return err
	}

	return nil
}

/* UpdateSubmission - Updates a user's submission for a problem. */
func (f *FirebaseService) UpdateSubmission(userID string, problemID string, submission Submission) {
	if f.DB != nil {
		err := f.Refs[userID].Child("Submissions").Update(context.Background(), map[string]interface{}{
			problemID: submission,
		})

		if err != nil {
			log.Println(err)
		}
	}
}

/* UpdatePoints - Updates a user's points for a problem. */
func (f *FirebaseService) UpdatePoints(userID string, problemID string, points int) {
	if f.DB != nil {
		err := f.Refs[userID].Child("Points").Update(context.Background(), map[string]interface{}{
			problemID: points,
		})

		if err != nil {
			log.Println(err)
		}
	}
}

/* UpdateCache - Updates the memory cache. */
func (f *FirebaseService) UpdateCache() error {
	temp := map[string]*UserEntity{}
	if f.DB != nil {
		err := f.DB.NewRef("/users").Get(context.Background(), &temp)
		if err != nil {
			return err
		}

		for ID, user := range temp {
			f.Cache.Users.Map.Store(ID, user)
			f.Refs[ID] = f.DB.NewRef("/users/" + ID)
		}
	}

	return nil
}

func (f *FirebaseService) addUser(userID string) {
	f.Refs[userID] = f.DB.NewRef("/users/" + userID)

	f.DB.NewRef("/users/" + userID + "/Submissions")
	f.DB.NewRef("/users/" + userID + "/Points")
}
