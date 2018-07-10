package main

import (
	"fmt"
	cron2 "github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("cron starting......")
	cron := cron2.New()
	//每秒执行
	cron.AddFunc("* * * * * *", func() {
		fmt.Println("aaaa")
	})
	//每秒执行
	cron.AddFunc("* * * * * *", func() {
		fmt.Println("BBB")
	})
	cron.Start()
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
