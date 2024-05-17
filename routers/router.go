package routers

import "net/http"

func Routes(r *http.ServeMux) {

	r.Handle("/", http.FileServer(http.Dir("./static/")))
}
