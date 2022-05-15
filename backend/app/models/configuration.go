package models

import "time"

/* Configuration — fully parsed configuration object */
type Configuration struct {
	AuthTokenLength        int
	ContestID              string
	ContestURL             string
	MaxTeamSize            int
	PointsPolicy           string
	Points                 map[string]int
	FirebaseURL            string
	LogInterval            time.Duration
	ScoreboardFreeze       time.Duration
	ScoreboardUnfreeze     time.Duration
	ScrapeIntervalAll      time.Duration
	ScrapeIntervalVerdicts time.Duration
	StartTime              time.Time
	StopTime               time.Time
	SubmissionGrace        time.Duration
	WriteInterval          time.Duration
	Standard               []string // only 2 divisions for now
	Advanced               []string
}
