package utils

import (
	"fmt"
	"os"
	"screenshot-ai/pkg/log"
	"sync"
	"time"
)

var lock sync.Mutex

const (
	Pic = "\\static\\screenshot\\"
)

func cleanStaitc() {
	wd, _ := os.Getwd()
	staticPath := wd + Pic
	now := time.Now().Format("2006-01-02")
	if _, err := os.Stat(staticPath); err != nil {
		log.Logobj.Errorf("dir not found,clean failed:%s", err)
	}
	scrFilePath := staticPath + now + "\\"
	err := os.RemoveAll(scrFilePath)
	if err != nil {
		log.Logobj.Errorf("remove dir failed: %s", err)
	}
	fmt.Println("Clean files success")
}

func CleanPeriodly() {
	ticker := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-ticker.C:
			// 每小时触发清理
			fmt.Println("Cleaning static files...")
			lock.Lock()
			cleanStaitc()
			lock.Unlock()
		}
	}
}
