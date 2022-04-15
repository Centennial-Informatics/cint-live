package workers

import (
	"log"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	const bInMb = 1024 * 1024

	return b / bInMb
}

/* MemLogger - spits out memory stats every [interval] seconds. */
func MemLogger(interval time.Duration) {
	go func() {
		t := time.NewTicker(interval)

		for ; true; <-t.C {
			var m runtime.MemStats

			runtime.ReadMemStats(&m)
			log.Printf("Alloc = %v MiB", bToMb(m.Alloc))
			log.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
			log.Printf("\tSys = %v MiB", bToMb(m.Sys))
			log.Printf("\tNumGC = %v\n", m.NumGC)
		}
	}()
}
