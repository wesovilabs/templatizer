package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesovilabs/templatizer/pkg/resources"
)

func LoadParamaters(c *gin.Context) {
	var req resources.ConnectionRequest
	if err := c.BindJSON(&req); err != nil {
		processError(err, c)
		return
	}
	executor := buildExecutor(req)

	templateSettings, err := executor.LoadTemplatizerconfig()
	if err != nil {
		processError(err, c)
		return
	}
	c.JSON(http.StatusOK, templateSettings)
}
