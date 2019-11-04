package main

import (
	"fmt"
	"os"
	"time"
	"github.com/mazrean/TraqWebhook/webhook"
)

var FeedURL = os.Getenv("FEED_URL")

func main(){
	webhook.Establish()
	fmt.Println("タイマーを開始")

	t := 0
	for {
		t++
		err := webhook.UpdateFeed()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Printf("通知: %d \n", t)
		time.Sleep(time.Second * 10)
	}
}