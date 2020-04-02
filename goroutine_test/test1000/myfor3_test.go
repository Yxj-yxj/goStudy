package main

import (
	"fmt"
	"testing"
	"time"
)

var MAXVALUE = 1000000

func myfor() error {
	done := make(chan bool)
	arr := make([]int, MAXVALUE)
	//rtn := make(chan int)
	start := time.Now()
	for i := 0; i < MAXVALUE; i++ {
		arr[i] = i
	}
	for i := 0; i < 1000; i++ {
		go func(j int) {
			for k := 0; k < 1000; k++ {
				if arr[j*1000+k]%3 == 0 {
					fmt.Println(arr[j*1000+k])
				}
			}
			done <- true
		}(i)
	}
	for i := 0; i < 1000; i++ {
		<-done
	}
	fmt.Println("myfor2 proj end")
	fmt.Println(time.Since(start))
	return nil
}

func TestMyfor(t *testing.T) {
	tests := []struct{ a, b int }{
		{1, 2},
	}

	for _, tt := range tests {
		if err := myfor(); err != nil {
			t.Error("error")
		}
		fmt.Println(tt)
	}

}
