// source code: https://pkg.go.dev/github.com/mnmtanish/go-graphiql@v0.0.0-20160921055525-cef5a61bd62b#section-readme

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mnmtanish/go-graphiql"
)

func myGraphQLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from a Go Lambda! pkg_go_dev part#2")
}

func main() {
	// Retrieve the value of the "PORT" environment variable
	port := os.Getenv("PORT")
	if port == "" { //local port
		fmt.Println("PORT environment variable not set. Using default port.")
		port = ":8080" // Default port if not set
	}
	fmt.Printf("Listening on port: %s\n", port)

	http.HandleFunc("/api/health", myGraphQLHandler)
	http.HandleFunc("/api/graphql", graphiql.ServeGraphiQL)
	http.ListenAndServe(port, nil)
}
