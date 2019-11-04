package main

import (
	"fmt"
	"time"
	"github.com/mazrean/TraqWebhook/webhook"
)

func main(){
	err := webhook.Establish()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("タイマーを開始")

	t := 0
	for {
		t++
		err = webhook.UpdateFeed()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Printf("通知: %d \n", t)
		time.Sleep(time.Hour * 1)
	}
}