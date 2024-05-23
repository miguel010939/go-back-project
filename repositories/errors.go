package repositories

// TODO si tengo tiempo, explicar mejor porque creo este tipo de error
import "net/http"

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
	return "KABOOOM!!"
}
