package tarball

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

type targz struct {
	path    string
	rootDir string
	files   []string
}

func (c *targz) Compress() error {
	out, err := os.Create(c.path)
	if err != nil {
		return err
	}
	defer out.Close()
	return c.createArchive(out)
}

func (c *targz) createArchive(buf io.Writer) error {
	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range c.files {
		if err := addToArchive(tw, c.rootDir, file); err != nil {
			return err
		}
	}
	return nil
}
