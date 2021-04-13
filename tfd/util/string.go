package util

import (
	"fmt"
	"regexp"

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

func SliceEmpty(slice []string) bool {
	for _, s := range slice {
		if len(s) > 0 {
			return false
		}
	}
	return true
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
