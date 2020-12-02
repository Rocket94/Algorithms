package main

import "fmt"

func bubblingSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
func main() {
	a := []int{1, 5, 3, 6, 8, 9, 1, 2, 4}
	bubblingSort(a)
	fmt.Println(a)
}
