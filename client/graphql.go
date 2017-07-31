package client

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"log"
	"net/http"
)

type GraphQl struct {
	IsTesting bool
	Client    MongoDb
}

func NewGraphQl(client MongoDb, isTesting bool) GraphQl {
	return GraphQl{
		IsTesting: isTesting,
		Client:    client,
	}
}

func (g *GraphQl) Handler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query()["query"]) == 0 {
			return
		}
		result := g.executeQuery(r.URL.Query()["query"][0], g.schema())
		json.NewEncoder(w).Encode(result)
	}
}

func ParseTestJson(name string, model interface{}) {
	data, err := ioutil.ReadFile("./resources/" + name + ".json")
	if err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(data, model); err != nil {
		log.Fatalln(err)
	}
}

func (g *GraphQl) schema() graphql.Schema {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Queries(g.IsTesting, g.Client),
		Mutation: Mutations(g.IsTesting, g.Client),
	})
	return schema
}

func (g *GraphQl) executeQuery(q string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: q,
	})
	return result
}
