package datastruct

import (
	"fmt"
	"testing"
)

func TestBuildHeap(t *testing.T) {
	h:=BuildHeap([]int{5,3,4,6,7,8,1,2})
	fmt.Println(h)
	fmt.Println(h.RemoveTop())
	fmt.Println(h.RemoveTop())
	fmt.Println(h.RemoveTop())
	h.Insert(10)
	h.Insert(4)
	h.Insert(0)
	fmt.Println(h)
}
