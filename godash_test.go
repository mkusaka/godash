package godash

import (
	"fmt"
	"reflect"
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

func TestFilter(t *testing.T) {
	array := []int{12, 5, 8, 130, 44}

	filtered := Filter(array, func(i int) bool {
		return array[i] > 10
	})

	fmt.Println("hoge")
	fmt.Println(reflect.TypeOf(filtered))
	fmt.Println(reflect.TypeOf(array))
	fmt.Println(array)

	expected := []int{12, 130, 44}

	if !cmp.Equal(filtered, expected) {
		t.Errorf("expected %v but got %v", expected, filtered)
	}
}
