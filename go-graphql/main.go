package main

import (
	"fmt"
	"graphql/schema"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/handler"
)

func handleGraphQL(w http.ResponseWriter, r *http.Request) { //old return type was;  events.APIGatewayV2HTTPResponse
	greeting := "Hi, Lambda Earthling!"
	apiKey := r.Header.Get("x-api-key")
	if apiKey != "valid_key" { // Optional: Validate API Key
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unauthorized."))
		return
		// return events.APIGatewayV2HTTPResponse{StatusCode: 401, Body: "Unauthorized"}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(greeting + " you are allowed with mux"))
	// return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: greeting + " you are allowed with mux"}
}
func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Healthy World! Watching the Tamron Hall Show.")
	//return events.LambdaFunctionURLResponse{StatusCode: 200, Body: message}
}
func myHandler() {
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/api/graphiql", h)
	http.HandleFunc("/api/health", handleHealth)
	http.HandleFunc("/api/graphql", handleGraphQL)
	port := os.Getenv("PORT") // Retrieve the value of the "PORT" environment variable
	if port == "" {
		port = "8080" // Provide a default value if the environment variable is not set
	}
	http.ListenAndServe(port, nil)
}

func main() {

	// Start the Lambda runtime with the adapter's proxy method
	lambda.Start(myHandler)
}
