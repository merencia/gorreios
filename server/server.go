package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/merencia/gorreios/config"
	"github.com/merencia/gorreios/controllers"
)

func New() *http.Server {

	mux := mux.NewRouter()

	mux.HandleFunc("/track/{code}", controllers.TrackPackage)

	return &http.Server{
		Handler: mux,
		Addr:    ":" + config.Get().Port,
	}
}
