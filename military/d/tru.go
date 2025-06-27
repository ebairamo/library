package main

import "fmt"

func main() {
	supply := map[int]int{
		1: 21,
		2: 23,
		3: 234,
		4: 4325,
	}
	for k := range supply {
		fmt.Println(k, supply[k])
	}
}
