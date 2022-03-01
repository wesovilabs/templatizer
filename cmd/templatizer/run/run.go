package run

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wesovilabs/templatizer/internal/action"
)

var debug bool

var repoPath, username, password, templateMode, targetDir, inputPath string

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Create a project from a given template",
		Long: `If flag --input is provided will create the repository from the template otherwise
		It will create a param file to be filled`,
		Example: "templatizer run --source github.com/ivancorrales/go-rest-template",
		PreRun: func(cmd *cobra.Command, args []string) {
			log.SetFormatter(&log.TextFormatter{})
			if debug {
				log.Info("Debug logs enabled...")
				log.SetLevel(log.DebugLevel)
			}
		},
		Run: run,
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

func run(cmd *cobra.Command, args []string) {
	log.Info("running templatizer...")
	for i := range args {
		println(args[i])
	}
	println()
	if repoPath == "" {
		log.Error("missing required flag 'from'")
		log.Info("Please run `templatizer help")
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
