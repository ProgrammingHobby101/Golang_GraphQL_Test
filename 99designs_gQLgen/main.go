package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
	"gitlab.com/pragmaticreviews/graph-server/graph"
	//below are imports for mux and HandleFunc  https://www.google.com/search?q=can+I+use+%22net%2Fhttp%22+with+lambda&rlz=1C1RXQR_enUS1067US1069&gs_lcrp=EgZjaHJvbWUyBggAEEUYOTIHCAEQIRigATIHCAIQIRigATIHCAMQIRigATIHCAQQIRigATIHCAUQIRigATIHCAYQIRiPAjIHCAcQIRiPAtIBCDc5NjZqMGo3qAIAsAIA&sourceid=chrome&ie=UTF-8&udm=50&ved=2ahUKEwjF1d-C2NSQAxWtMlkFHchiPAcQ0NsOegQIHhAA&aep=10&ntc=1&mstk=AUtExfB2hr32okaOUa757LzurBIfk44poTLekm_NO8BjgLQHA26Po3zsTGWI77rh5xGF_nvsUhdXxTqg-RIOi4Pu1HQ59X7b_Vg440bRQPezA5SA1WN-OTWYpTTse-9adHe1Lwe07SvS-xnXO5h3mAJjRwlQB5yVO_tc7UfX3PacZqdU0jTJkNiwWgXfVqs6Mo5oN_3rXOPDfMi-guxpbnNFue14OwHREc3HD1IfbfsJa9jBySox7Qljf3gE_PyMIMmao80eTsQ7fy7Ee1-1eanYcwewHUOTdFKysIGT0zN3V1U_9ywl6Iak1iPitbbPfRoj82gP6zbYoh7A1w&csuir=1&mtid=qPIHaaGwA7jR5NoP1ZDrgQM
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
