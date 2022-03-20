package scraper

import (
	"servermodule/app/models"
	"strings"

	"github.com/gocolly/colly"
)

/* Problems - Scrape problems from contest page of codeforces gym [contestID]. */
func (client *Client) Problems(contestURL string) ([]models.Problem, error) {
	var rowLength = 4 // columns in problem row, necessary to ignore the "Add new problem" row

	problems := make([]models.Problem, 0)

	client.C.OnHTML("table[class=\"problems\"] tr", func(e *colly.HTMLElement) {
		if e.DOM.Find("td").Length() >= rowLength {
			problemID := e.DOM.Find("td:nth-of-type(1)")
			problemName := e.DOM.Find("td:nth-of-type(2) a:nth-of-type(1)")
			problems = append(problems, models.Problem{
				ID:   strings.TrimSpace(problemID.Text()),
				Name: strings.TrimSpace(problemName.Text()),
			})
		}
	})

	err := client.C.Visit(contestURL)
	client.C.OnHTMLDetach("table[class=\"problems\"] tr")

	return problems, err
}

func parseStatementDiv(rootSelector string, e *colly.HTMLElement) []models.ProblemStatementParagraph {
	ret := make([]models.ProblemStatementParagraph, 0)

	e.ForEach(rootSelector+" > p, center, pre", func(i int, h *colly.HTMLElement) {
		par := models.ProblemStatementParagraph{
			Text:  h.Text,
			Image: h.ChildAttr("img", "src"),
		}

		if h.Attr("class") == "verbatim" {
			code, err := h.DOM.Html()
			if err != nil {
				return
			}

			par.Code = code
		}

		ret = append(ret, par)
	})

	return ret
}

/* ProblemPage - Scrape problem statement DOM element from problem [problemID] of codeforces gym [contestID]. */
func (client *Client) ProblemPage(contestURL string, problemID string) (models.ProblemPage, error) {
	problemPage := models.ProblemPage{
		Header:              models.ProblemHeader{},
		Statement:           []models.ProblemStatementParagraph{},
		InputSpecification:  []models.ProblemStatementParagraph{},
		OutputSpecification: []models.ProblemStatementParagraph{},
		SampleTests:         []models.ProblemSampleTest{},
		Note:                []models.ProblemStatementParagraph{},
	}

	client.C.OnHTML(".problem-statement .header", func(e *colly.HTMLElement) {

		problemPage.Header.Title = e.ChildText(".title")
		problemPage.Header.TimeLimit = strings.ReplaceAll(e.ChildText(".time-limit"),
			e.ChildText(".time-limit .property-title"), "")
		problemPage.Header.MemLimit = strings.ReplaceAll(e.ChildText(".memory-limit"),
			e.ChildText(".memory-limit .property-title"), "")
		problemPage.Header.Input = strings.ReplaceAll(e.ChildText(".input-file"),
			e.ChildText(".input-file .property-title"), "")
		problemPage.Header.Output = strings.ReplaceAll(e.ChildText(".output-file"),
			e.ChildText(".output-file .property-title"), "")
	})

	client.C.OnHTML(".problem-statement > div:nth-child(2)", func(e *colly.HTMLElement) {
		problemPage.Statement = parseStatementDiv(".problem-statement > div:nth-child(2)", e)
	})

	client.C.OnHTML(".problem-statement .input-specification", func(e *colly.HTMLElement) {
		problemPage.InputSpecification = parseStatementDiv(".problem-statement .input-specification", e)
	})

	client.C.OnHTML(".problem-statement .output-specification", func(e *colly.HTMLElement) {
		problemPage.OutputSpecification = parseStatementDiv(".problem-statement .output-specification", e)
	})

	client.C.OnHTML(".problem-statement .sample-tests .sample-test", func(e *colly.HTMLElement) {
		problemPage.SampleTests = make([]models.ProblemSampleTest, 0)

		e.ForEach(".input", func(i int, h *colly.HTMLElement) {
			problemPage.SampleTests = append(problemPage.SampleTests, models.ProblemSampleTest{
				Input: h.ChildText("pre"),
			})
		})

		e.ForEach(".output", func(i int, h *colly.HTMLElement) {
			problemPage.SampleTests[i].Output = h.ChildText("pre")
		})
	})

	client.C.OnHTML(".problem-statement .note", func(e *colly.HTMLElement) {
		problemPage.Note = parseStatementDiv(".problem-statement .note", e)
	})

	err := client.C.Visit(contestURL + "/problem/" + problemID)
	client.C.OnHTMLDetach(".problem-statement .header")
	client.C.OnHTMLDetach(".problem-statement div :nth-of-type(1)")
	client.C.OnHTMLDetach(".problem-statement .input-specification")
	client.C.OnHTMLDetach(".problem-statement .output-specification")
	client.C.OnHTMLDetach(".problem-statement .sample-tests .sample-test")
	client.C.OnHTMLDetach(".problem-statement .note")

	return problemPage, err
}
