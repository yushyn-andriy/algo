package arrays

const (
	AWAY_TEAM_WON = 0
	HOME_TEAM_WON = 1
	AWAY_TEAM_IDX = 1
	HOME_TEAM_IDX = 0
)

func TournamentWinner(competitions [][]string, results []int) string {
	points := make(map[string]int)

	for idx := range competitions {
		comp := competitions[idx]
		winnerIdx := getWinnerIdx(results[idx])
		points[comp[winnerIdx]] += 3
	}

	currMax := 0
	currWinnerName := ""
	for name, score := range points {
		if score > currMax {
			currMax = score
			currWinnerName = name
		}
	}

	return currWinnerName
}

func getWinnerIdx(idx int) int {
	if idx == AWAY_TEAM_WON {
		return AWAY_TEAM_IDX
	}

	return HOME_TEAM_IDX
}
