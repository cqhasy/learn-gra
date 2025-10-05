package beauter

import "github.com/graphql-go/graphql"

func NewSchema(r *Resolver) (graphql.Schema, error) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hello": r.HelloField(),
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
}
