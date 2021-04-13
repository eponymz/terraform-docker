package test

import (
	"testing"
	"tfd/util"
)

func TestSliceContains(t *testing.T) {
	slice := []string{"some", "test", "values"}
	term := "values"
	if !util.SliceContains(slice, term) {
		t.Fatal("SliceContains(slice, term) wants true, got false.")
	}
}

func TestSliceEmpty(t *testing.T) {
	slice := []string{}
	if !util.SliceEmpty(slice) {
		t.Fatal("SliceEmpty(slice) wants true, got false.")
	}
}

func TestInExceptions(t *testing.T) {
	exceptions := []string{"this", "that"}
	term := "thats/so/raven"
	if !util.InExceptions(exceptions, term) {
		t.Fatal("InExceptions wants true, got false")
	}
}
