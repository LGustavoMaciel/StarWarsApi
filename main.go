package main

import (
	"StarWarsApi/database"
	"StarWarsApi/routes"
)

func main() {

	database.Init()
	routes.Routes()

	

}