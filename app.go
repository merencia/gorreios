package main

import (
	"fmt"

	"github.com/merencia/gorreios/server"
)

func main() {
	server := server.New()

	fmt.Println("Starting Gorreios on port: " + server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
