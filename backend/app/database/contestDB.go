package database

import (
	"servermodule/app/models"
	"servermodule/utils"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Username string `json:"username"`
	Email    string `json:"email"`
	TeamCode string `json:"team_code"`
	Division string `json:"division"`
}

type Team struct {
	ID       int
	Code     string `json:"code"`
	Name     string `json:"name"`
	Division string `json:"division"`
}

type Submission struct {
	ID           int
	SubmissionID string `json:"submission_id"`
	ProblemID    string `json:"problem_id"`
	TeamCode     string `json:"team_code"`
	Points       int    `json:"points"`
	Verdict      string `json:"verdict"`
	Status       string `json:"status"`
	Time         int64  `json:"time"`
}

// publicly visible
type StandingsEntry struct {
	TeamID   int                   `json:"team_id"`
	TeamName string                `json:"team_name"`
	Verdicts map[string]Submission `json:"verdicts"`
	Members  []string              `json:"members"`
	Division string                `json:"division"`
}

type ContestDB struct {
	db            *gorm.DB
	ProblemsCache *models.ContestData
}

func NewDB(path string) (*ContestDB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrate schema
	db.AutoMigrate(&User{}, &Team{}, &Submission{})

	return &ContestDB{db: db}, err
}

func NewPostgresDB(dsn string) (*ContestDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrate schema
	db.AutoMigrate(&User{}, &Team{}, &Submission{})

	return &ContestDB{db: db}, err
}

func NewEmptyDB(path string) (*ContestDB, error) {
	db, err := NewDB(path)
	if err != nil {
		return db, err
	}

	db.Clear()
	return NewDB(path)
}

func (db *ContestDB) Clear() {
	db.db.Migrator().DropTable(&User{})
	db.db.Migrator().DropTable(&Team{})
	db.db.Migrator().DropTable(&Submission{})
}

// authorize email after google form response
func (db *ContestDB) CreateUser(username string, email string, division string) *User {
	var user User
	if !db.HasUser(email) {
		user = User{
			Username: username,
			Email:    email,
			TeamCode: db.CreateTeam(username, division).Code,
			Division: division,
		}
		db.db.Create(&user)
	} else {
		db.UpdateUser(email, username, division)
		user = *db.GetUser(email)
	}

	return &user
}

func (db *ContestDB) UpdateUser(email string, username string, division string) {
	db.db.Model(&User{}).Where("email = ?", email).Updates(User{
		Username: username,
		Division: division,
	})
}

func (db *ContestDB) HasUser(email string) bool {
	var user []User
	db.db.Where("email = ?", email).Find(&user)
	return len(user) > 0
}

// for user only after gmail authentication
func (db *ContestDB) GetUser(email string) *User {
	var user User

	db.db.First(&user, "email = ?", email)

	return &user
}

// for user only after gmail authentication, before contest starts
func (db *ContestDB) UpdateUserTeam(email string, teamCode string) *Team {
	var team Team
	user := db.GetUser(email)

	if db.HasTeam(teamCode) {
		// delete old team if no members left
		members := db.GetTeamMembers(user.TeamCode)
		if len(members) == 1 {
			db.DeleteTeam(user.TeamCode)
		}
		db.db.Model(user).Update("team_code", teamCode)
		team, _ = db.GetTeamByCode(teamCode)
	} else {
		db.db.First(&team, "code = ?", user.TeamCode)
	}

	return &team
}

// leaves existing team (if more than 1 member) and forms new team
func (db *ContestDB) LeaveTeam(email string) {
	user := db.GetUser(email)
	_, members := db.GetTeamByCode(user.TeamCode)
	if len(members) > 1 {
		newTeam := db.CreateTeam(user.Username, user.Division)
		db.UpdateUserTeam(email, newTeam.Code)
	}
}

// 1 team created with each user, users can then join other user teams
func (db *ContestDB) CreateTeam(name string, division string) *Team {
	team := Team{
		Code:     utils.GenerateSecureToken(3),
		Name:     name,
		Division: division,
	}

	db.db.Create(&team)

	return &team
}

func (db *ContestDB) UpdateTeam(code string, name string) {
	db.db.Model(&Team{}).Where("code = ?", code).Update("name", name)
}

func (db *ContestDB) HasTeam(code string) bool {
	var team []Team
	db.db.Where("code = ?", code).Find(&team)
	return len(team) > 0
}

func (db *ContestDB) DeleteTeam(code string) {
	db.db.Where("code = ?", code).Delete(&Team{})
}

// for team members only
func (db *ContestDB) GetTeamByCode(code string) (Team, []User) {
	if !db.HasTeam(code) {
		return Team{}, nil
	}
	var team Team
	db.db.First(&team, "code = ?", code)

	members := db.GetTeamMembers(code)

	team.Division = "Standard"
	for _, member := range members {
		if member.Division != "Standard" {
			team.Division = "Advanced"
		}
	}

	db.db.Model(&Team{}).Where("code = ?", code).Update("division", team.Division)

	return team, members
}

func (db *ContestDB) GetTeam(ID int) (Team, []User) {
	var team Team
	db.db.First(&team, ID)

	members := db.GetTeamMembers(team.Code)

	team.Division = "Standard"
	for _, member := range members {
		if member.Division != "Standard" {
			team.Division = "Advanced"
		}
	}

	db.db.First(&team, ID).Update("division", team.Division)

	return team, members
}

func (db *ContestDB) GetTeamMembers(code string) []User {
	var members []User
	db.db.Where("team_code = ?", code).Find(&members)
	return members
}

func (db *ContestDB) GetTeamSubmissions(code string) []Submission {
	var submissions []Submission
	db.db.Where("team_code = ?", code).Find(&submissions)
	return submissions
}

func (db *ContestDB) CreateSubmission(submissionID string, problemID string, teamCode string, timestamp int64) *Submission {
	submission := Submission{
		SubmissionID: submissionID,
		ProblemID:    problemID,
		TeamCode:     teamCode,
		Points:       0,
		Verdict:      "Pending",
		Status:       "Pending",
		Time:         timestamp,
	}

	db.db.Create(&submission)

	return &submission
}

func (db *ContestDB) UpdateSubmission(submissionID string, points int, verdict string, status string) {
	db.db.Model(&Submission{}).Where("submission_id = ?", submissionID).Updates(Submission{
		Points:  points,
		Verdict: verdict,
		Status:  status,
	})
}

func (db *ContestDB) UpdateSubmissionByID(ID int, submissionID string, points int, verdict string, status string) {
	db.db.Model(&Submission{}).Where(ID).Updates(Submission{
		SubmissionID: submissionID,
		Points:       points,
		Verdict:      verdict,
		Status:       status,
	})
}

func (db *ContestDB) GetSubmission(submissionID string) *Submission {
	var submission Submission
	db.db.First(&submission, "submission_id = ?", submissionID)

	return &submission
}

func AggStandings(db *ContestDB) []StandingsEntry {
	var teams []Team
	var members []User
	var submissions []Submission
	db.db.Find(&teams)

	standings := make([]StandingsEntry, 0)

	for _, team := range teams {
		db.db.Where("team_code = ?", team.Code).Find(&members)
		if len(members) > 0 {
			entry := StandingsEntry{
				TeamID:   team.ID,
				TeamName: team.Name,
				Verdicts: map[string]Submission{},
				Members:  make([]string, 0),
				Division: team.Division, // TODO: make this portable
			}

			for _, member := range members {
				entry.Members = append(entry.Members, member.Username)
			}

			db.db.Where("team_code = ?", team.Code).Find(&submissions)

			for _, submission := range submissions {
				submission.TeamCode = "REDACTED"
				if verdict, exists := entry.Verdicts[submission.ProblemID]; exists {
					if verdict.Time < submission.Time {
						entry.Verdicts[submission.ProblemID] = submission
					}
				} else {
					entry.Verdicts[submission.ProblemID] = submission
				}
			}

			standings = append(standings, entry)
		}
	}

	return standings
}
