package routes

import (
	"database/sql"
	"net/http"

	"github.com/muhammadadna/CRUD-GO/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWordController())
	server.HandleFunc("/employee", controller.NewEmployeeController(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db))
}
