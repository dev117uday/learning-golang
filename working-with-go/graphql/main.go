package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type Tutorial struct {
	Id       int
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

func populate() []Tutorial {
	Author := &Author{Name: "Elliot Forbes", Tutorials: []int{1, 2}}
	tutorial := Tutorial{
		Id:     1,
		Title:  "Go GraphQL Tutorial",
		Author: *Author,
		Comments: []Comment{
			{Body: "First Comment"},
		},
	}
	tutorial2 := Tutorial{
		Id:     2,
		Title:  "Go GraphQL Tutorial - Part 2",
		Author: *Author,
		Comments: []Comment{
			{Body: "Second Comment"},
		},
	}

	var Tutorials []Tutorial
	Tutorials = append(Tutorials, tutorial)
	Tutorials = append(Tutorials, tutorial2)

	return Tutorials
}

var AuthorType = graphql.NewObject(
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
			"Body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var TutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.Int,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Author": &graphql.Field{
				Type: AuthorType,
			},
			"Comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	},
)

func main() {

	Tutorials := populate()

	// Schema
	fields := graphql.Fields{
		"Tutorial": &graphql.Field{
			Type:        TutorialType,
			Description: "Get Tutorial By Id",
			Args: graphql.FieldConfigArgument{
				"Id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				Id, ok := p.Args["Id"].(int)
				if ok {

					for _, Tutorial := range Tutorials {
						if int(Tutorial.Id) == Id {
							return Tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(TutorialType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return Tutorials, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			Tutorial (Id:1) {
				Title
				Author {
					Name
				}
				Comments {
					Body
				}
			}
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
