package models

/* UserEntity - User data model in Firebase. */
type UserEntity struct {
	Username    string
	Email       string
	Submissions map[string]*Submission
	Points      map[string]int
}
