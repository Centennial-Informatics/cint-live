package models

import "time"

/* Configuration — fully parsed configuration object */
type Configuration struct {
	AuthTokenLength        int
	ContestID              string
	ContestURL             string
	PointsPolicy           string
	Points                 map[string]int
	FirebaseURL            string
	LogInterval            time.Duration
	ScrapeIntervalAll      time.Duration
	ScrapeIntervalVerdicts time.Duration
	StartTime              time.Time
	StopTime               time.Time
	WriteInterval          time.Duration
}
