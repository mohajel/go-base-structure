package main

import (
	"base-1/internal/server"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// @title Good-Food
// @version 1.0
// @allout solutin for reservation managment
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://good-food/support
// @contact.email good-food-support@gmail.com

// @host localhost:8338
func main() {

	log.Infof("Starting App ...")

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
