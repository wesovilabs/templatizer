package action

import (
	"fmt"
	"io/ioutil"

	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	"gopkg.in/yaml.v3"

	log "github.com/sirupsen/logrus"
)

type action struct {
	repoURL       string
	username      string
	password      string
	branch        string
	templateMode  string
	variablesPath string
	targetDir     string
	inputVars     map[string]interface{}
}

type Option func(*action)

func WithRepoPath(repoPath string) Option {
	return func(act *action) {
		act.repoURL = fmt.Sprintf("https://%s", repoPath)
	}
}

func WithUsername(username string) Option {
	return func(act *action) {
		act.username = username
	}
}

func WithTemplateMode(templateMode string) Option {
	return func(act *action) {
		act.templateMode = templateMode
	}
}

func WithPassword(password string) Option {
	return func(act *action) {
		act.password = password
	}
}

func WithBranch(branch string) Option {
	return func(act *action) {
		act.branch = branch
	}
}

func WithTargetDir(targetDir string) Option {
	return func(act *action) {
		act.targetDir = targetDir
	}
}

func WithVariables(inputVarsPath string) Option {
	return func(act *action) {
		if inputVarsPath == "" {
			return
		}
		yfile, err := ioutil.ReadFile(inputVarsPath)
		if err != nil {
			log.Fatalf("unexpected error reading the input file %s: '%s'", inputVarsPath, err)
		}
		if err := yaml.Unmarshal(yfile, &act.inputVars); err != nil {
			log.Fatalf("unexpected error unmarhsaling the input file %s: '%s'", inputVarsPath, err)
		}
		act.variablesPath = inputVarsPath

	}
}

func New(opts ...Option) *action {
	act := &action{
		templateMode:  "goTemplate",
		variablesPath: "tempaltize-vars.yml",
	}
	for _, opt := range opts {
		opt(act)
	}
	return act
}

func (act *action) Execute() error {
	var auth http.AuthMethod
	if act.username != "" || act.password != "" {
		auth = &http.BasicAuth{
			Username: act.username,
			Password: act.password,
		}
		log.Debug("- set repository credentials")
	}
	w, err := cloneRepositorty(act.repoURL, act.branch, auth)
	if err != nil {
		return err
	}
	repoFiles := fetchFilesContent(w)

	log.Debug("- extract variables from files in the repository.")
	variables := extract(repoFiles, act.templateMode)
	if len(variables) > 0 && act.inputVars == nil {
		variables.ToYAML(act.variablesPath)
		log.Warnf("The repository contains variables that need to be provided.\nPlease, set the values for file '%s' and invoke the command again.", act.variablesPath)
		return nil
	}
	log.Debugf("create repository in path %s", act.targetDir)

	saveRepository(repoFiles, act.targetDir, act.inputVars)
	return nil
}
