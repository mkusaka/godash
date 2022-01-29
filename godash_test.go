package godash

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntArrayFilter(t *testing.T) {
	array := []int{12, 5, 8, 130, 44}

	filtered := Filter(array, func(i int) bool {
		return i > 10
	})

	expected := []int{12, 130, 44}

	if !cmp.Equal(filtered, expected) {
		t.Errorf("expected %v but got %v", expected, filtered)
	}
}

func TestStringArrayFilter(t *testing.T) {
	array := []string{"1", "a", "sss", "333", "22", "2"}

	filtered := Filter(array, func(i string) bool {
		return len(i) <= 2
	})

	expected := []string{"1", "a", "22", "2"}

	if !cmp.Equal(filtered, expected) {
		t.Errorf("expected %v but got %v", expected, filtered)
	}
}

func MapStringToIntFilter(t *testing.T) {
	array := []string{"11", "234", "234234"}

	mapped := Map(array, func(i string) int {
		return len(i)
	})

	expected := []int{2, 3, 6}

	if !cmp.Equal(mapped, expected) {
		t.Errorf("expected %v but got %v", expected, mapped)
	}
}
