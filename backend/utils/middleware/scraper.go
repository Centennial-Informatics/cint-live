package middleware

import (
	"servermodule/app/scraper"
	"time"
)

/* ScrapeWorker - Starts a goroutine that creates a cache for codeforces gym [contestID]
and updates it every [interval] seconds. */
func ScrapeWorker(client *scraper.Client, contestURL string, contestID string,
	intervalSlow time.Duration) chan struct{} {
	client.InitCache(contestURL, contestID)

	ticker := time.NewTicker(intervalSlow)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				client.UpdateCache()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}
