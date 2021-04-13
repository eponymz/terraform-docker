package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func ExecExcept(exceptions []string, commandName string, args ...string) string {
	directory := args[0]
	if !SliceEmpty(exceptions) && InExceptions(exceptions, directory) {
		logrus.Tracef("Length of exceptions: %d", len(exceptions))
		logrus.Tracef("Skipping directory %s as it is in the passed exceptions %s.", directory, exceptions)
		return ""
	}
	var safeCommand string
	var safeArgs []string
	if strings.Contains(commandName, " ") {
		logrus.Tracef("CommandName '%s' contained spaces, splitting...", commandName)
		split := strings.Split(commandName, " ")
		safeCommand = split[0]
		safeArgs = append(split[1:], args[1:]...)
	} else {
		safeCommand = commandName
		safeArgs = args[1:]
	}
	logrus.Tracef("Running command %s on %s with args %s", safeCommand, directory, safeArgs)
	command := exec.Command(safeCommand, append(safeArgs, directory)...)
	logrus.Tracef("Command %s path: %s", safeCommand, command.Path)
	if !strings.Contains(command.Path, "/") {
		logrus.Fatalf("Command %s not in PATH!", command.Path)
	}
	out, _ := command.CombinedOutput()
	return string(out)
}

func DirTreeList(directory string) []string {
	logrus.Tracef("Running DirTreeList on %s", directory)
	var response []string
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				logrus.Errorf("DirTreeList had an error with %s: %s", directory, err.Error())
				return err
			}
			if info.IsDir() {
				response = append(response, path)
			}
			return nil
		})
	if err != nil {
		logrus.Error(err)
	}
	logrus.Tracef("DirTreeList returned %s for %s", response, directory)
	return response
}

func ExecExceptR(exceptions []string, command string, args ...string) string {
	var response string
	directory := args[0]
	for _, dir := range DirTreeList(directory) {
		response += ExecExcept(exceptions, command, append([]string{dir}, args[1:]...)...) + "\n"
	}
	return response
}

func ExecExceptRCompare(exceptions []string, compare string, command string, args ...string) string {
	var response string
	directory := args[0]
	for _, dir := range DirTreeList(directory) {
		file := fmt.Sprintf("%s/%s", dir, compare)
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			logrus.Tracef("ExecExceptRCompare found file %s", file)
			currentGen := ExecExcept(exceptions, command, append([]string{dir}, args[1:]...)...)
			existing, _ := ioutil.ReadFile(file)
			if strings.Compare(currentGen, string(existing)) != 0 {
				response += fmt.Sprintf("Command %s returned differences from %s\n", command, file)
			}
		} else {
			logrus.Tracef("ExecExceptRCompare could not find file %s, skipping %s", file, dir)
		}
	}
	return response
}
