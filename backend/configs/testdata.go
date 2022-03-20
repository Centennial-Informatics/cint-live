package configs

import "servermodule/app/models"

var testData = models.TestDataConfiguration{
	ContestID:  "328477",
	ContestURL: "https://codeforces.com/gym/328477",
	Problems: []models.Problem{
		{
			ID:   "A",
			Name: "Breakout Rooms",
		},
		{
			ID:   "B",
			Name: "Lockpicker",
		},
		{
			ID:   "C",
			Name: "Ramen",
		},
		{
			ID:   "D",
			Name: "Zombies",
		},
		{
			ID:   "E",
			Name: "Festival",
		},
		{
			ID:   "F",
			Name: "Heist",
		},
	},
	ProblemID:   "A",
	Submission:  "./testdata/test.cpp", // relative to tested package
	Language:    "54",
	FirebaseURL: "https://rosy-zoo-315017-default-rtdb.firebaseio.com/",
	Info: models.ContestInfo{
		Name: "CCC Demo",
		Description: "Programming competition open to all Howard County students." +
			" Make sure to fill out the sign up form for the competition on our website: " +
			"https://chscp.weebly.com/events.html to get tournament updates and receive prizes.",
		Website: "https://chscp.weebly.com",
	},
}

func TestData() models.TestDataConfiguration {
	return testData
}
