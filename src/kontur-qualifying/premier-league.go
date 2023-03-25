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
	teamSlice := make([]string, num)
	fillSlice(num, teamSlice)

	mapTeam := make(map[string]int)
	NewMapTeam(teamSlice, mapTeam)

	for i, val := range mapTeam {
		fmt.Println(i, " : ", val)
	}
}

func fillSlice(num int, teamSlice []string) {
	reader := bufio.NewReader(os.Stdin)
	var elem string
	for i := 0; i < num; i++ {
		fmt.Fscan(reader, &elem)
		teamSlice[i] = elem
	}

	fmt.Println("teamSlice:")
	for i, val := range teamSlice {
		fmt.Println(i, ": ", val)
	}
}

func NewMapTeam(teamSlice []string, mapTeam map[string]int) {
	for _, value := range teamSlice {
		splt := strings.Split(value, " ")
		team := splt[0]
		scoreStr := splt[1]
		score, err := strconv.Atoi(scoreStr)
		if err != nil {
			fmt.Println("spql[1]: ", splt[1])
			fmt.Println("converting string to int: ", err.Error())
		}
		mapTeam[team] = score
	}
}
