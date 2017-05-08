package data

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var userType *graphql.Object

var queryType *graphql.Object

var Schema graphql.Schema

func init() {

	/**
	 * We get the node interface and field from the Relay library.
	 *
	 * The first method defines the way we resolve an ID to its object.
	 * The second defines the way we resolve an object to its GraphQL type.
	 */
	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A person who uses our app",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"creation_time": &graphql.Field{
				Type: graphql.String,
			},
			"update_time": &graphql.Field{
				Type: graphql.String,
			},
			"deletion_time": &graphql.Field{
				Type: graphql.String,
			},
			"first_name": &graphql.Field{
				Type: graphql.String,
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	/**
	 * This is the type that will be the root of our query,
	 * and the entry point into our schema.
	 */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// Add you own root fields here
			"user": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetUser("3"), nil
				},
			},
		},
	})

	/**
	 * This is the type that will be the root of our mutations,
	 * and the entry point into performing writes in our schema.
	 */
	//	mutationType := graphql.NewObject(graphql.ObjectConfig{
	//		Name: "Mutation",
	//		Fields: graphql.Fields{
	//			// Add you own mutations here
	//		},
	//	})

	/**
	* Finally, we construct our schema (whose starting query type is the query
	* type we defined above) and export it.
	 */
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		panic(err)
	}
}
