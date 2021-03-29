package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	// time area
	// create time area
	// now -time

	cTime, _ := time.Parse("2006-01-02 15:04:05", "2021-03-29 09:00:00")
	fmt.Println(cTime.Format("2006-01-02"))

	dura := time.Since(cTime) // 从 cTime到当前时间间隔多少,前后用正负数表示
	// 如ctime= 2021-04-02 09:00:00
	// 当前时间为 2021/03/29 11:51
	// 则 dura = -89h9m48.378678484s
	fmt.Println(dura)

	dura02 := time.Until(cTime) // 当前时间到cTime时间间隔
	fmt.Println(dura02)

	dura03, err := time.ParseDuration("1h10m10s")

	fmt.Println(dura03, err) // 1h10m10s <nil>

	fmt.Println(dura03.Hours()) // 将dura03换算为小时打印
	// 1.1694444444444445

	fmt.Println(dura03.Minutes()) // 换算为分钟
	fmt.Println(dura03.Seconds()) // 换算为秒

	dayInterval, _ := time.ParseDuration("-24h5m10s")
	fmt.Println(now, now.Add(dayInterval)) // now + dayinterval 时间,相当于添加24h5m10s
	// 2021-03-30 00:24:55.201122408 +0800 CST m=+0.000042499 2021-03-31 00:30:05.201122408 +0800 CST m=+86710.000042499

	yestday := now.Add(dayInterval) // 如果dayInterval 为负数时间,则减少相应时间
	fmt.Println(now)
	fmt.Println(yestday)
	// 2021-03-30 00:27:04.432052327 +0800 CST m=+0.000042301
	// 2021-03-29 00:21:54.432052327 +0800 CST m=-86709.999957699

	fmt.Println(yestday.After(now))  // yestday 时间是否在当前时间之后,为真则为ture, 否则为false
	fmt.Println(yestday.Before(now)) // yestday 时间是否在当前时间之前,为真则为ture, 否则为false

	time.Sleep(time.Second * 3) // 休眠3秒后继续执行
	fmt.Println(time.Now())
}
