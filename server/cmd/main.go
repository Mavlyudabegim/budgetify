package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	authhandlers "server/internal/auth/auth-handlers"

	"github.com/gorilla/mux"
	"server/internal/auth"
	"server/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	queries := database.New(db)
	authSvc := auth.AuthService{Q: queries}

	r := mux.NewRouter()
	r.HandleFunc("/register", authhandlers.RegisterHandler(&authSvc)).Methods("POST")
	r.HandleFunc("/login", authhandlers.LoginHandler(&authSvc)).Methods("POST")
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
