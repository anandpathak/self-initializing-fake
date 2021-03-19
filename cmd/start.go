package cmd

import (
	"github.com/spf13/cobra"
	server2 "self_initializing_fake/internal/server"
)

func start() *cobra.Command {

	return &cobra.Command{
		Use:   "start",
		Short: "start the http-server",
		Long:  `runs the http-server on port 8080 default.`,
		Run: func(cmd *cobra.Command, args []string) {
			server2.Start()
		},
	}
}
