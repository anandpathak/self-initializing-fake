package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "self-initializing-fake",
		Short: "setup fake server",
		Long:  ``,
	}
)

func init() {
	config := initConfig()
	setDefaultCommandIfNonePresent(config)

	rootCmd.AddCommand(start(config))
}

func Execute( ) error{
	return  rootCmd.Execute()
}

func setDefaultCommandIfNonePresent(c config) {
	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "start", "-a", c.setupServerPort, "-f", c.fakeServerPort})
	}
}