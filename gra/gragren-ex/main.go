package main

import (
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"gragen-ex/graph"
	"gragen-ex/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := defaultPort

	// 创建 GraphQL server
	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))
	srv.AddTransport(transport.POST{}) // 支持 POST 请求
	srv.AddTransport(transport.GET{})  // 支持 GET 请求
	srv.AddTransport(transport.Websocket{})
	http.Handle("/graphql", srv)
	http.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
