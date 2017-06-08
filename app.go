package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"github.com/dlaize/homedatakeeper/activity"
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
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
