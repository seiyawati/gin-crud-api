package main

import (
	"app/database"
	"app/server"
)

func main() {
	database.InitDB()

	defer database.CloseDB()

	server.InitServer()
}
