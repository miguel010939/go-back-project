package main

import (
	"fmt"
	"main.go/routers"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	routers.Routes(r)

	fmt.Println("Server working on port 8090...")
	err := http.ListenAndServe(":8090", r)
	if err != nil {
		return
	}

}
