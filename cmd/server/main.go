package main

import "github.com/wesovilabs/templatizer/pkg/api"

const (
	localAddress = ":5001"
	basePath     = "/api"
)

func main() {
	router := api.SetUpRouter(basePath)
	router.Run(localAddress)
}
