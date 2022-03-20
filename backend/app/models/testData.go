package models

type TestDataConfiguration struct {
	ContestID   string
	ContestURL  string
	Info        ContestInfo
	Problems    []Problem
	ProblemID   string
	Submission  string
	Language    string
	FirebaseURL string
}
