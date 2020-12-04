package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Hello")
	}
}

func main() {

	/*** Data ***/
	allleague := make([]league, 1)
	teams := make([]team, 4)
	allleague[0] = league{id, "premier", "1", "30", "15-09-2020", "17-12-2020", [2]string{"edit", "delete"}, initScoreBoard(), nil}
	teams[0] = team{teamID, "Chelsea", "premier", time.Now().Format("01-02-2006"), [2]string{"edit", "delete"}}
	teams[1] = team{teamID, "Arsenal", "premier", time.Now().Format("01-02-2006"), [2]string{"edit", "delete"}}
	teams[2] = team{teamID, "Manchester City", "premier", time.Now().Format("01-02-2006"), [2]string{"edit", "delete"}}
	teams[3] = team{teamID, "Liverpool", "premier", time.Now().Format("01-02-2006"), [2]string{"edit", "delete"}}

	/*** Routings ***/
	r := mux.NewRouter()
	r.HandleFunc("/", index())
	r.HandleFunc("/leagues", getLeagueList(&allleague)).Methods("GET")
	r.HandleFunc("/leagues", addLeague(&allleague))
	r.HandleFunc("/leagues/{id}", getLeague(&allleague))
	r.HandleFunc("/leagues/{id}/teams", getTeamList(&teams)).Methods("GET")
	r.HandleFunc("/leagues/{id}/teams", addTeam(&teams))
	r.HandleFunc("/leagues/{id}/distribute-fixture", distributeFixture(&allleague))
	r.HandleFunc("/leagues/{id}/play-one-week", playOneWeek(&allleague))
	r.HandleFunc("/leagues/{id}/play-all", playAll(&allleague))
	r.HandleFunc("/leagues/{id}/predict-leaders", predict())

	http.ListenAndServe(":80", r)
}
