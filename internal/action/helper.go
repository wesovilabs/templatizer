package action

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-git/go-billy/v5"
	log "github.com/sirupsen/logrus"
)

type repoFile struct {
	folder  string
	name    string
	content string
}

func iterateDirs(fs billy.Filesystem, path string) []repoFile {
	repoFiles := make([]repoFile, 0)
	files, err := fs.ReadDir(path)
	if err != nil {
		log.Warnf("unexpected error reading path '%s': %s", path, err.Error())
		return repoFiles
	}
	for _, file := range files {
		filename := filepath.Join(path, file.Name())
		if file.IsDir() {
			repoFiles = append(repoFiles, iterateDirs(fs, filename)...)
			continue
		}
		src, err := fs.Open(filename)
		if err != nil {
			log.Warnf("unexpected error opening file '%s': %s", filename, err.Error())
			continue
		}
		bytes := make([]byte, file.Size())
		if _, err = src.Read(bytes); err != nil {
			log.Warnf("unexpected error reading file '%s': %s", filename, err.Error())
			continue
		}
		repoFiles = append(repoFiles, repoFile{
			folder:  path,
			name:    filepath.Base(filename),
			content: string(bytes),
		})
	}
	return repoFiles
}

func processRepositoryFiles(targetDir string, repoFiles []repoFile, path string, variables Variables) {
	for _, file := range repoFiles {
		folderPath := executeTemplate(file.folder, file.folder, variables)
		if err := os.MkdirAll(filepath.Join(targetDir, folderPath), os.ModePerm); err != nil {
			log.Warnf("unexpected error creating folder '%s': %s", folderPath, err.Error())
		}
		filename := executeTemplate(file.name, filepath.Join(path, folderPath, file.name), variables)
		content := executeTemplate(filename, file.content, variables)
		f, err := os.Create(filepath.Join(targetDir, filename))
		if err != nil {
			log.Warnf("unexpected error creating the file '%s': %s", filename, err.Error())
			continue
		}
		if _, err := f.WriteString(content); err != nil {
			log.Warnf("unexpected error writing content into the file '%s': %s", filename, err.Error())
			continue
		}
	}
}

func executeTemplate(name, content string, variables interface{}) string {
	//t := template.Must(template.New(name).Parse(content))
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
