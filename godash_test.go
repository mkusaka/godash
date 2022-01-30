package godash

import (
	"reflect"
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

func TestStringToIntMap(t *testing.T) {
	array := []string{"11", "234", "234234"}

	mapped := Map(array, func(i string) int {
		return len(i)
	})

	expected := []int{2, 3, 6}

	if !cmp.Equal(mapped, expected) {
		t.Errorf("expected %v but got %v", expected, mapped)
	}
}

func TestReverseString(t *testing.T) {
	array := []string{"11", "234", "234234"}

	reversed := Reverse(array)

	expected := []string{"234234", "234", "11"}

	if !cmp.Equal(reversed, expected) {
		t.Errorf("expected %v but got %v", expected, reversed)
	}
}

func TestDropStringArr(t *testing.T) {
	type args[T any] struct {
		input []T
		count int
	}
	tests := []struct {
		name string
		args args[string]
		want []string
	}{
		{
			name: "String Array with 0 drop",
			args: args[string]{
				input: []string{"1", "3", "4"},
				count: 0,
			},
			want: []string{"1", "3", "4"},
		},
		{
			name: "String Array with -1 drop",
			args: args[string]{
				input: []string{"1", "3", "4"},
				count: -1,
			},
			want: []string{"1", "3", "4"},
		},
		{
			name: "String Array with len(array) drop",
			args: args[string]{
				input: []string{"1", "3", "4"},
				count: 3,
			},
			want: []string{},
		},
		{
			name: "String Array with len(array) + 1 drop",
			args: args[string]{
				input: []string{"1", "3", "4"},
				count: 4,
			},
			want: []string{},
		},
		{
			name: "String Array with 0 < count < len(array) drop(1)",
			args: args[string]{
				input: []string{"3", "4"},
				count: 1,
			},
			want: []string{},
		},
		{
			name: "String Array with 0 < count < len(array) drop(2)",
			args: args[string]{
				input: []string{"4"},
				count: 1,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Drop(tt.args.input, tt.args.count)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGenerator() got = %v, want %v", got, tt.want)
			}
		})
	}
	// TODO: test with other types
}
