package scraper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"net/http"
	"regexp"
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

	client.C.OnHTML("body", func(e *colly.HTMLElement) {
		link := client.SetRCPC(e)
		client.C.OnHTMLDetach("body")
		client.C.OnHTML("input", func(e *colly.HTMLElement) {
			form[e.Attr("name")] = e.Attr("value")
			form["handleOrEmail"] = client.User
			form["password"] = client.Pass
		})
		client.C.Visit(link)
	})

	err := client.C.Visit(loginURL)
	if err != nil {
		return err
	}

	client.C.OnHTMLDetach("input")

	return client.C.Post(loginURL, form)
}

/*

thanks to
https://codeforces.com/blog/entry/80151
https://github.com/xalanq/cf-tool/pull/151/files

*/
func (client *Client) SetRCPC(e *colly.HTMLElement) string {
	link := "https://codeforces.com"
	reg := regexp.MustCompile(`Redirecting...`)
	is_redirected := (len(reg.FindStringSubmatch(e.Text)) > 0)
	if is_redirected {
		reg := regexp.MustCompile(`toNumbers\("(.+?)"\)`)
		res := reg.FindAllStringSubmatch(e.Text, -1)
		text, _ := hex.DecodeString(res[2][1])
		key, _ := hex.DecodeString(res[0][1])
		iv, _ := hex.DecodeString(res[1][1])

		block, _ := aes.NewCipher(key)
		mode := cipher.NewCBCDecrypter(block, iv)

		mode.CryptBlocks([]byte(text), []byte(text))

		var cookies []*http.Cookie
		cookie := &http.Cookie{
			Name:   "RCPC",
			Value:  hex.EncodeToString(text),
			Path:   "/",
			Domain: ".codeforces.com",
		}
		cookies = append(cookies, cookie)
		client.C.SetCookies(link, cookies)

		// get the redirect link
		reg = regexp.MustCompile(`href="(.+?)"`)
		link = reg.FindStringSubmatch(e.Text)[1]
	}

	return link
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
