package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wesovilabs/templatizer/pkg/resources"
	"github.com/wesovilabs/templatizer/pkg/tarball"
)

func ProcessTemplate(c *gin.Context) {
	var req resources.ProcessTemplateRequest
	if err := c.BindJSON(&req); err != nil {
		processError(err, c)
		return
	}
	executor := buildExecutor(req.ConnectionRequest)
	logrus.Info("processing teplate...")
	rootDir, filePaths, err := executor.ProcessTemplate(req.Mode, req.Params)
	if err != nil {
		processError(err, c)
		return
	}

	outputPath, err := os.CreateTemp("", "templatizer")
	if err != nil {
		processError(err, c)
		return
	}
	logrus.Infof("create temporary file %s", outputPath.Name())
	defer func() {
		outputPath.Close()
		//	os.Remove(outputPath.Name())
	}()
	logrus.Infof("create tarzg file %s", outputPath.Name())
	if err := tarball.New(outputPath.Name(), rootDir, filePaths).Compress(); err != nil {
		processError(err, c)
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=templatizer.tar.gz")
	c.Header("Content-Type", "application/octet-stream")
	http.ServeFile(c.Writer, c.Request, outputPath.Name())
}
