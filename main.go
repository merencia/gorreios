package main

import (
	"github.com/merencia/gorreios/server"
)

func main() {
	server := server.New()

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
