package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var teamID int = 3

type team struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	LeagueName string    `json:"league.name"`
	UpdatedAt  string    `json:"updated_at"`
	Actions    [2]string `json:"actions"`
}

func getTeamList(list *[]team) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.WriteHeader(http.StatusCreated)

		enc := json.NewEncoder(w)
		enc.Encode(*list)
	}
}

// temp data ekliyorum
func addTeam(list *[]team) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if teamID < 3 {
			teamID++
			*list = append(*list, team{teamID, "takım1", "premier", time.Now().Format("01-02-2006"), [2]string{"edit", "delete"}})
			fmt.Fprintln(w, "Success")
		} else {
			fmt.Fprintln(w, "Failure")
		}
	}
}
