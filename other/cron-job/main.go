package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("程式開始執行時間:", getNowStr())

	// 建立排程
	cronjob := cron.New(cron.WithSeconds())

	// 每兩秒執行一次動作
	cronjob.AddFunc("*/2 * * * * *", func() {
		printCronJobProgessBySec(2)
	})

	// 每五秒執行一次動作
	cronjob.AddFunc("*/5 * * * * *", func() {
		printCronJobProgessBySec(5)
	})

	// 排程開始
	cronjob.Start()

	time.Sleep(10 * time.Second)

	fmt.Println("程式結束時間:", getNowStr())
}

// printCronJobProgessBySec 透過秒數打印排程進度
func printCronJobProgessBySec(sec int) {
	fmt.Printf("--- [%s] 執行了每隔%d秒執行一次的排程。\n", getNowStr(), sec)
}

// getNow 取得現在時間字串
func getNowStr() string {
	var taipeiZone, _ = time.LoadLocation("Asia/Taipei")

	return time.Now().In(taipeiZone).Format("15:04:05")
}
