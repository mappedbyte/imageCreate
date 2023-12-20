package task

import (
	"image-Designer/internal/config"
	"log"
	"time"
)

func ClearIdCache() {
	go func() {
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day()+1, 1, 0, 0, 0, now.Location())
			duration := next.Sub(now)
			log.Println("定时任务已经启动...")
			time.Sleep(duration)
			config.Cache = make(map[string]string)
			log.Println("定时任务已经执行...")
		}
	}()
}
