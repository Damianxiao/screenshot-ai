package keyinput

import (
	"errors"
	"fmt"
	"screenshot-ai/network"
	"screenshot-ai/pkg/utils"
	"sync"
	"time"

	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
)

type bind struct {
	keys     []string
	donkeys  []string
	done     chan struct{}
	lastShot time.Time
	mu       sync.Mutex
	coolDown time.Duration
}

func NewBind(keys, donekeys []string) *bind {
	return &bind{
		keys:     keys,
		donkeys:  donekeys,
		done:     make(chan struct{}),
		lastShot: time.Unix(0, 0),
		coolDown: 3 * time.Second,
	}
}

func (b *bind) StartEvents() {
	hook.Register(hook.MouseWheel, b.keys, func(e hook.Event) {
		b.mu.Lock()
		defer b.mu.Unlock()
		now := time.Now()
		if now.Sub(b.lastShot) < b.coolDown {
			waitTime := b.coolDown - now.Sub(b.lastShot)
			fmt.Printf("too many type in, retry in %s\n", waitTime)
			return
		}
		img, err := TakeScreenShot()
		if err != nil {
			fmt.Println("screenshot failed", err)
		}
		b.lastShot = now
		// fmt.Println("screenshot success")

		// broadcast to phone
		base64Img, _ := utils.ImgToBase64(img)
		network.HubServer.Broadcast <- network.WebSocketMessage{
			Type:    network.TypeScreenShot,
			Content: base64Img,
		}
	})
	hook.Register(hook.KeyDown, b.donkeys, func(e hook.Event) {
		b.done <- struct{}{}
		hook.End()
	})
	fmt.Println("listening start")

	s := hook.Start()
	<-hook.Process(s)
}

func (b *bind) Listen() {
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("listening keyborad...")
		case <-b.done:
			// hook.End()
			fmt.Println("stoping the listener...")
		}
	} // hook.Register(hook.KeyDown)
}

func TakeScreenShot() (string, error) {
	nums := screenshot.NumActiveDisplays()
	if nums == 0 {
		return "", errors.New("none screen captcha")
	}
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	file, err := utils.SaveImgToStatic(img)
	if err != nil {
		return "", err
	}
	return file, nil

	// call ai
	// msg, err := api.CallApiWithimg(file)
	// fmt.Println("ai resp: " + msg)
	// if err != nil {
	// 	return "", err
	// }

	// pick monitor 1
}
