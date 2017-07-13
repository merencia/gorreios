package server

import (
	"net/http"

	"github.com/apex/log"
	"github.com/gorilla/mux"
	"github.com/merencia/gorreios/config"
	"github.com/merencia/gorreios/controllers"
)

func New() *http.Server {

	mux := mux.NewRouter()

	mux.HandleFunc("/track/{code}", controllers.TrackPackage)
	conf := config.Get()

	log.Info("starting on port " + conf.Port)

	return &http.Server{
		Handler: mux,
		Addr:    ":" + conf.Port,
	}
}
