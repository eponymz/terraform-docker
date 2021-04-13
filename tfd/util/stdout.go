package util

import (
	"io/ioutil"
	"os"
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
