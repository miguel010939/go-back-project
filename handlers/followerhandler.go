package handlers

import (
	"database/sql"
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
}
func (foh *FollowerHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (user), method Post
}
func (foh *FollowerHandler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (user), method Delete
}
