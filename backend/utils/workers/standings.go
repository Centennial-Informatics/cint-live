package workers

import (
	"servermodule/app/database"
	"time"
)

/* StandingsWorker - Starts a goroutine that updates Standings every [interval] seconds. */
func StandingsWorker(db *database.ContestDB,
	interval time.Duration) chan struct{} {

	ticker := time.NewTicker(interval)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				db.UpdateStandings()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}
