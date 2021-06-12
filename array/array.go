package array

import "fmt"

func TestArray() {
	var a [3]int
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
}
