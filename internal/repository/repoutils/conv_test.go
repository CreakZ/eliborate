package repoutils_test

import (
	"eliborate/internal/repository/repoutils"
	"reflect"
	"testing"
)

func TestConvertMeiliHitsToIntSlice(t *testing.T) {
	hits1 := []any{
		map[string]any{
			"id":        2,
			"something": (*int)(nil),
		},
		map[string]any{
			"id":        3,
			"something": (*int)(nil),
		},
		map[string]any{
			"id":        5,
			"something": (*int)(nil),
		},
	}

	expectedInts1 := []int{2, 3, 5}

	ints1 := repoutils.ConvertMeiliHitsToIntSlice(hits1)

	if !reflect.DeepEqual(expectedInts1, ints1) {
		t.Errorf("test 1 failed: expected %v, got %v", expectedInts1, ints1)
	}
}
