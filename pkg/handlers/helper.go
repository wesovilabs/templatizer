package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wesovilabs/templatizer/internal/templatizer"
	"github.com/wesovilabs/templatizer/pkg/resources"
)

func processError(err error, c *gin.Context) {
	logrus.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}

func buildExecutor(req resources.ConnectionRequest) templatizer.Executor {
	options := []templatizer.Option{
		templatizer.WithRepoURL(req.URL),
	}
	if req.Auth != nil {
		if req.Auth.Mechanism == "basic" {
			options = append(options, templatizer.WithBasicAuth(req.Auth.Username, req.Auth.Password))
		}
		if req.Auth.Mechanism == "token" {
			options = append(options, templatizer.WithTokenAuth(req.Auth.Token))
		}
	}
	return templatizer.New(options...)
}
