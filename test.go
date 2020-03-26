package main

import (
	"fmt"
	"log"
)

func g() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

func protect(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		//if err := recover();err!=nil{
		//	log.Printf("run time panic: %v", err)
		//}
	}()
	log.Println("start")
	log.Println("start")
	g() //   possible runtime-error
}

func main() {
	protect(g)
	fmt.Println("program end")
}
