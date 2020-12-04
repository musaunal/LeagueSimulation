package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type scores struct {
	Name     string `json:"team.name"`
	Won      int    `json:"won"`
	Drawn    int    `json:"drawn"`
	Lost     int    `json:"lost"`
	For      int    `json:"for"`
	Against  int    `json:"against"`
	GoalDiff int    `json:"goal_diff"`
	Points   int    `json:"points"`
}

type fixture struct {
	ID        int    `json:"id"`
	Week      int    `json:"week"`
	Home      string `json:"home_team_name"`
	Away      string `json:"away_team_name"`
	Played    bool   `json:"is_played"`
	HomeScore int    `json:"home_team_score"`
	AwayScore int    `json:"away_team_score"`
}

type prediction struct {
	ID   int    `json:"team_id"`
	Team string `json:"team_name"`
	Rate int    `json:"rate"`
}

func sampleFixture() []fixture {
	pool := make([]fixture, 8)
	pool[0] = fixture{0, 1, "Chelsea", "Arsenal", false, 0, 0}
	pool[1] = fixture{1, 1, "Manchester City", "Liverpool", false, 0, 0}
	pool[2] = fixture{2, 2, "Chelsea", "Liverpool", false, 0, 0}
	pool[3] = fixture{3, 2, "Manchester City", "Arsenal", false, 0, 0}
	pool[4] = fixture{4, 3, "Arsenal", "Chelsea", false, 0, 0}
	pool[5] = fixture{5, 3, "Liverpool", "Manchester City", false, 0, 0}
	pool[6] = fixture{6, 4, "Liverpool", "Chelsea", false, 0, 0}
	pool[7] = fixture{7, 4, "Arsenal", "Manchester City", false, 0, 0}
	return pool
}

func initScoreBoard() []scores {
	s := make([]scores, 4)
	s[0] = scores{"Chelsea", 0, 0, 0, 0, 0, 0, 0}
	s[1] = scores{"Arsenal", 0, 0, 0, 0, 0, 0, 0}
	s[2] = scores{"Manchester City", 0, 0, 0, 0, 0, 0, 0}
	s[3] = scores{"Liverpool", 0, 0, 0, 0, 0, 0, 0}
	return s
}

func getRandomTeam() string {
	// TODO: implement random fixture
	return ""
}

func teamNameToID(s string) int {
	switch s {
	case "Chelsea":
		return 0
	case "Arsenal":
		return 1
	case "Manchester City":
		return 2
	case "Liverpool":
		return 3
	}
	return -1
}

func updateScoreBoard(s1 int, s2 int, board *[]scores, team1 string, team2 string) {
	if s1 == s2 {
		(*board)[teamNameToID(team1)].Drawn++
		(*board)[teamNameToID(team2)].Drawn++
	} else if s1 > s2 {
		(*board)[teamNameToID(team1)].Won++
		(*board)[teamNameToID(team1)].Points += 3
		(*board)[teamNameToID(team2)].Lost++
	} else {
		(*board)[teamNameToID(team2)].Won++
		(*board)[teamNameToID(team2)].Points += 3
		(*board)[teamNameToID(team1)].Lost++
	}

	(*board)[teamNameToID(team1)].GoalDiff += s1
	(*board)[teamNameToID(team2)].GoalDiff += s2

}

func initPredictions() []prediction {
	var pool = make([]prediction, 4)
	pool[0] = prediction{0, "Chelsea", 25}
	pool[1] = prediction{1, "Arsenal", 25}
	pool[2] = prediction{2, "Manchester City", 25}
	pool[3] = prediction{3, "Liverpool", 25}
	return pool
}

func distributeFixture(list *[]league) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		i, _ := strconv.Atoi(vars["id"])
		(*list)[i].Fixture = sampleFixture()
		isLeagueStarted = true

		fmt.Fprintln(w, "Success")
	}
}

func playOneWeek(list *[]league) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		j, _ := strconv.Atoi(vars["id"])

		if currentWeek < 4 && isLeagueStarted {
			r1 := rand.NewSource(time.Now().UnixNano())
			r := rand.New(r1)

			for i := 0; i < 2; i++ {
				score1 := r.Intn(5)
				score2 := r.Intn(5)
				match := (*list)[j].Fixture[2*currentWeek+i]

				(*list)[j].Fixture[2*currentWeek+i].HomeScore = score1
				(*list)[j].Fixture[2*currentWeek+i].AwayScore = score2
				(*list)[j].Fixture[2*currentWeek+i].Played = true

				updateScoreBoard(score1, score2, &(*list)[j].ScoreBoard, match.Home, match.Away)
			}

			currentWeek++
			fmt.Fprintln(w, "Success")
		}
		fmt.Fprintln(w, "Failure")
	}
}

func playAll(list *[]league) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		j, _ := strconv.Atoi(vars["id"])

		if currentWeek < 4 && isLeagueStarted {
			r1 := rand.NewSource(time.Now().UnixNano())
			r := rand.New(r1)

			for k := currentWeek; k < 4; k++ {
				for i := 0; i < 2; i++ {
					score1 := r.Intn(5)
					score2 := r.Intn(5)
					match := (*list)[j].Fixture[2*currentWeek+i]

					(*list)[j].Fixture[2*currentWeek+i].HomeScore = score1
					(*list)[j].Fixture[2*currentWeek+i].AwayScore = score2
					(*list)[j].Fixture[2*currentWeek+i].Played = true

					updateScoreBoard(score1, score2, &(*list)[j].ScoreBoard, match.Home, match.Away)
				}
				currentWeek++
			}
			fmt.Fprintln(w, "Success")
		}
		fmt.Fprintln(w, "Failure")
	}
}

func predict() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.WriteHeader(http.StatusCreated)

		enc := json.NewEncoder(w)
		enc.Encode(initPredictions())
	}
}
