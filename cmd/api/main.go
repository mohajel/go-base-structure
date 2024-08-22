package main

import (
	"base-1/internal/server"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Infof("Starting App ...")

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
