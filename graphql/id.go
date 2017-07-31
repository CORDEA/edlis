package graphql

import (
	"github.com/graphql-go/graphql"
	"edlis/models"
)

var Id = graphql.NewObject(graphql.ObjectConfig{
	Name: "Id",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if id, ok := p.Source.(models.Id); ok {
					return id.Id, nil
				}
				return nil, nil
			},
		},
	},
})
