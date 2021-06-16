package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"gitlab.com/edquity/devops/terraform-docker.git/tfd/cmd"
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logrus.Warnf("Received %s", sig)
		os.Exit(1)
	}()

	cmd.Execute()
}
