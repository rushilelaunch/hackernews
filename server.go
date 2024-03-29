package main

import (
	"github.com/go-chi/chi/v5"
	generated "github.com/rushilelaunch/hackernews/graph"
	"github.com/rushilelaunch/hackernews/internal/auth"
	database "github.com/rushilelaunch/hackernews/internal/pkg/db/mysql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &generated.Resolver{}}))
	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
