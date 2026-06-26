package unittest

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		testName string
		numbers  []int
		expected int
	}{
		{
			testName: "multiple parameters",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 55,
		},
		{
			testName: "single parameter",
			numbers:  []int{20},
			expected: 20,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := sum(testCase.numbers...)
			expected := testCase.expected

			if expected != actual {
				t.Errorf("expected %d but got %d", expected, actual)
			}
		})
	}
}
