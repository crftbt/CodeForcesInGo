package main

import (
	"testing"
)

func TestExamples(t *testing.T) {

	example := input_calc(6, 6, 4)
	if example != 4 {
		t.Log("error should be 4, but got", example)
		t.Fail()
	}
}
