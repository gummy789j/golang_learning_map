package monster

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := Monster{
		Name : "Steven",
		Age : 24,
		Skill : "hard-working",
	}
	res := monster.Store()

	if !res {
		t.Fatalf("monster.Store() Fail Expectation: %v Return Value: %v\n",true,res)
	}
	t.Logf("monster.Store() Success !")
}

func TestReStore(t *testing.T) {

	var monster Monster

	res := monster.ReStore()

	if !res {
		t.Fatalf("monster.ReStore() Fail Expectation: %v Return Value: %v\n",true,res)
	}
	if monster.Name != "Steven" {
		t.Fatalf("monster.ReStore() Fail Expectation: %v Return Value: %v\n","Steven",monster.Name)
	}
	t.Logf("monster.ReStore() Success !")
}