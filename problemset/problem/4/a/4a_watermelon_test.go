package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	answer := input_calc(8)
	if answer != "YES" {
		t.Log("error should be YES, but got", answer)
		t.Fail()
	}
}
