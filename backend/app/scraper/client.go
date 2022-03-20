package scraper

import (
	"log"
	"servermodule/app/models"

	"github.com/gocolly/colly"
)

type Client struct {
	C         *colly.Collector           // web scraper
	Cached    *models.ContestData        // cf contest problem statements/info
	Verdict   map[string]*models.Verdict // (submissionID, verdict)
	Available bool                       // is this in use
	Name      string                     // identifier for logging
	User      string
	Pass      string
}

/* Set default callbacks and configs on collector [c]. */
func prepareCollector(c *colly.Collector, name string) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Scraper", name, "visiting", r.URL)
	})

	c.AllowURLRevisit = true
}

/* NewClient - Create new colly client and authenticate with codeforces handle or email [user] and password [pass]. */
func NewClient(name string, user string, pass string) (*Client, error) {
	client := Client{
		C:         colly.NewCollector(),
		Available: true,
		Name:      name,
		User:      user,
		Pass:      pass,
	}

	prepareCollector(client.C, client.Name)

	err := client.Login()

	return &client, err
}

func (client *Client) Login() error {
	form := map[string]string{}
	loginURL := "https://codeforces.com/enter"

	client.C.OnHTML("input", func(e *colly.HTMLElement) {
		form[e.Attr("name")] = e.Attr("value")
		form["handleOrEmail"] = client.User
		form["password"] = client.Pass
	})

	err := client.C.Visit(loginURL)
	if err != nil {
		return err
	}

	client.C.OnHTMLDetach(("input"))

	return client.C.Post(loginURL, form)
}

func (client *Client) CopyCache(cache *models.ContestData) {
	client.Cached = cache
	client.Verdict = map[string]*models.Verdict{}
}

/* InitCache - Initialize new data cache for codeforces gym [contestID]. */
func (client *Client) InitCache(contestURL string, contestID string) {
	client.Cached = &models.ContestData{}
	client.Cached.ContestID = contestID
	client.Cached.ContestURL = contestURL
	client.Cached.Problems = []models.Problem{}
	client.Cached.ProblemPages = map[string]models.ProblemPage{}
	client.Cached.Info = models.ContestInfo{}
	client.Verdict = map[string]*models.Verdict{}
	client.UpdateCache()
}

/* UpdateCache - Update cached resources for contest. */
func (client *Client) UpdateCache() {
	log.Println("Updating cache")

	problems, err := client.Problems(client.Cached.ContestURL)
	if err != nil {
		log.Println(err)
	} else { // safeguard existing data if scrape fails
		if len(problems) == 0 {
			log.Println("Scraper", client.Name, "disconnected. Logging back in.")
			client.Login()
		} else {
			client.Cached.Problems = problems
		}

		for _, prob := range client.Cached.Problems {
			page, err := client.ProblemPage(client.Cached.ContestURL, prob.ID)
			if err != nil {
				log.Println(err)
			} else {
				client.Cached.ProblemPages[prob.ID] = page
			}
		}
	}

	info, err := client.Info(client.Cached.ContestURL)
	if err != nil {
		log.Println(err)
	} else {
		client.Cached.Info = *info
	}
}
