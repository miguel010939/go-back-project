package routers

import (
	"database/sql"
	"main.go/handlers"
	"net/http"
)

// Routes makes the connection between the different routes and their corresponding handler functions
func Routes(r *http.ServeMux, db *sql.DB) {
	uh := handlers.NewUserHandler(db)
	r.HandleFunc("/users", uh.GetAllUsersHandler)

	r.Handle("/", http.FileServer(http.Dir("./static/")))
}
