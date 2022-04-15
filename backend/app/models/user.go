package models

/* UserEntity - User data model in Firebase. */
type TeamEntity struct {
	Name        string
	Submissions map[string]*Submission
	Points      map[string]int
	MemberCount int
}

/* UserEntity - User data model in Firebase. */
type UserEntity struct {
	Username string
	Email    string
	TeamID   string
	Division string
}
