package main

import (
	"testing"
)

func TestExamples(t *testing.T) {

	example1 := input_calc("ahhellllloou")
	if example1 != "YES" {
		t.Log("error should be YES, but got", example1)
		t.Fail()
	}

	example2 := input_calc("hlelo")
	if example2 != "NO" {
		t.Log("error should be NO, but got", example2)
		t.Fail()
	}
}
