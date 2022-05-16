package database

import (
	"encoding/csv"
	"os"
	"servermodule/app/models"
	"sort"
	"strconv"
)

func WriteStandingsToCSV(db *ContestDB, problemsCache *models.ContestData, division string, config *models.Configuration) error {
	var standings []StandingsEntry
	if db.SecretStandingsCache != nil {
		standings = db.SecretStandingsCache
	} else {
		standings = db.StandingsCache
	}

	f, err := os.Create("Standings" + division + ".csv")
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	cols := []string{"rank", "division", "team_name", "member_1", "member_2", "member_3", "member_4"}
	for _, problem := range problemsCache.Problems {
		cols = append(cols, "points_"+problem.ID)
	}

	cols = append(cols, "total_points")
	writer.Write(cols)

	rows := [][]string{}

	for _, standingsEntry := range standings {
		if standingsEntry.Division == division {
			row := []string{"1", standingsEntry.Division, standingsEntry.TeamName}
			row = append(row, standingsEntry.Members...)
			// add blanks when less than 4 members
			for i := len(standingsEntry.Members); i < 4; i++ {
				row = append(row, "")
			}
			totalPoints := 0
			for _, problem := range problemsCache.Problems {
				if verdict, good := standingsEntry.Verdicts[problem.ID]; good {
					totalPoints += verdict.Points
					row = append(row, strconv.Itoa(verdict.Points))
				} else {
					row = append(row, "0")
				}
			}
			row = append(row, strconv.Itoa(totalPoints))
			rows = append(rows, row)
		}
	}

	sort.Slice(rows, func(i, j int) bool {
		pts_i, _ := strconv.Atoi(rows[i][len(rows[i])-1])
		pts_j, _ := strconv.Atoi(rows[j][len(rows[j])-1])
		return pts_i > pts_j
	})

	for i, row := range rows {
		row[0] = strconv.Itoa(i + 1) // apply rank
		writer.Write(row)
	}

	return nil
}
