package main

import (
	"net/http"

	"github.com/muhammadadna/CRUD-GO/database"
	"github.com/muhammadadna/CRUD-GO/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8090", server)
}
