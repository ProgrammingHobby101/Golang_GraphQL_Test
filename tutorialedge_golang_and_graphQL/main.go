// source code; https://tutorialedge.net/golang/go-graphql-beginners-tutorial/

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
)

var tutorials []Tutorial

type Tutorial struct {
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

var schema graphql.Schema

func populate() []Tutorial {
	author := &Author{Name: "Elliot Forbes", Tutorials: []int{1, 2, 3}}
	tutorial := Tutorial{
		ID:     0,
		Title:  "Go GraphQL Tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "First Comment"},
		},
	}
	tutorial2 := Tutorial{
		ID:     1,
		Title:  "Go GraphQL Tutorial - Part 2",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "Second Comment"},
		},
	}
	tutorial3 := Tutorial{
		ID:     2,
		Title:  "Go GraphQL Tutorial - Part 3",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "Third Comment"},
		},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	tutorials = append(tutorials, tutorial2)
	tutorials = append(tutorials, tutorial3)

	return tutorials
}

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	},
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        tutorialType,
			Description: "Create a new Tutorial",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				tutorial := Tutorial{
					Title: params.Args["title"].(string),
				}
				tutorials = append(tutorials, tutorial)
				return tutorial, nil
			},
		},
	},
})

// GraphQLRequest represents the structure of a typical GraphQL request
type GraphQLRequest struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

func main() {
	my_init()
	lambda.Start(HandleRequest)
}

func my_init() { //only run on cold starts
	tutorials = populate() //only run on cold starts

	// Schema
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get Tutorial By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					// Find tutorial
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: mutationType,
	}

	var err error //for the following line.
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
}

func handleHealth() events.LambdaFunctionURLResponse {
	message := "Hello Healthy World2! Watching the Tamron Hall Show."
	return events.LambdaFunctionURLResponse{StatusCode: 200, Body: message}
}

// func handleGraphQL(apiKey string) events.LambdaFunctionURLResponse { //Overriden by "func GraphQLEntryPoint"
// 	greeting := "Hi, Earthling!"
// 	if apiKey != "valid_key" { // Optional: Validate API Key
// 		return events.LambdaFunctionURLResponse{StatusCode: 401, Body: "Unauthorized"}
// 	}
// 	return events.LambdaFunctionURLResponse{StatusCode: 200, Body: greeting + " you are allowed"}
// }

func GraphQLEntryPoint(request events.LambdaFunctionURLRequest) events.LambdaFunctionURLResponse {
	apiKey := request.Headers["x-api-key"]
	//greeting := "Hi, Earthling!"
	if apiKey != "valid_key" { // Optional: Validate API Key
		return events.LambdaFunctionURLResponse{StatusCode: 401, Body: "Unauthorized"}
	}
	// // Mutation manual test
	// query := `
	// 		    mutation {
	// 				create(title: "medium blog #3") {
	// 					title
	// 					id
	// 				}
	// 			}
	// 	`
	query1 := request.Body
	fmt.Printf("Request.Body was; %s \n", query1)
	params := graphql.Params{Schema: schema, RequestString: query1}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON1, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON1)
	return events.LambdaFunctionURLResponse{StatusCode: 200, Body: "completed query/mutation: " + string(rJSON1)}

	// // // Query manual test
	// // query = `
	// // 		{
	// // 			list {
	// // 				id
	// // 				title
	// // 			}
	// // 		}
	// // 	`
	// query2 := request.Body
	// params = graphql.Params{Schema: schema, RequestString: query2}
	// r = graphql.Do(params)
	// if len(r.Errors) > 0 {
	// 	log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	// }
	// rJSON2, _ := json.Marshal(r)
	// fmt.Printf("%s \n", rJSON2)
	// return events.LambdaFunctionURLResponse{StatusCode: 201, Body: greeting + " and you are allowed. GraphQL Mutate(Create) response: " + string(rJSON1) + "GraphQL Query(list tutorials) response:" + string(rJSON2)}
}
func HandleRequest(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	path := request.RequestContext.HTTP.Path
	httpMethod := request.RequestContext.HTTP.Method
	//apiKey := request.Headers["x-api-key"]

	var response events.LambdaFunctionURLResponse

	switch path {
	case "/":
		fmt.Printf("The HTTP method in the /api/graphQL  path is: %s\n", httpMethod)
		// PostCompare := "POST"
		// if httpMethod == PostCompare {
		// 	// response = handleGraphQL(apiKey)
		// } else {
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       string("Method not allowed. CODE:400.1"), // Explicitly convert the untyped string constant
			// Other fields like Headers, Cookies, IsBase64Encoded can be added
		}, nil //return error in the "events.LambdaFunctionURLResponse" struct, don't return in this error field.

		//}// when I uncomment this, also uncomment the Method check in this switch-case.
	case "/api/graphQL":
		apiKey := request.Headers["x-api-key"]
		//greeting := "Hi, Earthling!"
		if apiKey != "valid_key" { // Optional: Validate API Key
			return events.LambdaFunctionURLResponse{StatusCode: 401, Body: "Unauthorized"}, nil
		}

		var gqlRequest GraphQLRequest
		if err := json.Unmarshal([]byte(request.Body), &gqlRequest); err != nil {
			// http.Error(w, "Error unmarshaling JSON", http.StatusBadRequest)
			// return
			return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Error unmarshaling JSON"}, nil
		}

		// Now gqlRequest.Query contains the GraphQL query string
		// gqlRequest.Variables contains any variables passed
		// gqlRequest.OperationName contains the operation name (if specified)

		fmt.Printf("Received GraphQL Query: %s\n", gqlRequest.Query)
		fmt.Printf("Variables: %v\n", gqlRequest.Variables)

		// ... process the GraphQL query ...

		// w.WriteHeader(http.StatusOK)
		// w.Write([]byte("Query received and processed (placeholder)"))
		//responseBodyPart1 := "Query received and processed (placeholder). Received GraphQL Query was: %s\n" + gqlRequest.Query
		//query1 := request.Body
		//fmt.Printf("Request.Body was; %s \n", query1)
		// Query manual test
		// querytest := `
		// 		{
		// 			list {
		// 				id
		// 				title
		// 			}
		// 		}
		// 	`
		params := graphql.Params{Schema: schema, RequestString: gqlRequest.Query} // gqlRequest.Query (this is for non-hardcoded queries), querytest (this is for hard-coded queries)
		r := graphql.Do(params)

		// // Marshal the User struct into a JSON byte slice
		graphqlResultJSON_Data, err := json.Marshal(r.Data)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "tried to json.Marshal(r.Data) but got error; " + err.Error()}, nil
		}
		// responseBodyPart2, ok := r.Data.(string)
		// if ok {
		// 	// reresponseBodyPart2 now holds the string value
		// 	fmt.Println(responseBodyPart2)
		// 	responseBodyPart1 = responseBodyPart1 + " . The graphql query result is: " + responseBodyPart2
		// } else {
		// 	// Handle the case where the assertion fails
		// 	fmt.Println("Value is not a string")
		// 	return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "r.Data.(string) was not a string because it was "}, nil
		// }

		// check for null value in the resulting query
		if string(graphqlResultJSON_Data) == "null" {
			return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Error graphql result was: " + string(graphqlResultJSON_Data) + " , gqlRequest.Query: " + gqlRequest.Query}, nil
		}
		return events.LambdaFunctionURLResponse{StatusCode: 200, Body: "graphql result is: " + string(graphqlResultJSON_Data)}, nil

		//old working code is below!

		// query1 := request.Body
		// fmt.Printf("Request.Body was; %s \n", query1)
		// params := graphql.Params{Schema: schema, RequestString: query1}
		// r := graphql.Do(params)
		// return events.LambdaFunctionURLResponse{StatusCode: 200, Body: "Query received and processed (placeholder). Received GraphQL Query: %s\n" + gqlRequest.Query + "  "}, nil
		// ... your GraphQL execution logic resulting in 'r' ...
		//var r *graphql.Result // Assume 'r' is populated here

		// Example: Create a dummy result for demonstration
		// r = &graphql.Result{
		// 	Data: map[string]interface{}{
		// 		"query": r.Data,
		// 	},
		// }

		// jsonBytes, err := json.Marshal(r)
		// if err != nil {
		// 	fmt.Println("Error marshalling to JSON:", err)
		// 	//return
		// 	return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Error occurred." + err.Error()}, nil
		// }

		// jsonString := string(jsonBytes)
		// fmt.Println(jsonString)

		// if len(r.Errors) > 0 {
		// 	log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
		// 	return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Error occurred." + string(jsonString)}, nil
		// }
		// rJSON1, _ := json.Marshal(r)
		// fmt.Printf("%s \n", rJSON1)
		// return events.LambdaFunctionURLResponse{StatusCode: 200, Body: "completed query/mutation: " + string(rJSON1)}, nil

		//TEST november 7, 2025
		// // // Mutation manual test
		// // query := `
		// // 		mutation {
		// // 			create(title: "Hello Lambda World") {
		// // 				title
		// // 			}
		// // 		}
		// // 	`
		// // var message string
		// query1 := request.Body
		// // message = "query recieved " + query1
		// // Marshal the User struct into a JSON byte slice
		// jsonData, err := json.Marshal(query1)
		// if err != nil {
		// 	fmt.Println("Error marshalling JSON:", err)
		// }

		// // Convert the byte slice to a string for printing
		// fmt.Println(string(jsonData)) // Output: {"Name":"Alice","Age":30,"IsAdmin":true}

		// return events.LambdaFunctionURLResponse{StatusCode: 200, Body: string(jsonData)}, nil

		// fmt.Printf("request.Body was; %s \n", query1)
		// params := graphql.Params{Schema: schema, RequestString: query1}
		// r := graphql.Do(params)
		// if len(r.Errors) > 0 {
		// 	//log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
		// 	return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Error detected in /api/graphQL ."}, nil
		// }
		// rJSON1, _ := json.Marshal(r)
		// fmt.Printf("%s \n", rJSON1)
		// return events.LambdaFunctionURLResponse{StatusCode: 200, Body: "completed query/mutation: " + string(rJSON1)}, nil

		// //response = GraphQLEntryPoint(request)
	case "/api/graphiql":
		response = handleHealth()
	// case "/api/graphiql":
	// 	response = handleGraphiQL()
	case "/api/health":
		response = handleHealth()
	default:
		response = events.LambdaFunctionURLResponse{StatusCode: 404, Body: "Not Found"}
	}

	return response, nil
}
