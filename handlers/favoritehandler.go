package handlers

import (
	"database/sql"
	"encoding/json"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
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
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	// if the query param is not found, the var is assigned a nonsensical value, so it is ignored by the repo method
	if limitStr == "" {
		limitStr = "-1"
	}
	if offsetStr == "" {
		offsetStr = "-1"
	}
	limit, err2 := strconv.Atoi(limitStr)
	offset, err3 := strconv.Atoi(offsetStr)
	if err2 != nil || err3 != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		return
	}

	// user
	token := r.Header.Get("sessionid")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	userId, err2 := fah.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}

	favorites, err3 := fah.repo.GetFavorites(userId, limit, offset)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	} // TODO maybe later simplify this so that the repo method itself returns an array of values, not pointers
	var favoritesArray []models.ProductRepresentation
	for _, favoritePtr := range favorites {
		favoritesArray = append(favoritesArray, *favoritePtr)
	}

	jsonData, e := json.Marshal(favoritesArray)
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
	w.WriteHeader(http.StatusOK)
}
func (fah *FavoriteHandler) SaveFavorite(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (product), method Post
	id, err := ParseIntPathParam(r.URL.Path, "favorites/")
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
	userId, err2 := fah.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}

	err3 := fah.repo.SaveFavorite(userId, id)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (fah *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (product), method Delete
	id, err := ParseIntPathParam(r.URL.Path, "favorites/")
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
	userId, err2 := fah.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}

	err3 := fah.repo.DeleteFavorite(userId, id)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
