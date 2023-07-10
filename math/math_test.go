package math

import (
	"testing"
)

func TestAdd(t *testing.T) {

	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Add(1,3) FAILED. Expected %d,got %d\n", 4, result)
	} else {
		t.Logf("Add(1,3) PASSED. Expected %d, got %d\n", 4, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(5, 0)
	if result != 0 {
		t.Errorf("Divide(5,0) FAILED. Expected %d, got %f\n", 0, result)
	} else {
		t.Logf("Divide(5,0) PASSED. Expected %d, got %f\n", 0, result)
	}

}
