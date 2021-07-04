package test

import (
	_ "fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	
	res := addUpper(10)
	expectation := 55
	if res != expectation {
		t.Fatalf("AddUpper(10) error, Expectation: %d, Return Value: %d\n",expectation, res)
	}

	t.Logf("AddUpper(10) Success...")
}