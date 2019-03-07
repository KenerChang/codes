package fibonacci

import (
	"testing"
	"fmt"
)

func TestCompute(t *testing.T) {
	target := 2
	result := Compute(target)
	msgTemplate := "compute %d failed, got: %d"
	if result.Uint64() != 1 {
		t.Errorf(msgTemplate, target, result)
	}

	target = 10
	result = Compute(target)
	if result.Uint64() != 55 {
		t.Errorf(msgTemplate, target, result)
	}

	target = 20
	result = Compute(target)
	if result.Uint64() != 6765 {
		t.Errorf(msgTemplate, target, result)
	}

	target = 30
	result = Compute(target)
	if result.Uint64() != 832040 {
		t.Errorf(msgTemplate, target, result)
	}

	target = 8181
	fmt.Printf("%s", Compute(target))
}