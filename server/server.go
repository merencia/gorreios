package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() *http.Server {

	mux := mux.NewRouter()

	return &http.Server{
		Addr: ":5000",
	}
}
