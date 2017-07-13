package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/merencia/gorreios/crawler"
)

func TrackPackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	events := crawler.GetPackage(vars["code"])
	json.NewEncoder(w).Encode(events)
}
