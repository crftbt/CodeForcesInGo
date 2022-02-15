package main

import (
	"testing"
)

func TestExamples(t *testing.T) {

	example1 := input_calc(2, 4)
	if example1 != 4 {
		t.Log("error should be 4, but got", example1)
		t.Fail()
	}

	example2 := input_calc(3, 3)
	if example2 != 4 {
		t.Log("error should be 4, but got", example2)
		t.Fail()
	}
}
