package handlers

import (
	"database/sql"
	"main.go/repositories"
	"net/http"
)

type FavoriteHandler struct {
	repo repositories.FavoriteRepo
	auth repositories.AuthRepo
}

func NewFavoriteHandler(db *sql.DB) *FavoriteHandler {
	return &FavoriteHandler{
		repo: *repositories.NewFavoriteRepo(db),
		auth: *repositories.NewAuthRepo(db),
	}
}

func (fah *FavoriteHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & query params limit & offset, method Get
}
func (fah *FavoriteHandler) SaveFavorite(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (product), method Post

}
func (fah *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (product), method Delete

}
