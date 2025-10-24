package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handleHealth() events.LambdaFunctionURLResponse {
	message := "Hello World!"
	return events.LambdaFunctionURLResponse{StatusCode: 200, Body: message}
}

func handleGraphQL(apiKey string) events.LambdaFunctionURLResponse {
	greeting := "Hi!"
	if apiKey != "valid_key" { // Optional: Validate API Key
		return events.LambdaFunctionURLResponse{StatusCode: 401, Body: "Unauthorized"}
	}
	return events.LambdaFunctionURLResponse{StatusCode: 200, Body: greeting + " you are allowed"}
}

func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	path := request.RequestContext.HTTP.Path
	httpMethod := request.RequestContext.HTTP.Method
	apiKey := request.Headers["x-api-key"]

	var response events.LambdaFunctionURLResponse

	switch path {
	case "/api/graphQL":
		fmt.Printf("The HTTP method in the /api/graphQL  path is: %s\n", httpMethod)
		PostCompare := "POST"
		if httpMethod == PostCompare {
			response = handleGraphQL(apiKey)
		} else {
			return events.LambdaFunctionURLResponse{
				StatusCode: 404,
				Body:       string("Not Found."), // Explicitly convert the untyped string constant
				// Other fields like Headers, Cookies, IsBase64Encoded can be added
			}, errors.New("input value 3 is not allowed") // Return an error // throw error

		}

	case "/api/health":
		response = handleHealth()
	default:
		response = events.LambdaFunctionURLResponse{StatusCode: 404, Body: "Not Found"}
	}

	return response, nil
}
