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
	if _, err := os.Stat(staticPath); err != nil {
		log.Logobj.Errorf("dir not found,clean failed:%s", err)
	}
	err := os.RemoveAll(staticPath)
	if err != nil {
		log.Logobj.Errorf("remove dir failed: %s", err)
	}
	fmt.Println("Clean files success")
}

func CleanPeriodly() {
	ticker := time.NewTicker(24 * time.Hour)
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
