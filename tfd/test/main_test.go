package test

import (
	"os"
	"testing"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/cmd"
)

func TestMain(m *testing.M) {
	cmd.InitConfig() // Unified logging configuration from app
	code := m.Run()
	os.Exit(code)
}
