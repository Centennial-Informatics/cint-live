package scraper

import (
	"servermodule/app/models"
	"strings"

	"github.com/gocolly/colly"
)

func (client *Client) Info(contestURL string) (*models.ContestInfo, error) {
	info := models.ContestInfo{}

	client.C.OnHTML("#sidebar .sidebox:nth-of-type(1) a", func(e *colly.HTMLElement) {
		info.Name = e.Text
	})

	client.C.OnHTML("#sidebar .sidebox", func(e *colly.HTMLElement) {
		e.ForEach(".titled", func(i int, h *colly.HTMLElement) {
			if strings.TrimSpace(h.Text) == "→ About Contest" {
				info.Logo, _ = e.DOM.Find("div:nth-of-type(4) img").Attr("src")
				info.Logo = strings.TrimSpace(info.Logo)
				info.Description = strings.TrimSpace(e.DOM.Find("div:nth-of-type(5)").Text())
				info.Website, _ = e.DOM.Find("div:nth-of-type(6) a").Attr("href")
			}
		})
	})

	err := client.C.Visit(contestURL)
	client.C.OnHTMLDetach("#sidebar div:nth-of-type(1) a")

	return &info, err
}
