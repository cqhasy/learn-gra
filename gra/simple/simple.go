package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// 定义 User 数据结构
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 模拟数据库
var users = []User{
	{ID: "1", Name: "Alice", Age: 23},
	{ID: "2", Name: "Bob", Age: 25},
}

func main() {
	// 定义 GraphQL Object Type：User
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"name": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"age":  &graphql.Field{Type: graphql.Int},
		},
	})

	// 定义 Query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(userType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return users, nil
				},
			},
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					for _, u := range users {
						if u.ID == id {
							return u, nil
						}
					}
					return nil, nil
				},
			},
		},
	})

	// 定义 Mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name := p.Args["name"].(string)
					age, _ := p.Args["age"].(int)
					user := User{
						ID:   string(rune(len(users) + 1)),
						Name: name,
						Age:  age,
					}
					users = append(users, user)
					return user, nil
				},
			},
		},
	})

	// 构建 Schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatalf("创建 Schema 失败: %v", err)
	}

	// 创建 GraphQL handler（带 playground）
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // 打开 GraphQL Playground
	})
	// 注册路由
	http.Handle("/graphql", h)

	log.Println("GraphQL server running at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
