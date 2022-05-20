package scraper

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"servermodule/app/models"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var submissionVerdicts = []string{
	"Accepted",
	"Ok",
	"Partial",
	"Wrong answer",
	"Failed",
	"Compilation error",
	"Runtime error",
	"Presentation error",
	"Time limit exceeded",
	"Memory limit exceeded",
	"Idleness limit exceeded",
	"Security violated",
	"Crashed",
	"Input preparation crashed",
	"Challenged",
	"Skipped",
	"Rejected",
}

func getFileBytes(file *multipart.FileHeader) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	b, err := io.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return b, nil
}

/* Submit - Submit code [submission] in [language] to problem [problemID] of codeforces gym [contestID]
and return the submission ID. */
func (client *Client) Submit(contestURL string, problemID string, submission string,
	file *multipart.FileHeader, language string) (string, error) {
	var submissionID string

	submitURL := contestURL + "/submit"
	form := map[string]string{}
	submitFile := false

	var b []byte

	var err error

	if file != nil {
		b, err = getFileBytes(file)
		if err == nil {
			submitFile = true
		}
	}

	client.C.OnHTML(".submit-form input", func(e *colly.HTMLElement) {
		form[e.Attr("name")] = e.Attr("value")
		form["submittedProblemIndex"] = problemID
		form["programTypeId"] = language
		if submitFile {
			form["source"] = string(b) + client.comment(language) +
				fmt.Sprintf("%d", time.Now().UnixNano())
		} else {
			form["source"] = submission + client.comment(language) +
				fmt.Sprintf("%d", time.Now().UnixNano())
		}
	})

	err = client.C.Visit(submitURL)
	if err != nil {
		client.Available = true
		return "Error", err
	}

	client.C.OnHTMLDetach((".submit-form input"))
	client.C.OnHTML(".status-frame-datatable tr:nth-of-type(2) td:nth-of-type(1)", func(e *colly.HTMLElement) {
		client.Available = true // client must wait until submission ID is scraped before submititng again
		submissionID = strings.TrimSpace(e.Text)
		log.Println("Scraper", client.Name, "submitted", submissionID, "to", problemID)
	})

	client.Available = false
	err = client.C.Post(submitURL, form)
	client.C.OnHTMLDetach(".status-frame-datatable tr:nth-of-type(2) td:nth-of-type(1)")

	if err != nil {
		log.Println(err)
		log.Println("Scraper", client.Name, "disconnected. Logging back in.")
		client.Login()
		client.Available = true
		return submissionID, err
	}

	if len(submissionID) > 0 {
		client.Verdict[submissionID] = &models.Verdict{ // remember to set UserID outside
			Status:    "Pending",
			Verdict:   "Pending",
			ProblemID: &problemID,
		}
	} else {
		client.Available = true
		return submissionID, errors.New("error during submission")
	}

	return submissionID, err
}

/* Status - Get verdict status for submission [submissionID] in codeforces gym [contestID]. */
func (client *Client) Status(contestURL string, submissionID string) *models.Verdict {
	submissionURL := contestURL + "/submission/" + submissionID

	verdict := client.Verdict[submissionID]

	client.C.OnHTML("table:nth-of-type(1) tr:nth-of-type(2) td:nth-of-type(5)", func(e *colly.HTMLElement) {
		verdict.Verdict = strings.TrimSpace(e.Text)
	})

	err := client.C.Visit(submissionURL)
	if err != nil {
		log.Println(err)
	}

	client.C.OnHTMLDetach("table:nth-of-type(1) tr:nth-of-type(2) td:nth-of-type(5)")

	for _, v := range submissionVerdicts {
		if len(verdict.Verdict) >= len(v) && verdict.Verdict[:len(v)] == v {
			verdict.Status = "Final"
		}
	}

	return verdict
}
