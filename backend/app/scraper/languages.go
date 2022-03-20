package scraper

import "servermodule/app/models"

var langs = []models.Lang{
	{
		ID:   "54",
		Name: "C++17",
		Ext:  "cpp",
	},
	{
		ID:   "60",
		Name: "Java 11",
		Ext:  "java",
	},
	{
		ID:   "31",
		Name: "Python 3",
		Ext:  "py",
	},
}

var comments = map[string]string{
	"54": "//",
	"60": "//",
	"31": "#",
}

/* Languages - Get languages for problem/contest, static for now. */
func (client *Client) Languages() []models.Lang {
	return langs
}

func (client *Client) comment(language string) string {
	return comments[language]
}
