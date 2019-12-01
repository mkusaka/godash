package godash

import (
	"fmt"
	"log"
	"reflect"
	"sort"
)

// use sort for slice
func Sort(elements []int, f func(i, j int) bool) []int {
	cached := elements
	sort.Slice(cached, f)
	return cached
}

func IntFilter(elements []int, f func(e int) bool) []int {
	var filtered []int
	for _, element := range elements {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func Filter(elements interface{}, f func(i int) bool) []interface{} {
	cached := []interface{}{}
	items := reflect.ValueOf(elements)
	log.Println(items.Kind())
	if items.Kind() == reflect.Slice {
		for i := 0; i < items.Len(); i++ {
			if f(i) {
				cached = append(cached, reflect.Indirect(items.Index(i)).Interface())
				fmt.Println(reflect.TypeOf(reflect.Indirect(items.Index(i)).Interface()))
			}
			// fmt.Println(reflect.Indirect(items.Index(i)).Interface())
		}
	} else {
		panic("elements must be slice")
	}
	fmt.Println(reflect.TypeOf(elements))
	return cached
}
