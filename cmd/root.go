package cmd

import (
	"log"
	"net/http"

	"github.com/Noah-Heil/golang-relay-starter-kit/data"
	"github.com/graphql-go/handler"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{Use: "app"}

func init() {
	RootCmd.AddCommand(seedCmd)
	RootCmd.AddCommand(serveCmd)
}

var seedCmd = &cobra.Command{
	Use:   "devdockerseed",
	Short: "seed db for dev in docker env",
	Long:  `Seeds the db with fake data for development in a docker container`,
	Run: func(cmd *cobra.Command, args []string) {
		data.SeedDockerdb()
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start up the server",
	Long:  `starts the server so we can use it`,
	Run: func(cmd *cobra.Command, args []string) {
		StartMainServer()
	},
}

func StartMainServer() {
	// simplest relay-compliant graphql server HTTP handler
	h := handler.New(&handler.Config{
		Schema: &data.Schema,
		Pretty: true,
	})

	fs := http.FileServer(http.Dir("static"))

	// create graphql endpoint
	http.Handle("/graphql", h)

	// create graphiql endpoint
	http.Handle("/", fs)

	// serve!
	port := ":8080"
	log.Printf(`GraphQL server starting up on http://localhost%v`, port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
