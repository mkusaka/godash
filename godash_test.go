package godash

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntFilter(t *testing.T) {
	filter := func(e int) bool {
		return e >= 10
	}
	array := []int{12, 5, 8, 130, 44}

	filtered := IntFilter(array, filter)

	expected := []int{12, 130, 44}

	if !cmp.Equal(filtered, expected) {
		t.Errorf("expected %v but got %v", expected, filtered)
	}
}
