package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := 0
	cur := 0
	return func() int {
		if prev == 0 && cur == 0 {
			cur = 1
			return 0
		}
		tmp := prev + cur
		prev = cur
		cur = tmp
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
