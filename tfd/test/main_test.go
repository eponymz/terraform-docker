package test

import (
	"os"
	"testing"
	"tfd/cmd"
)

func TestMain(m *testing.M) {
	cmd.InitConfig() // Unified logging configuration from app
	code := m.Run()
	os.Exit(code)
}
