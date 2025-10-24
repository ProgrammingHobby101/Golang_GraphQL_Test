package main

import (
	"context"
	//"encoding/json"//uncomment later?
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type AppSyncEvent struct {
	Info struct {
		FieldName string `json:"fieldName"`
	} `json:"info"`
	Arguments map[string]interface{} `json:"arguments"`
}

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func HandleRequest(ctx context.Context, event AppSyncEvent) (interface{}, error) {
	fmt.Printf("Received event: %+v\n", event)

	switch event.Info.FieldName {
	case "/getPost":
		id := event.Arguments["id"].(string)
		// Implement logic to fetch post by ID from a database or other source
		return Post{ID: id, Title: "Sample Post", Content: "This is a sample post content."}, nil
	case "/listPosts":
		// Implement logic to list all posts
		return []Post{{ID: "1", Title: "Post 1"}, {ID: "2", Title: "Post 2"}}, nil
	case "createPost":
		title := event.Arguments["title"].(string)
		content := event.Arguments["content"].(string)
		// Implement logic to create a new post
		return Post{ID: "new-id", Title: title, Content: content}, nil
	default:
		return nil, fmt.Errorf("unknown field: %s", event.Info.FieldName)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
