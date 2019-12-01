# godash
provide lodash like api

# useage
## filter

```go
package main

import "github.com/mkusaka/godash"

func main() {
	filtered := godash.Filter(&array, func(i int) bool {
		return array[i] > 10
	})

	returned := []int{}
	for _, f := range filtered {
		returned = append(returned, f.(int)) // still typecast needed...
  }
  fmt.Println(returned)
}
```

# TODO
- use reflect package
- add test
- filter/map etc...
