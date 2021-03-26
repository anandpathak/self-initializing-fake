package cmd

import (
	server2 "self_initializing_fake/internal/server"

	"github.com/spf13/cobra"
)

func start(conf config) *cobra.Command {
	var setupServerPort, fakeServerPort string
	var fakeServerTimeout int

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start the http-server",
		Long:  `runs the http-server on port 8080 default.`,
		Run: func(cmd *cobra.Command, args []string) {

			server2.Start(setupServerPort, fakeServerPort, fakeServerTimeout)
		},
	}
	startCmd.Flags().StringVarP(&setupServerPort, "setup-server-port", "s", conf.setupServerPort, "port to run the setup route")
	startCmd.Flags().StringVarP(&fakeServerPort, "fake-server-port", "f", conf.fakeServerPort, "port to run the fake route")
	startCmd.Flags().IntVarP(&fakeServerTimeout, "fake-server-timeout", "t", conf.fakeServerTimeout, "port to run the fake route")
	return startCmd
}
