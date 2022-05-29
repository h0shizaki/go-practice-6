package gql

import (
	"go-gql/postgres"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	Db *postgres.Db
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {

	name, ok := p.Args["name"].(string)

	if ok {
		users := r.Db.GetUserByName(name)
		return users, nil
	}

	return nil, nil
}
