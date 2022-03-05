package server

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/wesovilabs/templatizer/pkg/api"
)

func Run(port int) {
	router := api.SetUpRouter("/api")
	address := fmt.Sprintf(":%d", port)
	if err := router.Run(address); err != nil {
		logrus.Fatal(err)
	}
}
