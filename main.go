package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"main.go/config"
	db2 "main.go/db"
	"main.go/factories"
	"main.go/handlers"
	"main.go/routers"
	"net/http"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err1 := db2.NewDBConnection(connStr)
	if err1 != nil {
		if config.EnforceSuccessfulDBConnection {
			panic(err1)
		}
		fmt.Println(err1)
		fmt.Println("Warning: Failed to connect to database")
	}
	db2.DropTables(db)
	db2.CreateTables(db)
	factories.UralFactories(db, 100, 200, 50, 30)

	r := http.NewServeMux()
	routers.Routes(r, db)

	corsHandler := handlers.CorsMiddleware(r)

	fmt.Println("Server working on port 8090...")
	err2 := http.ListenAndServe(":8090", corsHandler)
	if err2 != nil {
		return
	}

}
