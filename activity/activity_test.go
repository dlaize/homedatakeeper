package activity_test

import (
	"testing"
	"net/http"
	"encoding/json"
	"bytes"
	"strconv"
	"log"
	"github.com/dlaize/homedatakeeper/util"
	"github.com/dlaize/homedatakeeper/database"
	"os"
	"github.com/gorilla/mux"
	"github.com/dlaize/homedatakeeper/activity"
)

var r *mux.Router

func TestMain(m *testing.M) {
	r = mux.NewRouter()
	database.Connect("5433")
	ensureTableExists()
	activity.InitializeActivityRoutes(r)
	code := m.Run()
	database.Close()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := database.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	database.DB.Exec("DELETE FROM activities")
	database.DB.Exec("ALTER SEQUENCE activities_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS activities
(
	id SERIAL,
    	name TEXT NOT NULL,
    	unit TEXT NOT NULL,
    	value NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    	CONSTRAINT activities_pkey PRIMARY KEY (id)
)`

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/activities", nil)
	response := util.ExecuteRequest(req, r)

	util.CheckResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentActivity(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/activities/11", nil)
	response := util.ExecuteRequest(req, r)

	util.CheckResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Activity not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Activity not found'. Got '%s'", m["error"])
	}
}

func TestCreateActivity(t *testing.T) {

	payload := []byte(`{"name":"test activity","value":11.22, "unit":"cm"}`)

	req, _ := http.NewRequest("POST", "/activities", bytes.NewBuffer(payload))
	response := util.ExecuteRequest(req, r)

	util.CheckResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test activity" {
		t.Errorf("Expected product name to be 'test activity'. Got '%v'", m["name"])
	}

	if m["value"] != 11.22 {
		t.Errorf("Expected activity value to be '11.22'. Got '%v'", m["value"])
	}

	if m["unit"] != "cm" {
		t.Errorf("Expected activity unit to be 'cm'. Got '%v'", m["value"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetActivity(t *testing.T) {
	clearTable()
	addActivities(1)

	req, _ := http.NewRequest("GET", "/activities/1", nil)
	response := util.ExecuteRequest(req, r)

	util.CheckResponseCode(t, http.StatusOK, response.Code)
}

func addActivities(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		database.DB.Exec("INSERT INTO activities(name, value, unit) VALUES($1, $2, $3)", "Product "+strconv.Itoa(i), (i+1.0)*10, "minutes")
	}
}