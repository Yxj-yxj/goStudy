package main

import "fmt"

func main() {
	a := 1
	for {
		a += 1
		a -= 1
		fmt.Print(a)
	}
}
