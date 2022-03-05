package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/templatizer/cmd/client"
	"github.com/wesovilabs/templatizer/cmd/server"
)

func main() {
	serverPort := 16917
	log.Infof("server running on %d", serverPort)
	go server.Run(serverPort)
	clientPort, err := getFreePort()
	if err != nil {
		log.Fatalf("error allocating client port: %s", err.Error())
	}
	log.Infof("client running on %d", clientPort)
	go client.Run(clientPort)
	browserURL := fmt.Sprintf("http://localhost:%d", clientPort)
	go openbrowser(browserURL)

	c := make(chan os.Signal)
	done := make(chan bool, 1)

	go func() {
		<-c
		log.Info("Hasta la vista!")
		done <- true
	}()
	<-done
}

func openbrowser(url string) {
	log.Infof("open browser at %s", url)
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	if err := exec.Command(cmd, args...).Start(); err != nil {
		log.Warn("browser could not be open automatically")
		log.Infof("Try open the followinf URL %s", url)
	}
}

func getFreePort() (int, error) {
	a, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return -1, err
}
