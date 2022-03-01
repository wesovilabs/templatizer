/*
Copyright 2022 Iván Corrales
*/
package help

import (
	"strings"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "help [command]",
		Short: `Help about any command.`,
		Long: `Help provides help for any command in the application.
Simply type templatizer help [path to command] for full details.`,
		Example: "templatizer help run",
		Run:     run,
	}
}

// RunHelp checks given arguments and executes command.
func run(cmd *cobra.Command, args []string) {
	foundCmd, _, err := cmd.Root().Find(args)
	switch {
	case foundCmd == nil:
		cmd.Printf("Unknown help topic %#q.\n", args)
		if usageErr := cmd.Root().Usage(); usageErr != nil {
			panic(usageErr)
		}
		return
	case err != nil:
		cmd.Println(err)
		argsString := strings.Join(args, " ")
		matchedMsgIsPrinted := false
		for _, foundCmd := range foundCmd.Commands() {
			if strings.Contains(foundCmd.Short, argsString) {
				if !matchedMsgIsPrinted {
					cmd.Printf("Matchers of string '%s' in short descriptions of commands: \n", argsString)
					matchedMsgIsPrinted = true
				}
				cmd.Printf("  %-14s %s\n", foundCmd.Name(), foundCmd.Short)
			}
		}
		if !matchedMsgIsPrinted {
			if err := cmd.Root().Usage(); err != nil {
				panic(err)
			}
		}
		return
	default:
		if len(args) == 0 {
			foundCmd = cmd
		}
		helpFunc := foundCmd.HelpFunc()
		helpFunc(foundCmd, args)
	}
}
