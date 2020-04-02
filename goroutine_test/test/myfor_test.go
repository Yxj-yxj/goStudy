package main

import (
	"fmt"
	"testing"
	"time"
)

var MAXVALUE = 10000000

func myfor() error {
	arr := make([]int, MAXVALUE)
	rtn := make(chan int)
	start := time.Now()
	for i := 0; i < MAXVALUE; i++ {
		arr[i] = i
	}
	go func() {
		for {
			<-rtn
		}
	}()
	for i := 0; i < MAXVALUE; i++ {
		if arr[i]%3 == 0 {
			rtn <- arr[i] % 3
			//			fmt.Println(arr[i])
		}
	}

	fmt.Println("myfor proj end")
	fmt.Println(time.Since(start))
	return nil
}

func TestMyfor(t *testing.T) {
	if err := myfor(); err != nil {
		t.Error("error")
	}
}
