package repositories

import (
	"fmt"
	"net/http"
)

// Error values with associated htpp.Status that i return from the repository layer, so i can deal with them in the handler layer
type relatedError int

const (
	Empty              relatedError = http.StatusNotFound
	NotFound           relatedError = http.StatusNotFound
	Conflict           relatedError = http.StatusConflict
	NoPermission       relatedError = http.StatusUnauthorized
	SomethingWentWrong relatedError = http.StatusInternalServerError
	InvalidInput       relatedError = http.StatusBadRequest
)

func (err relatedError) Error() string {
	errMessage := fmt.Sprintf("Oh, no! Error code:%d", int(err))
	return errMessage
}
