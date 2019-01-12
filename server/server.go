package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/handler"
	"github.com/deslee/cms"
	"github.com/deslee/cms/data"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"strings"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("sqlite3", "database.sqlite?_loc=auto")
	if err != nil {
		panic(err)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle(
		"/query",
		withAuth(handler.GraphQL(cms.NewExecutableSchema(cms.Config{Resolvers: &cms.Resolver{DB: db}}))),
	)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) == 2 {
			if auth[0] == "Bearer" {
				token := auth[1]
				ctx, err := data.ParseTokenToContext(token, r.Context())
				if err != nil {
					http.Error(w, "authorization failed", http.StatusUnauthorized)
					return
				}
				r = r.WithContext(ctx)
			} else {
				http.Error(w, "invalid scheme", http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}