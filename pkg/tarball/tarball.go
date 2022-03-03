package tarball

import (
	"archive/tar"
	"io"
	"os"
	"strings"
)

type Tarball interface {
	Compress() error
}

func addToArchive(tw *tar.Writer, rootDir string, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := tar.FileInfoHeader(info, strings.TrimPrefix(info.Name(), rootDir))
	if err != nil {
		return err
	}
	header.Name = strings.TrimPrefix(filename, rootDir)
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}

func New(path string, rootDir string, files []string) Tarball {
	return &targz{path, rootDir, files}
}
