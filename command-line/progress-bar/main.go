package main

import (
	"fmt"
	"time"

	progressbar "github.com/schollz/progressbar/v3"
)

func main() {
	fmt.Println("開始執行進度條測試。")

	progressBar := progressbar.NewOptions(
		1000,                                     // 進度條總數
		progressbar.OptionEnableColorCodes(true), // 是否可以使用字體顏色語法
		progressbar.OptionShowBytes(true),        // 是否顯示傳遞速率(xxx B/s)
		progressbar.OptionSetWidth(15),           // 進度條寬度
		progressbar.OptionSetDescription("[cyan]測試進度[reset] 請稍等..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	for i := 0; i < 1000; i++ {
		// 替進度條進度加一，加滿則結束
		progressBar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}

	progressBar.Clear()
	progressBar.Close()

	fmt.Println("進度條結束。")
}
