package handlers

import (
	"database/sql"
	"encoding/json"
	"main.go/models"
	"main.go/repositories"
	"net/http"
)

type FollowerHandler struct {
	repo repositories.FollowerRepo
	auth repositories.AuthRepo
}

func NewFollowerHandler(db *sql.DB) *FollowerHandler {
	return &FollowerHandler{
		repo: *repositories.NewFollowerRepo(db),
		auth: *repositories.NewAuthRepo(db),
	}
}
func (foh *FollowerHandler) GetUsersImFollowing(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token , method Get
	// user
	token := r.Header.Get("sessionid")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	userId, err2 := foh.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}
	followed, err3 := foh.repo.GetUsersWhomIFollow(userId)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	} // TODO maybe later simplify this so that the repo method itself returns an array of values, not pointers
	var followedArray []models.UserRepresentation
	for _, followedPtr := range followed {
		followedArray = append(followedArray, *followedPtr)
	}

	jsonData, e := json.Marshal(followedArray)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, e2 := w.Write(jsonData)
	if e2 != nil {
		http.Error(w, e2.Error(), http.StatusInternalServerError)
		return
	}
}
func (foh *FollowerHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (user), method Post
	followedId, err := ParseIntPathParam(r.URL.Path, "follow/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// user
	token := r.Header.Get("sessionid")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	userId, err2 := foh.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}

	err3 := foh.repo.FollowSomeone(userId, followedId)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (foh *FollowerHandler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (user), method Delete
	followedId, err := ParseIntPathParam(r.URL.Path, "follow/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// user
	token := r.Header.Get("sessionid")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	userId, err2 := foh.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}

	err3 := foh.repo.UnfollowSomeone(userId, followedId)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
