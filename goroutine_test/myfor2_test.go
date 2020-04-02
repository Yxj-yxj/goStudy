package main

import (
	"fmt"
	"testing"
	"time"
)

var MAXVALUE = 100000000

func myfor() error {
	done := make(chan bool)
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
	for i := 0; i < 100; i++ {
		go func(c chan int, arrslice []int) {
			for k := 0; k < 1000000; k++ {
				if arrslice[k]%3 == 0 {
					c <- arr[k] % 3
				}
			}
			done <- true
		}(rtn, arr[i*1000000:(i+1)*1000000])
	}
	for i := 0; i < 100; i++ {
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
