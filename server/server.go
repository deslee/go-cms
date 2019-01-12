package main

import (
	"database/sql"
	"github.com/deslee/cms/database/repository"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"github.com/99designs/gqlgen/handler"
	"github.com/deslee/cms"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		panic(err)
	}

	repo := repository.DbRepo{DB: db}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(cms.NewExecutableSchema(cms.Config{Resolvers: &cms.Resolver{Repo: repo}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
