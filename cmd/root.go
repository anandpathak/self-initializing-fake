package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "self-initializing-fake",
		Short: "setup mock server",
		Long:  ``,
	}
)

func init() {
	setDefaultCommandIfNonePresent()
	rootCmd.AddCommand(start())
}

func Execute( ) error{
	return  rootCmd.Execute()
}

func setDefaultCommandIfNonePresent() {
	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "start"})
	}
}