package main

import (
	"time"
)

//1. time包
//2. time.Time类型，用来表示时间
//3. 获取当前时间，now:=time.Now()
//4. time.Now().Day(),time.Now().Minute(),time.Now().Month(),time.Now().Year()
//5. 格式化，fmt.Printf("%02d/%02d%02d %02d:%02d:%02d) now.Year(),----)
//6. time.Duration用来表示纳秒
//7. 一些常量： 
// const (
// 	Nanosecond  Duration = 1
// 	Microsecond          = 1000 * Nanosecond //1微秒 = 1000 纳秒
// 	Millisecond          = 1000 * Microsecond //1毫秒 = 1000 微秒
// 	Second               = 1000 * Millisecond  // 1秒 = 1000 毫秒
// 	Minute               = 60 * Second
// 	Hour                 = 60 * Minute
// )

// 8. 格式化：

now := time.Now()
fmt.Println(now.Format("02/1/2006 15:04"))
fmt.Println(now.Format("2006/1/02 15:04"))
fmt.Println(now.Format("2006/1/02"))

//6、获取当前时间，并格式化成2017/06/15 08:05:00形式
func ass6() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05")) //必须是这个语法“2006/01/02 15:04:05”
}

//7、统计一段代码的执行耗时，单位精确到微秒

func main() {
	start := time.Now().UnixNano() //uninx时间，1970年utc年1月1日到现在的纳秒
	ass6()
	end := time.Now().UnixNano()
	fmt.Printf("cost:%d us\n", (end-start)/1000/1000)
}
