package templatizer

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	log "github.com/sirupsen/logrus"
)

type repoFile struct {
	folder  string
	name    string
	content string
}

func (rf repoFile) persist() error {
	if err := os.MkdirAll(rf.folder, os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(rf.name)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(rf.content); err != nil {
		return err
	}
	println(rf.name)
	return nil
}

type repoFileSystem struct {
	fs billy.Filesystem
}

func (rf repoFileSystem) read(path string) []repoFile {
	repoFiles := make([]repoFile, 0)
	files, err := rf.fs.ReadDir(path)
	if err != nil {
		log.Warnf("unexpected error reading path '%s': %s", path, err.Error())
		return repoFiles
	}
	for _, file := range files {
		filename := filepath.Join(path, file.Name())
		if file.IsDir() {
			repoFiles = append(repoFiles, rf.read(filename)...)
			continue
		}
		src, err := rf.fs.Open(filename)
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
