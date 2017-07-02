package activity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dlaize/homedatakeeper/util"
	"github.com/gorilla/mux"
)

var router *mux.Router

func InitializeActivityRoutes(r *mux.Router) {
	router = r
	r.HandleFunc("/activities", createActivity).Methods("POST")
	r.HandleFunc("/activities/{name}/{unit}/{value:[0-9]*\\.?[0-9]+}", handleActivity).Methods("GET")
}

// Zibase can only send GET requests so we generate the real POST
// from values of the GET request
// url : http://localhost:8000/activities/tv/min/50
func handleActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	value := vars["value"]
	name := vars["name"]
	unit := vars["unit"]

	payload := []byte(fmt.Sprintf(`{"name":"%s","value":%s,"unit":"%s"}`, name, value, unit))

	req, _ := http.NewRequest("POST", "/activities", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
}

func createActivity(w http.ResponseWriter, r *http.Request) {
	var act Activity
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&act); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := act.createActivity(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, act)
}
