package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num int
	fmt.Fscan(reader, &num)

	teamMap := make(map[string]int)

	teamAteamBStr := fillTeamMap(teamMap, num) // Заполняет карту - ок

	teamAteamBSls := strings.Split(teamAteamBStr, "-")
	teamA, teamB := teamAteamBSls[0], teamAteamBSls[1]

	positions := positions(teamMap, teamA, teamB)
	fmt.Println(positions)
}

func fillTeamMap(teamMap map[string]int, num int) string {
	var teamAteamB string

	for i := 0; i <= num; i++ {
		if i < num {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			line := scanner.Text()
			spltSls := strings.Split(line, " ")
			team := spltSls[0]
			score, err := strconv.Atoi(spltSls[1])
			if err != nil {
				fmt.Println("converting string to int: ", err.Error())
			}
			teamMap[team] = score
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			teamAteamB = scanner.Text()

		}
	}
	return teamAteamB
}

func positions(teamMap map[string]int, teamA string, teamB string) []int {
	results := []string{"victory", "draw", "defeat"}
	resultsSls := make([]int, len(results))

	for i, value := range results {
		position := calculatePosition(teamMap, teamA, teamB, value)
		resultsSls[i] = position
	}

	return resultsSls
}

func calculatePosition(teamMap map[string]int, teamA string, teamB string, result string) int {
	victory := 3
	defeat := 0
	draw := 1

	var position int

	teamMapCalculate := make(map[string]int)
	for key, value := range teamMap {
		teamMapCalculate[key] = value
	}

	if result == "victory" {
		teamMapCalculate[teamA] += victory
		teamMapCalculate[teamB] += defeat
	} else if result == "draw" {
		teamMapCalculate[teamA] += draw
		teamMapCalculate[teamB] += draw
	} else if result == "defeat" {
		teamMapCalculate[teamA] += defeat
		teamMapCalculate[teamB] += victory
	}

	scores := make([]int, 0)
	for _, value := range teamMapCalculate {
		scores = append(scores, value)
	}

	sort.Ints(scores)

	for idx, value := range scores {
		if value == teamMap[teamA] {
			position = idx + 1
		}
	}

	fmt.Println("position: ", position)

	return len(teamMap) - position
}
