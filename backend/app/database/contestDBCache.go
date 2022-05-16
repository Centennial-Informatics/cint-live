package database

import "servermodule/app/models"

func NewContestDB(path string, contestCache ...*models.ContestData) (*ContestDB, error) {
	db, err := NewEmptyDB(path)
	if err != nil {
		return db, err
	}

	if len(contestCache) > 0 {
		db.StandardCache = contestCache[0]
	}

	return db, err
}
