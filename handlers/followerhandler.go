package handlers

import (
	"database/sql"
	"main.go/repositories"
	"net/http"
)

type FollowerHandler struct {
	repo repositories.FollowerRepo
}

func NewFollowerHandler(db *sql.DB) *FollowerHandler {
	return &FollowerHandler{
		repo: *repositories.NewFollowerRepo(db),
	}
}
func (foh *FollowerHandler) GetUsersImFollowing(w http.ResponseWriter, r *http.Request) {

}
func (foh *FollowerHandler) FollowUser(w http.ResponseWriter, r *http.Request) {

}
func (foh *FollowerHandler) UnfollowUser(w http.ResponseWriter, r *http.Request) {

}
