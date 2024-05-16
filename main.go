package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	x := []byte("Hello World")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(x)
		if err != nil {
			return
		}

	})
	err := http.ListenAndServe(":8090", r)
	if err != nil {
		return
	}
	fmt.Println("Server working on port 8090...")

}
