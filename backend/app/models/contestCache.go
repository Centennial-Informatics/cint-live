package models

import (
	"mime/multipart"
	"time"
)

type CFAccount struct {
	User string
	Pass string
}

type ContestData struct {
	ContestID    string
	ContestURL   string
	Problems     []Problem
	ProblemPages map[string]ProblemPage
	Info         ContestInfo
}

type Schedule struct {
	Start int64
	Stop  int64
}

type Lang struct {
	ID   string
	Name string
	Ext  string
}

type Problem struct {
	ID   string
	Name string
}

type Submission struct {
	ID      string
	Verdict string
	Status  string
	Time    int64
	Points  int
}

type QueuedSubmission struct {
	UserID       string
	SubmissionID int
	Source       string
	Language     string
	ProblemID    string
	File         *multipart.FileHeader
	ContestURL   string
}

type Verdict struct {
	Status, Verdict string
	UserID          *string
	SubmissionID    int
	ProblemID       *string
}

type ProblemStatus struct {
	Status string
	Points int
	Time   string
}

type StandingsEntity struct {
	TotalPoints int
	Status      map[string]*ProblemStatus
}

type Standings struct {
	Team map[string]*StandingsEntity
}

type ProblemPage struct {
	Header              ProblemHeader
	Statement           []ProblemStatementParagraph
	InputSpecification  []ProblemStatementParagraph
	OutputSpecification []ProblemStatementParagraph
	SampleTests         []ProblemSampleTest
	Note                []ProblemStatementParagraph
}

type ProblemSampleTest struct {
	Input  string
	Output string
}

type ProblemStatementParagraph struct {
	Text  string
	Code  string
	Image string
}

type ProblemHeader struct {
	Title     string
	TimeLimit string
	MemLimit  string
	Input     string
	Output    string
}

type SubmissionCallback func(submissionID string, verdict *Verdict)

type Page struct {
	Info         ContestInfo
	Problems     []Problem
	ProblemPages map[string]ProblemPage
	Languages    []Lang
	StartTime    time.Time
	StopTime     time.Time
	Points       map[string]int
}

type ContestInfo struct {
	Name        string
	Description string
	Logo        string
	Website     string
}
