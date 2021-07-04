package test

import (
	_ "fmt"
	"testing"
)

func TestGetSub(t *testing.T) {
	
	res := getSub(10,8)
	expectation := 3
	if res != expectation {
		t.Fatalf("getSub(10) error, Expectation: %d, Return Value: %d\n",expectation, res)
	}

	t.Logf("getSub(10) Success...")
}