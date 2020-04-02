package main

//var MAXVALUE = 100000000
//
//func myfor() error {
//	done := make(chan bool)
//	arr := make([]int, MAXVALUE)
//	rtn := make(chan int)
//	start := time.Now()
//	for i := 0; i < MAXVALUE; i++ {
//		arr[i] = i
//	}
//	for i := 0; i < 10000; i++ {
//		go func(j int) {
//			for k := 0; k < 10000; k++ {
//				if arr[j*10000+k]%3 == 0 {
//					rtn <- arr[j*10000+k] % 3
//				}
//			}
//			done <- true
//		}(i)
//	}
//	go func() {
//		for {
//			<-rtn
//		}
//	}()
//	for i := 0; i < 10000; i++ {
//		<-done
//	}
//	fmt.Println("myfor2 proj end")
//	fmt.Println(time.Since(start))
//	return nil
//}
//
//func main() {
//	_ = myfor()
//}
