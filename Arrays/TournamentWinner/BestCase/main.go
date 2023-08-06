package main

import "fmt"

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	pointMap := make(map[string]int)
	winner := ""
	bestTeam := ""
	for k, v := range competitions {
		if results[k] == HOME_TEAM_WON {
			winner = v[0]
		} else {
			winner = v[1]
		}
		if _, ok := pointMap[winner]; ok {
			pointMap[winner] = pointMap[winner] + 3
		} else {
			pointMap[winner] = 3
		}
		if pointMap[winner] > pointMap[bestTeam] {
			bestTeam = winner
		}
	}
	return bestTeam
}

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	fmt.Println(TournamentWinner(competitions, results))
}
