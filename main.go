package main

import (
	"final/config"
	"final/database"
	"final/routes"
)

func main() {
	config.Init()

	database.StartDatabase()
	s := routes.NewServer()

	s.Run()
}
