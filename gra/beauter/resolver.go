package beauter

import "github.com/graphql-go/graphql"

//解析器

type Resolver struct {
	service *Service
}

func NewResolver(service *Service) *Resolver {
	return &Resolver{service}
}

func (resolver *Resolver) HelloField() *graphql.Field {
	return &graphql.Field{
		Name: "hello",
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			}},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name := p.Args["name"].(string)
			return resolver.service.SayHello(name), nil
		},
	}
}
