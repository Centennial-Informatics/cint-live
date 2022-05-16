package workers

import (
	"log"
	"servermodule/app/database"
	"servermodule/app/models"
	"servermodule/utils"
	"time"
)

/* StandingsWorker - Starts a goroutine that updates Standings every [interval] seconds. */
func StandingsWorker(db *database.ContestDB,
	interval time.Duration, config *models.Configuration) chan struct{} {

	ticker := time.NewTicker(interval)
	quit := make(chan struct{})

	go func() {
		updatedOnce := false // in case the app crashes after contest end
		startFreeze := utils.XBefore(config.ScoreboardFreeze, config.StopTime)
		endFreeze := utils.XAfter(config.ScoreboardUnfreeze, config.StopTime)
		for {
			select {
			case <-ticker.C:
				/* Stop updating standings when contest ends, but continue to allow submissions. */
				if !updatedOnce || utils.IsBefore(startFreeze) {
					db.UpdateStandings()
				} else if utils.IsAfter(startFreeze) && utils.IsBefore(endFreeze) {
					db.SecretStandingsCache = database.AggStandings(db) // only admin accounts can view secret standings
				} else {
					log.Println("Contest is over. Scoreboard no longer updating.")
					if db.SecretStandingsCache != nil {
						db.StandingsCache = db.SecretStandingsCache
					}
					ticker.Stop()
					return
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}
