package main

import (
	"app/server"
	"app/database"
)

func main() {
	database.InitDB()

	defer database.CloseDB()

	server.InitServer()
}
