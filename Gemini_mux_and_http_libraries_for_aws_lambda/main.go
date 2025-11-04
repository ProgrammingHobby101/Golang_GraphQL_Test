// https://www.google.com/search?q=can+I+use+%22net%2Fhttp%22+with+lambda&rlz=1C1RXQR_enUS1067US1069&gs_lcrp=EgZjaHJvbWUyBggAEEUYOTIHCAEQIRigATIHCAIQIRigATIHCAMQIRigATIHCAQQIRigATIHCAUQIRigATIHCAYQIRiPAjIHCAcQIRiPAtIBCDc5NjZqMGo3qAIAsAIA&sourceid=chrome&ie=UTF-8&udm=50&ved=2ahUKEwjF1d-C2NSQAxWtMlkFHchiPAcQ0NsOegQIHhAA&aep=10&ntc=1&mstk=AUtExfB2hr32okaOUa757LzurBIfk44poTLekm_NO8BjgLQHA26Po3zsTGWI77rh5xGF_nvsUhdXxTqg-RIOi4Pu1HQ59X7b_Vg440bRQPezA5SA1WN-OTWYpTTse-9adHe1Lwe07SvS-xnXO5h3mAJjRwlQB5yVO_tc7UfX3PacZqdU0jTJkNiwWgXfVqs6Mo5oN_3rXOPDfMi-guxpbnNFue14OwHREc3HD1IfbfsJa9jBySox7Qljf3gE_PyMIMmao80eTsQ7fy7Ee1-1eanYcwewHUOTdFKysIGT0zN3V1U_9ywl6Iak1iPitbbPfRoj82gP6zbYoh7A1w&csuir=1&mtid=qPIHaaGwA7jR5NoP1ZDrgQM

package main

import (
	"fmt"
	"log"
	"net/http"

	//https://medium.com/@youlserf.cardenas/mastering-graphql-in-go-a-geeky-adventure-with-star-wars-and-pok%C3%A9mon-5db63f852fa0
	"graphql/schema"

	"github.com/graphql-go/handler"
	//"github.com/aws/aws-lambda-go/events"
	//"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	//99designs_gQLgen imports
)

//var adapter *httpadapter.HandlerAdapterV2

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

//	func handleGraphiQL() events.LambdaFunctionURLResponse {
//		return events.LambdaFunctionURLResponse{StatusCode: 200, Body: "Enjoy graphiQL."}
//		//return.
//	}
func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Healthy World! Watching the Tamron Hall Show.")
	//return events.LambdaFunctionURLResponse{StatusCode: 200, Body: message}
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from a Go Lambda!")
}

// func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
// 	path := req.RequestContext.HTTP.Path
// 	httpMethod := req.RequestContext.HTTP.Method
// 	//apiKey := req.Headers["x-api-key"]

// 	var response events.APIGatewayV2HTTPResponse

// 	switch path {
// 	case "/api/graphQL":
// 		fmt.Printf("The HTTP method in the /api/graphQL  path is: %s\n", httpMethod)
// 		PostCompare := "POST"
// 		if httpMethod == PostCompare {
// 			response = handleGraphQL(apiKey)
// 		} else {
// 			return events.APIGatewayV2HTTPResponse{
// 				StatusCode: 400,
// 				Body:       string("Method not allowed. CODE:400.1"), // Explicitly convert the untyped string constant
// 				// Other fields like Headers, Cookies, IsBase64Encoded can be added
// 			}, nil //return error in the "events.LambdaFunctionURLResponse" struct, don't return in this error field.

//			}
//		// case "/api/graphiql":
//		// 	response = handleGraphiQL()
//		// case "/api/health":
//		// 	response = handleHealth()
//		default:
//			response = events.APIGatewayV2HTTPResponse{StatusCode: 404, Body: "Not Found"}
//		}
//		//func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
//		//return adapter.ProxyWithContext(ctx, req)
//		return response, nil // This is the line originally for the "handler" function.
//	}
func main() {
	// Create a standard Go ServeMux and register your handler
	//router := mux.NewRouter()
	//router.HandleFunc("/", playground.Handler("GraphQL playground", "/api/graphQL"))
	// router.HandleFunc("/api/hello", myHandler)
	// router.HandleFunc("/api/health", handleHealth)
	// router.HandleFunc("/api/graphQL", handleGraphQL)
	// err := http.ListenAndServe(":3001", router)
	// if err != nil {
	// 	log.Fatal("ListenAndServe error: ", err)
	// }
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	// http.Handle("/api/hello", myHandler)
	// http.Handle("/api/health", handleHealth)
	// http.Handle("/api/graphQL", handleGraphQL)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

	// Wrap the mux in the httpadapter for V2 API Gateway events
	//adapter = httpadapter.NewV2(mux)

	// Start the Lambda runtime with the adapter's proxy method
	//lambda.Start(handler)
}
