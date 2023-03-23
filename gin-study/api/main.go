package main

import (
	"gin-study/api/database"
	"gin-study/api/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
