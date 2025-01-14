package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type TeamSchedule struct {
	ID           int
	TeamName     string `json:"team_name"`
	ClockInTime  string `json:"clock_in_time"`
	ClockOutTime string `json:"clock_out_time"`
}

func main() {
	data, err := os.ReadFile("team-01.json")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("the file has been opened")
	}

	var teamSchedule []*TeamSchedule

	err = json.Unmarshal(data, &teamSchedule)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("the data has been unmarshalled")
	}

	for _, schedule := range teamSchedule {
		log.Printf("ID: %d, Team Name: %s, Clock In Time: %s, Clock Out Time: %s\n",
			schedule.ID, schedule.TeamName, schedule.ClockInTime.Format(time.RFC3339), schedule.ClockOutTime.Format(time.RFC3339))
	}
}
