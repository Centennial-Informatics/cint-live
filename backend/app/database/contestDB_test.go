package database_test

import (
	"servermodule/app/database"
	"strconv"
	"strings"
	"testing"
)

// clear the test db
func newTestDB(t *testing.T) *database.ContestDB {
	db, err := database.NewEmptyDB("test.db")
	if err != nil {
		t.Error(err)
	}

	return db
}

func TestUser(t *testing.T) {
	db := newTestDB(t)

	const testEmail = "user@test.edu"
	const testUsername = "test user"
	const testDivision = "test division"

	if db.HasUser(testEmail) {
		t.Fail()
	}

	user := db.CreateUser(testUsername, testEmail, testDivision)

	if user.Username != testUsername {
		t.Fail()
	}

	if !db.HasUser(testEmail) {
		t.Fail()
	}

	if db.GetUser(testEmail).Username != testUsername {
		t.Fail()
	}

	// user cant be made twice
	duplicate := db.CreateUser(testUsername, testEmail, testDivision)
	if duplicate.TeamCode != user.TeamCode {
		t.Fail()
	}

	const newTestUsername = "new username"

	db.UpdateUser(testEmail, newTestUsername, testDivision)

	user = db.GetUser(testEmail)
	if user.Username != newTestUsername {
		t.Fail()
	}
}

func TestTeam(t *testing.T) {
	db := newTestDB(t)

	const testName = "team name"
	team := db.CreateTeam(testName, "Standard")

	testTeam, users := db.GetTeamByCode(team.Code)
	if testTeam.Name != testName {
		t.Fail()
	}

	if len(users) > 0 {
		t.Fail()
	}

	testTeam2, _ := db.GetTeam(testTeam.ID)
	if testTeam2.Name != testName {
		t.Fail()
	}

	const newTestName = "new team name"

	db.UpdateTeam(team.Code, newTestName)

	testTeam, users = db.GetTeamByCode(team.Code)
	if testTeam.Name != newTestName {
		t.Fail()
	}

	if len(users) > 0 {
		t.Fail()
	}

	user := db.CreateUser("test", "test", "advanced")
	testTeam, _ = db.GetTeamByCode(user.TeamCode)
	if testTeam.Division != "Advanced" {
		t.Fail()
	}
	testTeam, _ = db.GetTeam(testTeam.ID)
	if testTeam.Division != "Advanced" {
		t.Fail()
	}

	db.DeleteTeam(testTeam.Code)
	_, members := db.GetTeamByCode(testTeam.Code)
	if members != nil {
		t.Fail()
	}
}

func TestUserTeamAssociation(t *testing.T) {
	db := newTestDB(t)

	user := db.CreateUser("user", "email", "division")

	initialTeam, _ := db.GetTeamByCode(user.TeamCode)
	if initialTeam.Name != user.Username {
		t.Fail()
	}

	newTeam := db.CreateTeam("new team", "Standard")

	// user is assigned to new team
	newUserTeam := db.UpdateUserTeam(user.Email, newTeam.Code)
	if newTeam.Code != newUserTeam.Code {
		t.Fail()
	}

	if newUserTeam.ID == initialTeam.ID {
		t.Fail()
	}

	user = db.GetUser(user.Email)
	if user.TeamCode != newTeam.Code {
		t.Fail()
	}

	_, users := db.GetTeamByCode(user.TeamCode)
	if users[0].ID != user.ID {
		t.Fail()
	}

	// initial team is deleted
	_, members := db.GetTeamByCode(initialTeam.Code)
	if members != nil {
		t.Fail()
	}

	diffTeam := db.CreateTeam("different team", "Standard")

	// user is re-assigned to original team
	newTeam = db.UpdateUserTeam(user.Email, diffTeam.Code)

	_, users = db.GetTeamByCode(newTeam.Code)
	if users[0].ID != user.ID {
		t.Fail()
	}

	// initial team is deleted
	_, members = db.GetTeamByCode(newUserTeam.Code)
	if members != nil {
		t.Fail()
	}

	// user is assigned to an invalid team
	newTeam = db.UpdateUserTeam(user.Email, "123")

	_, users = db.GetTeamByCode(newTeam.Code)
	if users[0].ID != user.ID {
		t.Fail()
	}

	if newTeam.Code != diffTeam.Code {
		t.Fail()
	}

	// leaving a team
	user2 := db.CreateUser("user 2", "email 2", "Standard")
	db.UpdateUserTeam(user2.Email, newTeam.Code)

	db.LeaveTeam(user.Email)
	_, users = db.GetTeamByCode(newTeam.Code)
	if len(users) != 1 {
		t.Fail()
	}
	if users[0].ID != user2.ID {
		t.Fail()
	}

	user = db.GetUser(user.Email)
	_, users = db.GetTeamByCode(user.TeamCode)
	if len(users) != 1 {
		t.Fail()
	}
	if users[0].ID != user.ID {
		t.Fail()
	}
}

func TestSubmission(t *testing.T) {
	db := newTestDB(t)

	team := db.CreateTeam("test submission team", "Standard")

	const testSubID = "submission id"
	const testProbID = "problem id"
	const testTimestamp = 1
	const testPoints = 0
	const testVerdict = "Pending"
	const testStatus = "Pending"

	submission := db.CreateSubmission(testSubID, testProbID, team.Code, testTimestamp)
	if submission.SubmissionID != testSubID {
		t.Fail()
	}
	if submission.ProblemID != testProbID {
		t.Fail()
	}
	if submission.TeamCode != team.Code {
		t.Fail()
	}
	if submission.Time != testTimestamp {
		t.Fail()
	}
	if submission.Points != testPoints {
		t.Fail()
	}
	if submission.Verdict != testVerdict {
		t.Fail()
	}
	if submission.Status != testStatus {
		t.Fail()
	}

	newTestVerdict := "Good"
	newTestStatus := "Final"
	newTestPoints := 10
	db.UpdateSubmission(submission.SubmissionID, newTestPoints, newTestVerdict, newTestStatus)
	submission = db.GetSubmission(submission.SubmissionID)
	if submission.Verdict != newTestVerdict {
		t.Fail()
	}
	if submission.Status != newTestStatus {
		t.Fail()
	}
	if submission.Points != newTestPoints {
		t.Fail()
	}

	const newSubmissionID = "1234"
	newTestPoints = 20
	newTestVerdict = "Very good"
	newTestStatus = "More final"
	db.UpdateSubmissionByID(submission.ID, newSubmissionID, newTestPoints, newTestVerdict, newTestStatus)
	submission = db.GetSubmission(newSubmissionID)
	if submission.Verdict != newTestVerdict {
		t.Fail()
	}
	if submission.Status != newTestStatus {
		t.Fail()
	}
	if submission.Points != newTestPoints {
		t.Fail()
	}
	if submission.SubmissionID != newSubmissionID {
		t.Fail()
	}
}

func TestSubmissionTeamAssociation(t *testing.T) {
	db := newTestDB(t)

	const testSubID = "submission id"
	const testProbID = "problem id"
	const testTimestamp = 1

	team := db.CreateTeam("test submission team", "Standard")

	submission := db.CreateSubmission(testSubID, testProbID, team.Code, testTimestamp)

	subTeam, _ := db.GetTeamByCode(submission.TeamCode)
	submissions := db.GetTeamSubmissions(submission.TeamCode)
	if submissions[0].SubmissionID != testSubID {
		t.Fail()
	}
	if subTeam.Code != team.Code {
		t.Fail()
	}
}

func fillTestUsers(db *database.ContestDB, n int) []database.User {
	users := make([]database.User, 0)
	for i := 1; i <= n; i++ {
		users = append(users, *db.CreateUser("User "+strconv.Itoa(i), strconv.Itoa(i), "Standard"))
	}

	return users
}

func fillTestSubmissions(db *database.ContestDB, users []database.User) []database.Submission {
	submissions := make([]database.Submission, 0)
	for i, user := range users {
		db.CreateSubmission("submission A"+strconv.Itoa(i+1), "A", user.TeamCode, 1)
		db.CreateSubmission("submission B"+strconv.Itoa(i+1), "B", user.TeamCode, 2)
		db.CreateSubmission("submission 2A"+strconv.Itoa(i+1), "A", user.TeamCode, 3)
	}

	return submissions
}

func TestAggStandings(t *testing.T) {
	db := newTestDB(t)

	users := fillTestUsers(db, 10)
	used := map[string]bool{}

	for _, user := range users {
		used[user.Username] = false
	}

	standings := database.AggStandings(db)
	for _, team := range standings {
		for _, user := range team.Members {
			if used[user] {
				t.Fail()
			} else {
				used[user] = true
			}
		}
	}

	for _, user := range users {
		if !used[user.Username] {
			t.Fail()
		}
	}

	submissions := fillTestSubmissions(db, users)
	for _, submission := range submissions {
		used[submission.SubmissionID] = false
	}

	standings = database.AggStandings(db)
	for _, team := range standings {
		if len(team.Verdicts) != 2 {
			t.Fail()
		}
		if !strings.Contains(team.Verdicts["A"].SubmissionID, "2A") {
			t.Fail()
		}
		for _, submission := range team.Verdicts {
			if used[submission.SubmissionID] {
				t.Fail()
			} else {
				used[submission.SubmissionID] = true
			}
		}
	}

	for _, submission := range submissions {
		if strings.Contains(submission.SubmissionID, "2A") && !used[submission.SubmissionID] {
			t.Fail()
		} else if strings.Contains(submission.SubmissionID, " A") && used[submission.SubmissionID] {
			t.Fail()
		} else if !used[submission.SubmissionID] {
			t.Fail()
		}
	}

	t.Log(standings)
}
