package handlers

import (
	"database/sql"
	"encoding/json"
	"main.go/repositories"
	"net/http"
)

type UserHandler struct {
	repo repositories.UserRepo
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		repo: *repositories.NewUserRepo(db),
	}
}
func (uh *UserHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {

}
func (uh *UserHandler) LogInHandler(w http.ResponseWriter, r *http.Request) {

}
func (uh *UserHandler) LogOutHandler(w http.ResponseWriter, r *http.Request) {

}

func (uh *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uh.repo.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
