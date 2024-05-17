package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"main.go/config"
	db2 "main.go/db"
	"main.go/routers"
	"net/http"
)

func main() {
	db := db2.NewDBConnection(config.DefaultConnStr)
	if config.CreateTables {
		db2.CreateTables(db)
	}

	r := http.NewServeMux()
	routers.Routes(r, db)

	fmt.Println("Server working on port 8090...")
	err := http.ListenAndServe(":8090", r)
	if err != nil {
		return
	}

}
