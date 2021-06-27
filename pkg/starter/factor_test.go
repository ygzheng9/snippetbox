package starter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFactorsOf(t *testing.T) {
	type checkFunc func([]int) error

	// SECTION 1: checkers
	isEmptyList := func(have []int) error {
		if len(have) > 0 {
			return fmt.Errorf("Expected empty list, found %v.", have)
		}
		return nil
	}

	is := func(want ...int) checkFunc {
		return func(have []int) error {
			if !reflect.DeepEqual(have, want) {
				return fmt.Errorf("Expected list %v, found %v.", want, have)
			}
			return nil
		}
	}

	// SECTION 2: test cases
	tests := [...]struct {
		in    int
		check checkFunc
	}{
		// the first test: factors of 1
		{1, isEmptyList},
		{2, is(2)},
		{3, is(3)},
		{5, is(5)},
		{6, is(2, 3)},
		{7, is(7)},
		{8, is(2, 2, 2)},
		{9, is(3, 3)},
		{2 * 2 * 3 * 3 * 5 * 7 * 11 * 11 * 13, is(2, 2, 3, 3, 5, 7, 11, 11, 13)},
	}

	// SECTION 3: test logic
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Factors of %d", tc.in), func(t *testing.T) {
			factors := factorsOf(tc.in)
			if err := tc.check(factors); err != nil {
				t.Error(err)
			}
		})
	}
}
