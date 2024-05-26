package handlers

import (
	"errors"
	"main.go/repositories"
	"net/http"
)

func errorDispatch(w http.ResponseWriter, r *http.Request, e error) {
	switch {
	case errors.Is(e, repositories.SomethingWentWrong):
		http.Error(w, e.Error(), http.StatusInternalServerError)
	case errors.Is(e, repositories.NotFound):
		http.Error(w, e.Error(), http.StatusNotFound)
	case errors.Is(e, repositories.InvalidInput):
		http.Error(w, e.Error(), http.StatusBadRequest)
	case errors.Is(e, repositories.Conflict):
		http.Error(w, e.Error(), http.StatusConflict)
	case errors.Is(e, repositories.NoPermission):
		http.Error(w, e.Error(), http.StatusUnauthorized)
	default:
		http.Error(w, e.Error(), http.StatusTeapot)
	}
}
