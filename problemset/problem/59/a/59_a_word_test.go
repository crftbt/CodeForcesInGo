package main

import (
	"testing"
)

func TestExamples(t *testing.T) {

	example1 := input_calc("HoUse")
	if example1 != "house" {
		t.Log("error should be house, but got", example1)
		t.Fail()
	}

	example2 := input_calc("ViP")
	if example2 != "VIP" {
		t.Log("error should be VIP, but got", example2)
		t.Fail()
	}

	example3 := input_calc("maTRIx")
	if example3 != "matrix" {
		t.Log("error should be matrix, but got", example3)
		t.Fail()
	}
}
