package handlers

import (
	"database/sql"
	"encoding/json"
	"main.go/models"
	"main.go/repositories"
	"net/http"
)

type UserHandler struct {
	repo repositories.UserRepo
	auth repositories.AuthRepo
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		repo: *repositories.NewUserRepo(db),
		auth: *repositories.NewAuthRepo(db),
	}
}
func (uh *UserHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// takes json body usersingupform, method Post
	var user models.UserSignUpForm
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := uh.repo.UserSignUp(&user, &uh.auth)
	if err != nil {
		errorDispatch(w, r, err)
		return
	}
	w.Header().Set("sessionid", token)
	w.WriteHeader(http.StatusCreated)
}
func (uh *UserHandler) LogInHandler(w http.ResponseWriter, r *http.Request) {
	// takes json body userloginform, method Post
	var user models.UserLogInForm
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := uh.repo.UserLogIn(&user, &uh.auth)
	if err != nil {
		errorDispatch(w, r, err)
		return
	}
	w.Header().Set("sessionid", token)
	w.WriteHeader(http.StatusCreated)
}
func (uh *UserHandler) LogOutHandler(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" with token, method Delete
	token := r.Header.Get("sessionid")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	err := uh.repo.UserLogOut(token)
	if err != nil {
		errorDispatch(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
