package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = int(now.Year())
	var month int = int(now.Month())
	var day int = int(now.Day())
	fmt.Printf("오늘은%d년 %d월 %d일 입니다.\n", year, month, day)
	fmt.Printf("지금 시각은 %d시 %d분 %d초 입니다.\n", now.Hour(), now.Minute(), now.Second())
}
