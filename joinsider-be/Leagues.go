package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var id int = 0
var currentWeek = 0
var isLeagueStarted = false

type league struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CurrentWeek string    `json:"current_week"`
	TotalWeek   string    `json:"total_week"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	Actions     [2]string `json:"actions"`
	ScoreBoard  []scores  `json:"scoreboard"`
	Fixture     []fixture `json:"fixture"`
}

func getLeagueList(list *[]league) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.WriteHeader(http.StatusCreated)

		enc := json.NewEncoder(w)
		enc.Encode(*list)
	}
}

func getLeague(list *[]league) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.WriteHeader(http.StatusCreated)

		vars := mux.Vars(r)
		i, _ := strconv.Atoi(vars["id"])
		s := *list

		enc := json.NewEncoder(w)
		enc.Encode(s[i])
	}
}

/*** post ile gelen data okunmuyor body boş, form valuelerde boş ***/
func addLeague(list *[]league) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)

		// r.ParseForm()
		// fmt.Println(r.Form)
		// fmt.Println(r.PostForm)

		// for key, value := range r.Form {
		// 	fmt.Printf("%s = %s\n", key, value)
		// }
		// fmt.Println("request  " + r.RequestURI)
		// fmt.Println("url " + r.URL.Path)

		// body, _ := ioutil.ReadAll(r.Body)
		// if r.Body != http.NoBody {
		// 	fmt.Println("body is : " + string(body))
		// }
		// fmt.Println("mux : " + vars["name"])
		id++
		*list = append(*list, league{id, "tr", "2", "32", "1", "1", [2]string{"edit", "delete"}, nil, nil}) // temp data ekliyorum
		fmt.Fprintln(w, "Success")
	}
}
