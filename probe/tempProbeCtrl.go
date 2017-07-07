package probe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dlaize/homedatakeeper/util"
	"github.com/gorilla/mux"
)

var router *mux.Router

// InitializeTempProbeRoutes : set routes
func InitializeTempProbeRoutes(r *mux.Router) {
	router = r
	r.HandleFunc("/tempprobes", createTempProbe).Methods("POST")
	r.HandleFunc("/tempprobes/{name}/{etage:[0-9]{1}}/{temp:[0-9]*\\.?[0-9]+}/{hygro:[0-9]*\\.?[0-9]+}", handleTempProbe).Methods("GET")
}

// Zibase can only send GET requests so we generate the real POST
// from values of the GET request
// url : http://localhost:8000/tempprobes/ch_florent/1/20.5/55.2
func handleTempProbe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	etage := vars["etage"]
	name := vars["name"]
	temp := vars["temp"]
	hygro := vars["hygro"]

	//zibase send 231 for 23.1°
	i, _ := strconv.ParseFloat(temp, 64)
	ct := i / 10

	payload := []byte(fmt.Sprintf(`{"name":"%s","etage":%s,"temp":%f, "hygro":%s}`, name, etage, ct, hygro))

	req, _ := http.NewRequest("POST", "/tempprobes", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)
}

func createTempProbe(w http.ResponseWriter, r *http.Request) {
	var tp TempProbe
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tp); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := tp.createTempProbe(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, tp)
}
