package activity

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dlaize/homedatakeeper/database"
	"github.com/dlaize/homedatakeeper/util"
	"github.com/gorilla/mux"
)

var router *mux.Router

func InitializeActivityRoutes(r *mux.Router) {
	router = r
	r.HandleFunc("/activities", getActivities).Methods("GET")
	r.HandleFunc("/activities", createActivity).Methods("POST")
	r.HandleFunc("/activities/{id:[0-9]+}", getActivity).Methods("GET")
	r.HandleFunc("/activities/{name}/{unit}/{value:[0-9]*\\.?[0-9]+}", handleActivity).Methods("GET")
}

func getActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid activity ID")
		return
	}

	act := activity{ID: id}
	if err := act.getActivity(database.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "Activity not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	util.RespondWithJSON(w, http.StatusOK, act)
}

// Zibase can only send GET requests so we generate the real POST
// from values of the GET request
// url : http://localhost:8000/activities/50/tv/min
func handleActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	value := vars["value"]
	name := vars["name"]
	unit := vars["unit"]

	payload := []byte(fmt.Sprintf(`{"name":"%s","value":%s,"unit":"%s"}`, name, value, unit))

	req, _ := http.NewRequest("POST", "/activities", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
}

func getActivities(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	activities, err := getListActivities(database.DB, start, count)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, activities)
}

func createActivity(w http.ResponseWriter, r *http.Request) {
	var act activity
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&act); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := act.createActivity(database.DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, act)
}
