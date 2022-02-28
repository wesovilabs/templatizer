package main

import (
	"os"

	"github.com/wesovilabs/templatizer/internal/action"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debug bool
var repoPath, username, password, templateMode, targetDir, inputPath string

func command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "templatizer",
		Short: `Generate a json file with empty variables to be filled.`,
		Long: `

		`,
		Example: "init --from github.com/ivancorrales/go-rest-template",
		PreRun: func(cmd *cobra.Command, args []string) {
			log.SetFormatter(&log.TextFormatter{})
			if debug {
				log.Info("Debug logs enabled...")
				log.SetLevel(log.DebugLevel)
			}
		},
		Run: runTemplatizer,
	}

	cmd.PersistentFlags().BoolVar(&debug, "verbose", false,
		"verbose logging")
	cmd.PersistentFlags().StringVar(&repoPath, "source", "",
		"path to the repo i.e. github.com/ivancorrales/go-rest-template)")
	cmd.PersistentFlags().StringVarP(&username, "username", "u", "",
		"user handle")
	cmd.PersistentFlags().StringVarP(&password, "password", "p", "",
		"user secret for the provided username")
	cmd.PersistentFlags().StringVarP(&templateMode, "mode", "m", "goTemplate",
		"template mode used tod efine the variables. ")
	cmd.PersistentFlags().StringVarP(&inputPath, "input", "i", "",
		"path to the input files that contains the variables")
	cmd.PersistentFlags().StringVar(&targetDir, "target", "",
		"Path to the folder in which the repository will be created")
	return cmd
}

func main() {
	rootCmd := command()
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("unexpected error: %s", err)
		os.Exit(1)
	}
}

func runTemplatizer(cmd *cobra.Command, args []string) {
	log.Info("running templatizer...")
	if repoPath == "" {
		log.Error("missing required flag 'from'.\n Example: 'templatizer --from github.com/organization/repository.git'")
		os.Exit(1)
	}
	options := []action.Option{
		action.WithRepoPath(repoPath),
	}
	if inputPath != "" {
		options = append(options, action.WithVariables(inputPath))
	}
	if username != "" {
		options = append(options, action.WithUsername(username))
	}
	if password != "" {
		options = append(options, action.WithPassword(password))
	}
	if templateMode != "" {
		options = append(options, action.WithTemplateMode(templateMode))
	}
	if targetDir != "" {
		options = append(options, action.WithTargetDir(targetDir))
	}
	err := action.New(options...).Execute()
	if err != nil {
		os.Exit(1)
	}
}
