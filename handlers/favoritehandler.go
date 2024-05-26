package handlers

import (
	"database/sql"
	"main.go/repositories"
	"net/http"
)

type FavoriteHandler struct {
	repo repositories.FavoriteRepo
}

func NewFavoriteHandler(db *sql.DB) *FavoriteHandler {
	return &FavoriteHandler{
		repo: *repositories.NewFavoriteRepo(db),
	}
}

func (fah *FavoriteHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {

}
func (fah *FavoriteHandler) SaveFavorite(w http.ResponseWriter, r *http.Request) {

}
func (fah *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {

}
