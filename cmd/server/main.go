package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wesovilabs/templatizer/pkg/api"
)

const (
	localAddress = ":5001"
	basePath     = "/api"
)

func main() {
	router := api.SetUpRouter(basePath)
	if err := router.Run(localAddress); err != nil {
		logrus.Fatal(err)
	}
}
