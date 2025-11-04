package schema

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/graphql-go/graphql"
)

var client = resty.New()

// Define a GraphQL type for SWAPI people
var personType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"height": &graphql.Field{
			Type: graphql.String,
		},
		"mass": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// Define a GraphQL type for PokeAPI Pokémon
var pokemonType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Pokemon",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"height": &graphql.Field{
			Type: graphql.Int,
		},
		"weight": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// Define the root query
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"person": &graphql.Field{
			Type: personType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(string)
				if !ok {
					return nil, nil
				}
				resp, err := client.R().Get("https://swapi.dev/api/people/" + id + "/")
				if err != nil {
					return nil, err
				}
				var result map[string]interface{}
				err = json.Unmarshal(resp.Body(), &result)
				if err != nil {
					return nil, err
				}
				return result, nil
			},
		},
		"pokemon": &graphql.Field{
			Type: pokemonType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name, ok := p.Args["name"].(string)
				if !ok {
					return nil, nil
				}
				resp, err := client.R().Get("https://pokeapi.co/api/v2/pokemon/" + name + "/")
				if err != nil {
					return nil, err
				}
				var result map[string]interface{}
				err = json.Unmarshal(resp.Body(), &result)
				if err != nil {
					return nil, err
				}
				return result, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})
