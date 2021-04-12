package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

func SliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func InExceptions(exceptions []string, term string) bool {
	for _, e := range exceptions {
		regex := fmt.Sprintf("^%s.*", e)
		matched, _ := regexp.MatchString(regex, term)
		if matched {
			logrus.Tracef("Term %s was matched in %s with %s", term, exceptions, regex)
			return true
		}
	}
	logrus.Tracef("Term %s was not matched in %s", term, exceptions)
	return false
}

func ExecExcept(exceptions []string, commandName string, args ...string) string {
	directory := args[0]
	if !InExceptions(exceptions, directory) {
		var safeCommand string
		var safeArgs []string
		if strings.Contains(commandName, " ") {
			logrus.Tracef("CommandName '%s' contained spaces, splitting...", commandName)
			split := strings.Split(commandName, " ")
			safeCommand = split[0]
			safeArgs = append(split[1:], args...)
		} else {
			safeCommand = commandName
			safeArgs = args
		}
		logrus.Tracef("Running command %s on %s with args %s", safeCommand, directory, safeArgs)
		command := exec.Command(safeCommand, safeArgs...)
		out, _ := command.CombinedOutput()
		return string(out)
	}
	logrus.Tracef("Skipping directory %s as it is in the passed exceptions %s.", directory, exceptions)
	return ""
}

func DirTreeList(directory string) []string {
	logrus.Tracef("Running DirTreeList on %s", directory)
	var response []string
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				logrus.Errorf("DirTreeList had an error with %s: \n%s", directory, err.Error())
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
