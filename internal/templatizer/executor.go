package templatizer

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"text/template"

	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/templatizer/pkg/resources"
	"gopkg.in/yaml.v3"
)

type Executor interface {
	LoadTemplatizerconfig() (*resources.Config, error)
	ProcessTemplate(mode string, params map[string]interface{}) (string, []string, error)
}
type executor struct {
	repoURL      string
	configPath   string
	auth         http.AuthMethod
	branch       string
	templateMode string
}

type Option func(*executor)

func WithRepoURL(url string) Option {
	return func(act *executor) {
		act.repoURL = url
	}
}

func WithConfigPath(configPath string) Option {
	return func(act *executor) {
		act.configPath = configPath
	}
}

func WithBasicAuth(username, password string) Option {
	return func(act *executor) {
		act.auth = &http.BasicAuth{
			Username: username,
			Password: password,
		}
	}
}

func WithTokenAuth(token string) Option {
	return func(act *executor) {
		act.auth = &http.TokenAuth{
			Token: token,
		}
	}
}

func WithTemplateMode(templateMode string) Option {
	return func(act *executor) {
		act.templateMode = templateMode
	}
}

func WithBranch(branch string) Option {
	return func(act *executor) {
		act.branch = branch
	}
}

func New(opts ...Option) Executor {
	exec := &executor{
		templateMode: "goTemplate",
		configPath:   ".templatizer.yml",
	}
	for _, opt := range opts {
		opt(exec)
	}
	return exec
}

func (exec *executor) LoadTemplatizerconfig() (*resources.Config, error) {
	w, err := cloneRepositorty(exec.repoURL, exec.branch, exec.auth)
	if err != nil {
		return nil, fmt.Errorf("error while cloning the repository: '%s", err)
	}
	cfg := resources.Config{}
	cfgRemotePath := filepath.Join("/", exec.configPath)
	cfgFile, err := w.Filesystem.Open(cfgRemotePath)
	if err != nil {
		return nil, fmt.Errorf("error openning the config file: '%s", err)
	}
	parentRemoteDir := filepath.Dir(cfgRemotePath)
	files, err := w.Filesystem.ReadDir(parentRemoteDir)
	if err != nil {
		return nil, fmt.Errorf("error reading folder '%s': %s", parentRemoteDir, err.Error())
	}
	var fileInfo fs.FileInfo
	for index := range files {
		if files[index].Name() == filepath.Base(cfgRemotePath) {
			fileInfo = files[index]
			if err != nil {
				return nil, err
			}
		}
	}
	if fileInfo == nil {
		return nil, err
	}
	bytes := make([]byte, fileInfo.Size())
	if _, err = cfgFile.Read(bytes); err != nil {
		return nil, fmt.Errorf("error reading file '%s': %s", exec.configPath, err.Error())
	}
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, fmt.Errorf("error processing config file: '%s'", err)
	}
	return &cfg, nil
}

func (exec *executor) ProcessTemplate(mode string, params map[string]interface{}) (string, []string, error) {
	w, err := cloneRepositorty(exec.repoURL, exec.branch, exec.auth)
	if err != nil {
		return "", nil, fmt.Errorf("error while cloning the repository: '%s", err)
	}
	repoFiles := repoFileSystem{w.Filesystem}.read("/")
	outputDir, err := ioutil.TempDir("", "templatizer")
	if err != nil {
		return outputDir, nil, fmt.Errorf("error creating temporary folder: %s", err)
	}
	filePaths := make([]string, len(repoFiles))
	for idx, file := range repoFiles {
		file.folder = filepath.Join(outputDir, executeTemplate(file.folder, file.folder, params))
		file.name = executeTemplate(file.name, filepath.Join(file.folder, file.name), params)
		file.content = executeTemplate(file.name, file.content, params)
		if err := file.persist(); err != nil {
			return outputDir, nil, err
		}
		filePaths[idx] = file.name
	}

	return outputDir, filePaths, nil
}

func executeTemplate(name, content string, variables interface{}) string {
	t, err := template.New(name).Parse(content)
	if err != nil {
		log.Warn(err)
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, variables); err != nil {
		log.Println("error processing template:", err)
	}
	return buf.String()
}
