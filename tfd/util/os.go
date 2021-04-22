package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func CaptureStdout() (*os.File, *os.File, *os.File) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	return stdout, r, w
}

func ReleaseStdout(stdout *os.File, r *os.File, w *os.File) string {
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdout
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

func SafeChangeDir(path string) error {
	if err := os.Chdir(path); err != nil {
		if patherr := err.(*os.PathError); patherr.Err != nil {
			logrus.Trace(err)
			return patherr.Err
		}
	}
	return nil
}
