package main

import "fmt"

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	pointMap := make(map[string]int)
	winner := ""
	for k, v := range competitions {
		if results[k] == HOME_TEAM_WON {
			winner = v[0]
		} else {
			winner = v[1]
		}
		if _, ok := pointMap[winner]; ok {
			pointMap[winner] += 1
		} else {
			pointMap[winner] = 1
		}
	}
	maxSum := 0
	for key, val := range pointMap {
		if maxSum < val {
			maxSum = val
			winner = key
		}
	}
	return winner
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
