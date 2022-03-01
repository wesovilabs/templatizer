package main

import (
	"os"

	"github.com/wesovilabs/templatizer/cmd/templatizer/help"
	"github.com/wesovilabs/templatizer/cmd/templatizer/run"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "templatizer [cmd]",
	}
	runCmd := run.New()
	helpCmd := help.New()
	rootCmd.SetHelpCommand(helpCmd)
	rootCmd.AddCommand(helpCmd, runCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("unexpected error: %s", err)
		os.Exit(1)
	}
}
