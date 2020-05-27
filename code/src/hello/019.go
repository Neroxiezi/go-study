package main
import (
    "fmt"
    "time"
)

func main(){
    now := time.Now() // 时间对象
    fmt.Println(now)
    year :=now.Year()
    fmt.Println(year)

    // 定时器
    for tmp := range time.Tick(time.Second) {
    		fmt.Println(tmp)
    }
}
