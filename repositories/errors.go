package repositories

import (
	"fmt"
	"net/http"
)

// Error values with associated htpp.Status that i return from the repository layer, so i can deal with them in the handler layer
type RelatedError int

const (
	Empty              RelatedError = http.StatusNotFound
	NotFound           RelatedError = http.StatusNotFound
	Conflict           RelatedError = http.StatusConflict
	NoPermission       RelatedError = http.StatusUnauthorized
	SomethingWentWrong RelatedError = http.StatusInternalServerError
	InvalidInput       RelatedError = http.StatusBadRequest
)

func (err RelatedError) Error() string {
	errMessage := fmt.Sprintf("Oh, no! Error code:%d", int(err))
	return errMessage
}
