package gql

import (
	"log"

	"github.com/graphql-go/graphql"
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		log.Printf("Unexpected errors inside ExecuteQuery: %v \n", result.Errors)
	}

	return result
}
