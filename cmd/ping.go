package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thecrogers/echo-micro/server"
	"github.com/thecrogers/echo-micro/server/handlers"
)

func init() {
	rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Simple http server that returns a 200 for any request.",
	Long:  `Any HTTP request to the confiugred port will send 200 response back. `,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("starting ping command")
		server, err := server.InitServer(rootCmd)
		if err != nil {
			return err
		}
		server.Router.PathPrefix("/").HandlerFunc(handlers.PingHandler)
		server.StartServer()
		return nil
	},
}
