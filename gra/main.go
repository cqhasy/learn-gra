package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"gra/beauter"
)

func main() {
	query := `{ hello(name: "Alice") }`
	s := beauter.NewService()
	r := beauter.NewResolver(s)
	sh, _ := beauter.NewSchema(r)
	result := graphql.Do(graphql.Params{
		Schema:        sh,
		RequestString: query,
	})
	j, _ := json.Marshal(result)
	fmt.Println(string(j))

}
