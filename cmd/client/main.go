package client

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/sirupsen/logrus"
)

//go:embed templatizer-ui/build
var content embed.FS

func clientHandler() http.Handler {
	fsys := fs.FS(content)
	contentStatic, _ := fs.Sub(fsys, "templatizer-ui/build")
	return http.FileServer(http.FS(contentStatic))
}

func Run(port int) {
	mux := http.NewServeMux()
	mux.Handle("/", clientHandler())
	address := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(address, mux); err != nil {
		logrus.Fatal(err)
	}
}
