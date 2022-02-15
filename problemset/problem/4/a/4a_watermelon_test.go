package main

import (
	"testing"
)

func TestExamples(t *testing.T) {

	example := input_calc(8)
	if example != "YES" {
		t.Log("error should be YES, but got", example)
		t.Fail()
	}
}
