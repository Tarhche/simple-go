package unittest

import "testing"

func TestSum(t *testing.T) {
	var expected, actual int

	// 1. single parameter
	expected = 3
	actual = sum(3)
	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}

	// 2. multiple parameters
	expected = 10
	actual = sum(1, 2, 3, 4)

	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
