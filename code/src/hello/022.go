package main

import (
    "fmt"
)

func main(){
    ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
    		for i := 0; i < 100; i++ {
    			ch1 <- i
    		}
    		close(ch1)
    }()
    go func() {
        for {
            i, ok := <-ch1 // 通道关闭后再取值ok=false
            if !ok {
				break
			}
			ch2 <- i * i
        }
        close(ch2)
    }()
    // 在主goroutine中从ch2中接收值打印
    for i := range ch2 { // 通道关闭后会退出for range循环
    	fmt.Println(i)
    }
}
