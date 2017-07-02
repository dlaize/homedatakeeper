package activity_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/dlaize/homedatakeeper/activity"
	"github.com/dlaize/homedatakeeper/database"
	"github.com/dlaize/homedatakeeper/util"
	"github.com/gorilla/mux"
)

var r *mux.Router

func TestMain(m *testing.M) {
	r = mux.NewRouter()
	database.Initialize()
	activity.InitializeActivityRoutes(r)
	code := m.Run()
	database.Close()
	os.Exit(code)
}

func TestCreateActivity(t *testing.T) {

	payload := []byte(`{"name":"test_activity","value":11.22, "unit":"min"}`)

	req, _ := http.NewRequest("POST", "/activities", bytes.NewBuffer(payload))
	response := util.ExecuteRequest(req, r)

	util.CheckResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test_activity" {
		t.Errorf("Expected name to be 'test_activity'. Got '%v'", m["name"])
	}

	if m["value"] != 11.22 {
		t.Errorf("Expected value to be '11.22'. Got '%v'", m["value"])
	}

	if m["unit"] != "min" {
		t.Errorf("Expected unit to be 'min'. Got '%v'", m["value"])
	}
}
