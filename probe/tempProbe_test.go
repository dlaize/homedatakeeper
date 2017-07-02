package probe_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/dlaize/homedatakeeper/database"
	"github.com/dlaize/homedatakeeper/probe"
	"github.com/dlaize/homedatakeeper/util"
	"github.com/gorilla/mux"
)

var r *mux.Router

func TestMain(m *testing.M) {
	r = mux.NewRouter()
	database.Initialize()
	probe.InitializeTempProbeRoutes(r)
	code := m.Run()
	database.Close()
	os.Exit(code)
}

func TestCreateTempProbe(t *testing.T) {

	payload := []byte(`{"name":"test_tempProbe", "etage":1, "temp":21.2, "hygro":55.3}`)

	req, _ := http.NewRequest("POST", "/tempprobes", bytes.NewBuffer(payload))
	response := util.ExecuteRequest(req, r)

	util.CheckResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test_tempProbe" {
		t.Errorf("Expected name to be 'test_tempProbe'. Got '%v'", m["name"])
	}

	if m["etage"] != 1.0 {
		t.Errorf("Expected etage to be '1'. Got '%v'", m["etage"])
	}

	if m["temp"] != 21.2 {
		t.Errorf("Expected temp to be '21.2'. Got '%v'", m["value"])
	}

	if m["hygro"] != 55.3 {
		t.Errorf("Expected hygro to be '55.3'. Got '%v'", m["value"])
	}
}
