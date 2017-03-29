package main

import (
	"fmt"

	"s.mcquay.me/dm/pairsum/pairsum"
)

func main() {
	a := []int{1, 3, 5, 9}
	b := []int{1, 3, 4, 9}
	c := []int{1, 2, 4, 4}
	fmt.Println(pairsum.DoesItSum(a, 8))
	fmt.Println(pairsum.DoesItSum(b, 8))
	fmt.Println(pairsum.DoesItSum(c, 8))
}
