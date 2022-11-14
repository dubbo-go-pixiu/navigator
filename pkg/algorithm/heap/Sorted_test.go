package heap

import (
	"fmt"
	"testing"
)

func TestExampleIntheap(t *testing.T) {
	var compare = func(one int, other int) int {
		return -(one - other)
	}
	h := Sorted[int]{}
	h.Init(compare, []int{34, 231, 4231, 0})

	h.Push(8)
	h.Push(9)
	h.Push(8)
	h.Push(1)
	h.Push(8)
	h.Push(3)
	h.Push(2)
	h.Push(8)
	fmt.Printf("minimum: %d\n", h.array[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}

}
