package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dlaize/homedatakeeper/activity"
	"github.com/dlaize/homedatakeeper/probe"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeControllers(a.Router)
}

func (a *App) initializeControllers(r *mux.Router) {
	activity.InitializeActivityRoutes(r)
	probe.InitializeTempProbeRoutes(r)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
