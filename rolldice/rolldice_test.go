package rolldice

import (
	"testing"
	"fmt"
)

func TestCompute(t *testing.T) {
	msgTemplate := "expect %d but got %d"
	n := 1
	var target uint64 = 1
	result := Compute(n).Uint64()
	if target != result {
		t.Errorf(msgTemplate, target, result)
	}

	n = 2
	target = 2
	result = Compute(n).Uint64()
	if target != result {
		t.Errorf(msgTemplate, target, result)
	}
	
	n = 4
	target = 8
	result = Compute(n).Uint64()
	if target != result {
		t.Errorf(msgTemplate, target, result)
	}

	n = 7
	target = 63
	result = Compute(n).Uint64()
	if target != result {
		t.Errorf(msgTemplate, target, result)
	}

	n = 8
	target = 125
	result = Compute(n).Uint64()
	if target != result {
		t.Errorf(msgTemplate, target, result)
	}

	n = 10
	target = 492
	result = Compute(n).Uint64()
	if target != result {
		t.Errorf(msgTemplate, target, result)
	}

	n = 610
	fmt.Printf("%s", Compute(n))
}