package main

import (
	"log"
	"net/http"

	"todo-example/graph"
	"todo-example/graph/generated"

	"github.com/99designs/gqlgen/handler"
)

func main() {
	port := "8080"
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
