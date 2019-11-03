package main

import (
	"fmt"
	_ "net/http"
	"github.com/mmcdole/gofeed"
	"github.com/mazrean/TraqWebhook/webhook"
	"github.com/jmoiron/sqlx"
)

var FeedURL = os.Getenv("FEED_URL")

func main(){
	// 変数tickerに指定の時間を入れる
	ticker := time.NewTicker(time.hour * 1)
	fmt.Println("タイマーを開始")
	// go funcで並列処理を実行する
	go func() {
		// for文で繰り返しイベントを発火させるループを組む
		for t := range ticker.C {
			err := webhook.UpdateFeed()
			if err != nil{
				fmt.Println(err)
			}
			// ticker.Cの通知があれば、Printする（イベント）
			fmt.Printf("通知: %s \n", t)
		}
	}()
}