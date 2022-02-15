package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	answer := input_calc(6, 6, 4)
	if answer != 4 {
		t.Log("error should be 4, but got", answer)
		t.Fail()
	}
}
