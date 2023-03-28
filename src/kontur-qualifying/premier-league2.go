package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num int
	fmt.Fscan(reader, &num)

	teamMap := make(map[string]int)
	teamAteamBStr := fillTeamMap(teamMap, num)
	teamAteamBSls := strings.Split(teamAteamBStr, "-")
	teamA, teamB := teamAteamBSls[0], teamAteamBSls[1]

	position := positions(teamMap, teamA, teamB)
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

			return teamAteamB
		}
	}

}

func positions(teamMap map[string]int, teamA string, teamB string) []int {
	var victory int = 3 // Победа
	var defeat int = 0  // Поражение
	var draw int = 1    // Ничья

	position := make([]int, 3)

	// position[0]
}

func calculatePosition(teamMap map[string]int, teamA string, teamB string, result int) int {

}
